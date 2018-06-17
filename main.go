package main

import (
	"flag"
	"log"

	"github.com/knsh14/open-pull-request/executor"
)

var (
	mode string
)

func init() {
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
		log.Fatal("wrong mode")
	}
}
