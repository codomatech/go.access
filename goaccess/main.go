package main

import "bufio"
import "flag"
import "log"
import "os"
import "sort"

import "github.com/codomatech/go.access/common"
import parsers "github.com/codomatech/go.access/parsers"
import analyzers "github.com/codomatech/go.access/analyzers"

//import output "github.com/codomatech/go.access/output"

var arginformat = flag.String("format", "clf", "input log format")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		log.Fatalf("Wrong number of arguments %d. You need to specify a log file", len(args))
	}

	fpath := args[0]

	/*
		fmt.Println("parsers.plugins: ", parsers.Plugins())
		fmt.Println("analyzers.plugins: ", analyzers.Plugins())
		fmt.Println("output.plugins: ", output.Plugins())
	*/

	// parse
	//

	parse, found := parsers.Plugins()[*arginformat]

	if !found {
		log.Fatalf("Input format `%s` is not supported", *arginformat)
	}

	log.Printf("parse function: %+v\n", parse)

	file, err := os.Open(fpath)

	if err != nil {
		log.Fatalf("File `%s` cannot be opened", fpath)
	}
	defer file.Close()

	maxlines := int64(10240)
	stat, err := file.Stat()
	if err != nil {
		maxlines = stat.Size() / 512
	}

	records := make([]common.AccessRecord, 0, maxlines)

	scanner := bufio.NewScanner(file)
	maxdayseen := ""
	sorted := true
	for scanner.Scan() {
		record := parse(scanner.Text())
		if record.Day < maxdayseen {
			sorted = false
		}
		maxdayseen = record.Day
		records = append(records, record)
	}

	if !sorted {
		sort.Slice(records, func(i, j int) bool {
			return records[i].Day < records[j].Day
		})
	}

	log.Printf("read in %d records", len(records))

	// analyze
	//

	analyses := analyzers.Plugins()
	results := make([]common.AnalysisResult, 0, len(analyses))

	for _, analyze := range analyses {
		results = append(results, analyze(records))
	}

	log.Printf("results: %+v", results)

}

// code generation

//go:generate spluggy -func Parse -pkg github.com/codomatech/go.access/parsers ./parsers
//go:generate spluggy -func Analyze -pkg github.com/codomatech/go.access/analyzers ./analyzers
//go:generate spluggy -func Output -pkg github.com/codomatech/go.access/output ./output
