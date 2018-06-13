package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/knsh14/udemy/executor"
)

var (
	httpPort string
	mode     string
)

func init() {
	flag.StringVar(&httpPort, "http", "", "port to work as http server mode")
	flag.StringVar(&mode, "mode", "current", "which to open PR current branch or all branch in local")
}

func main() {
	flag.Parse()
	switch mode {
	case "current":
		err := executor.OpenCurrentBranch()
		if err != nil {
			log.Fatal(err)
		}
	case "all":
		err := executor.OpenAllBranches()
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("wrong mode ")
	}
}
