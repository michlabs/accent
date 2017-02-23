package main 

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/michlabs/accent"
)

func main() {
	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	inputFlag := removeCmd.String("i", "", "required, path to your input file")
	outputFlag := removeCmd.String("o", "", "required, path to your output file")

	if len(os.Args) < 2 {
		log.Println("Error: must specify a command")
		PrintHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "remove":
		removeCmd.Parse(os.Args[2:])
	case "help":
		PrintHelp()
		os.Exit(0)
	default:
		log.Println("Error: invalide command")
		PrintHelp()
		os.Exit(1)
	}

	if removeCmd.Parsed() {
		if *inputFlag == "" {
			log.Println("Error: input file is required")
			PrintHelp()
			os.Exit(1)
		}
		if *outputFlag == "" {
			log.Println("Error: output file is required")
			PrintHelp()
			os.Exit(1)
		}

		if err := accent.RemoveFromFile(*inputFlag, *outputFlag); err != nil {
			log.Println("Error:", err)
			PrintHelp()
			os.Exit(1)
		}
	}

	log.Println("Done!")
}

func PrintHelp() {
	var help string = `
Description:
accent CLI removes Vietnamese accents from a text file and writes result to output file.

Usage: accent <command> <option>
Available commands and corresponding options:
	remove 
	  -i string
	    	required, path to your input file
	  -o string
	    	required, path to your output file
	help
`
	fmt.Println(help)
}