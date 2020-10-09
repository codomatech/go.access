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

//go:generate spluggy -func Parse -pkg github.com/codomatech/goaccess.go/parsers ./parsers
//go:generate spluggy -func Analyze -pkg github.com/codomatech/goaccess.go/analyzers ./analyzers
//go:generate spluggy -func Output -pkg github.com/codomatech/goaccess.go/output ./output
