package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"salsadigitalauorg/shipshape/pkg/shipshape"
)

var projectDir string
var checksFile string
var outputFormat string

func main() {
	parseFlags()
	parseArgs()
	validateOutputFormat(&outputFormat)

	cfg, err := shipshape.ReadAndParseConfig(projectDir, "shipshape.yml")
	if err != nil {
		log.Fatal(err)
	}
	cfg.Init()
	r := cfg.RunChecks()

	if outputFormat == "json" {
		data, err := json.Marshal(r)
		if err != nil {
			log.Fatalf("Unable to convert result to json: %+v\n", err)
		}
		fmt.Println(string(data))
	} else if outputFormat == "table" {
		fmt.Printf("Result: %#v\n", r)
	}
}

func parseFlags() {
	flag.StringVar(&checksFile, "checks-file", "shipshape.yml", "Path to the file containing the checks")
	flag.StringVar(&outputFormat, "output", "table", "Output format (table|json); default is table.")
	flag.Parse()
}

func parseArgs() {
	args := flag.Args()
	if len(args) > 1 {
		log.Fatalf("Max 1 argument expected, got '%+v'\n", args)
	} else if len(args) == 1 {
		projectDir = args[0]
	}
}

func validateOutputFormat(of *string) {
	if *of != "json" && *of != "table" {
		log.Fatal("Invalid output format; needs to be 'table' or 'json'.")
	}
}
