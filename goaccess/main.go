package main

import "fmt"

import parsers "github.com/codomatech/goaccess.go/parsers"

func main() {
	fmt.Println("parsers.plugins: ", parsers.Plugins())
}

// code generation

//go:generate ./scripts/package-plugins.sh parsers github.com/codomatech/goaccess.go/parsers func(string)common.AccessRecord Parse parsers/plugins.go
