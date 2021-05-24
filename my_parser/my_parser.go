package my_parser

import (
	"github.com/davecgh/go-spew/spew"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func DemoMyParser() {
	//demoV1Dump()
	demoV2Printer()

}

func demoV1Dump() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "package main; var a = 3", parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(f)
}

func demoV2Printer() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "package main; var a = 3", parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	printer.Fprint(os.Stdout, fs, f)
}
