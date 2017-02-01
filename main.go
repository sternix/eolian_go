package main

import (
	"fmt"
	"log"
)

import (
	"github.com/sternix/eolian-go/eolian"
)

func main() {
	defer eolian.Shutdown()

	// some tests

	if !eolian.ScanSystemDirectory() {
		log.Print("Error : ScanSystemDirectory")
	}

	if !eolian.ValidateDatabase(true) {
		log.Println("Error : ValidateDatabase")
	}

	for _, f := range eolian.AllEoFiles() {
		fmt.Println(f)
	}

	if !eolian.ParseAllEoFiles() {
		log.Println("Error: ParseAllEoFiles")
	}

	if !eolian.ParseAllEotFiles() {
		log.Println("Error : ParseAllEotFiles")
	}

	for _, c := range eolian.AllClasses() {
		//fmt.Println(c.Name())
		fmt.Println(c.FullName())
		for _, ns := range c.Namespaces() {
			fmt.Println("-", ns)
		}
		//fmt.Println(c.LegacyPrefix())
		for _, f := range c.Functions(eolian.FunctionTypeProperty) {
			fmt.Println(f.Name())
		}
	}
}
