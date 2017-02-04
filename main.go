package main

import (
	"fmt"
	"log"
)

import (
	"github.com/sternix/eolian_go/eolian"
//	"./eolian"
)

func main() {
	defer eolian.Shutdown()

	// some tests

	if !eolian.ScanSystemDirectory() {
		log.Print("Error : ScanSystemDirectory")
	}

	if !eolian.ParseAllEoFiles() {
		log.Println("Error: ParseAllEoFiles")
	}

	if !eolian.ParseAllEotFiles() {
		log.Println("Error : ParseAllEotFiles")
	}


	for _, f := range eolian.AllEoFiles() {
		fmt.Println(f)
	}

	if !eolian.ValidateDatabase(true) {
		log.Println("Error : ValidateDatabase")
	}


	for _, c := range eolian.AllClasses() {
		if c.Type() != eolian.ClassTypeRegular {
			continue
		}
		//fmt.Println(c.Name())
		fmt.Println("FullName :" , c.FullName())
		fmt.Println("Data Type:" , c.DataType())
		for _, ns := range c.Namespaces() {
			fmt.Println("Namespace :", ns)
		}
		//fmt.Println(c.LegacyPrefix())
		/*
		for _, f := range c.Functions(eolian.FunctionTypeProperty) {
			if f.IsCOnly() {
				continue
			}
			fmt.Println("Property:" , f.Name())
			fmt.Println("FullCname:" , f.FullCName(eolian.FunctionTypeProperty,true))
			fmt.Println("FullCname:" , f.FullCName(eolian.FunctionTypeProperty,false))
		}

		for _, f := range c.Functions(eolian.FunctionTypePropSet) {
			if f.IsCOnly() {
				continue
			}
			fmt.Println("PropSet:" , f.Name())
			fmt.Println("FullCname:" , f.FullCName(eolian.FunctionTypePropSet,true))
			fmt.Println("FullCname:" , f.FullCName(eolian.FunctionTypePropSet,false))
		}


		for _, f := range c.Functions(eolian.FunctionTypePropGet) {
			fmt.Println("PropGet:" , f.Name())
		}

		for _, f := range c.Functions(eolian.FunctionTypeMethod) {
			fmt.Println("PropMethod:" , f.Name())
		}
		*/

		for _, f := range c.Functions(eolian.FunctionTypeUnresolved) {
			fmt.Println("PropUnresolved:" , f.Name())
		}

		for _ , i := range c.Inherits() {
			fmt.Println("Inherits :" , i)
		}

		for _ , i := range c.Implements() {
			fmt.Println("Implement:", i.FullName())
		}

		for _ , e := range c.Events() {
			fmt.Println("Event:" , e.Name())
		}

		for _ , co := range c.Constructors() {
			fmt.Println("Constructor: " , co.FullName())
		}

		fmt.Println("C Get func : " , c.CGetFunctionName())

		fmt.Println("-----------------------------------")
	}


	for _ , d := range eolian.AllDeclarations()  {
		if d.Type() == eolian.DeclarationTypeClass {
			fmt.Println(d.Name())
		}
	}
}
