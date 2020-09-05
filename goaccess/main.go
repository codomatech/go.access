package main

import "fmt"

import parsers "github.com/codomatech/goaccess.go/parsers"
import analyzers "github.com/codomatech/goaccess.go/analyzers"
import output "github.com/codomatech/goaccess.go/output"

func main() {
	fmt.Println("parsers.plugins: ", parsers.Plugins())
	fmt.Println("analyzers.plugins: ", analyzers.Plugins())
	fmt.Println("output.plugins: ", output.Plugins())
}

// code generation

//go:generate ./scripts/package-plugins.sh parsers github.com/codomatech/goaccess.go/parsers func(string)common.AccessRecord Parse parsers/plugins.go
//go:generate ./scripts/package-plugins.sh analyzers github.com/codomatech/goaccess.go/analyzers func([]common.AccessRecord)common.AnalysisResult Analyze analyzers/plugins.go
//go:generate ./scripts/package-plugins.sh output github.com/codomatech/goaccess.go/output func([]common.AnalysisResult,string) Output output/plugins.go
