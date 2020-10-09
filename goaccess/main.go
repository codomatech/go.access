package main

import "fmt"

import parsers "github.com/codomatech/go.access/parsers"
import analyzers "github.com/codomatech/go.access/analyzers"
import output "github.com/codomatech/go.access/output"

func main() {
	fmt.Println("parsers.plugins: ", parsers.Plugins())
	fmt.Println("analyzers.plugins: ", analyzers.Plugins())
	fmt.Println("output.plugins: ", output.Plugins())
}

// code generation

//go:generate spluggy -func Parse -pkg github.com/codomatech/go.access/parsers ./parsers
//go:generate spluggy -func Analyze -pkg github.com/codomatech/go.access/analyzers ./analyzers
//go:generate spluggy -func Output -pkg github.com/codomatech/go.access/output ./output
