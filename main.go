package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/htamakos/contran/translater"
)

const name = "contran"
const version = "0.0.1"

var revision = "HEAD"

var (
	sourceType  = flag.String("s", "eb", "Source Type")
	targetType  = flag.String("t", "compose", "Target Type")
	output      = flag.String("o", "", "Output")
	showVersion = flag.Bool("v", false, "Print the version")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("%s %s (rev: %s/%s)\n", name, version, revision, runtime.Version())
		return
	}

	inputValues, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Printf("[Error] %v", err)
		os.Exit(1)
	}

	var sourceTranslater, targetTranslater translater.Translater
	var outputWriter io.Writer

	switch *sourceType {
	case "eb":
		sourceTranslater = &translater.Eb{}
	default:
		fmt.Println("option: it must be in eb")
		os.Exit(1)
	}

	switch *targetType {
	case "compose":
		targetTranslater = &translater.Composer{}
	default:
		fmt.Println("option: it must be in compose")
		os.Exit(1)
	}

	if *output == "" {
		outputWriter = os.Stdout
	} else {
		outputWriter, err = os.OpenFile(*output, os.O_CREATE|os.O_WRONLY, 0700)
		if err != nil {
			log.Printf("[Error] %v", err)
			os.Exit(1)
		}
	}

	service := translater.NewTranslateService(
		sourceTranslater,
		targetTranslater,
		outputWriter,
	)

	err = service.Translate(inputValues)
	if err != nil {
		log.Printf("[Error] unexpected behaivor occurs: %v", err)
		os.Exit(1)
	}
}
