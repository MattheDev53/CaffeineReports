package main

import (
	"fmt"
	"os"
	"strings"
)

func NoImp() {
	fmt.Println("!! Not Implemented : Nothing will be done")
}

func PrintUsage() {
	argZeroList := strings.Split(os.Args[0], "/")
	filename := argZeroList[len(argZeroList)-1]
	fmt.Printf("Usage of %s:\n", filename)
	fmt.Printf("%s help : Prints this help message\n", filename)
	fmt.Printf("%s generate [data] [template] [output] : generates a report from the provided data and template\n", filename)
}

func GenerateReport(data string, template string, output string) {
	NoImp()
	fmt.Printf("Generating report from data '%s', template '%s' and output '%s'\n", data, template, output)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No arguments provided")
		PrintUsage()
		os.Exit(3)
	}
	switch os.Args[1] {
		default:
			PrintUsage()
			os.Exit(3)
		case "help":
			PrintUsage()
			os.Exit(0)
		case "generate":
			if len(os.Args) != 5 {
				fmt.Printf("Invalid number of arguments (%d provided, 3 expected)\n", len(os.Args)-2)
				os.Exit(3)
			}
			GenerateReport(os.Args[2], os.Args[3], os.Args[4])
			os.Exit(0)
	}
}
