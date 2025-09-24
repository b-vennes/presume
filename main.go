package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/b-vennes/presume/pkg/generate"
	"github.com/b-vennes/presume/pkg/xmldecoding"
)

/* Safely retrieves the value at the given index with the given fallback alternative. */
func safeGet[A any](slice []A, index int, fallback A) A {
	if len(slice) <= index {
		return fallback
	} else {
		return slice[index]
	}
}

/* Safely slices the given slice starting at the given index to the end.  If the index is beyond the slice then it returns an empty slice. */
func safeSlice[A any](slice []A, start int) []A {
	if len(slice) <= start {
		return []A{}
	} else {
		return slice[start:]
	}
}

/* Creates a static server at the given directory path. */
func runServe(directory string) {
	fs := http.FileServer(http.Dir(directory))

	log.Print("Listening on :5050...")

	err := http.ListenAndServe(":5050", fs)

	if err != nil {
		log.Fatal(err)
	}
}

/* Parses serve command arguments.  Returns the directory and a possible list of errors. */
func parseServe(args []string) (string, []string) {
	cmd := flag.NewFlagSet("serve", flag.ExitOnError)

	parsedDirectory := cmd.String("dir", "", "Path to directory to serve.")

	err := cmd.Parse(args)

	if err != nil {
		log.Fatalln("Failed to parse arguments: " + err.Error())
	}

	dir := ""

	errs := []string{}

	if parsedDirectory == nil || *parsedDirectory == "" {
		errs = append(errs, "No 'dir' argument provided.")
	}

	dir = *parsedDirectory

	return dir, errs
}

/* Runs the CV generator using the given options. */
func runGenerate(cvContent string, cvTemplate string, cvOutput string) {
	result, err := xmldecoding.Decode(cvContent)

	if err != nil {
		log.Println("Failed to decode XML file.", err)
	}

	output, err := os.Create(cvOutput)

	if err != nil {
		log.Println("Failed to open output file.")
		log.Println(err)
	}

	err = generate.Resume(result, cvTemplate, output)

	if err != nil {
		log.Println("Failed to generate resume.")
		log.Fatalln(err)
	}
}

/* Parses arguments to the generate command.  Returns the content path, template path, output path, and a list of errors. */
func parseGenerate(args []string) (string, string, string, []string) {
	cmd := flag.NewFlagSet("generate", flag.ExitOnError)

	if cmd == nil {
		log.Fatal("Failed to create 'generate' command processor.")
	}

	parsedCVContent := cmd.String("c", "", "CV content path")
	parsedCVTemplate := cmd.String("t", "", "CV template path")
	parsedCVOutput := cmd.String("o", "", "generated CV output")
	cmd.Parse(args)

	content := ""
	template := ""
	output := ""
	errs := []string{}

	if parsedCVContent == nil || *parsedCVContent == "" {
		errs = append(errs, "CV content path is required (-c).")
	} else {
		content = *parsedCVContent
	}

	if parsedCVTemplate == nil || *parsedCVTemplate == "" {
		errs = append(errs, "CV template path is required (-t).")
	} else {
		template = *parsedCVTemplate
	}

	if parsedCVOutput == nil || *parsedCVOutput == "" {
		errs = append(errs, "Generated CV output path is required (-o).")
	} else {
		output = *parsedCVOutput
	}

	return content, template, output, errs
}

func main() {
	switch safeGet(os.Args, 1, "") {
	case "serve":
		dir, errs := parseServe(safeSlice(os.Args, 2))

		if len(errs) > 0 {
			for _, err := range errs {
				log.Println(err)
			}

			log.Fatalln("Run `presume help` for additional info.")
		}

		runServe(dir)
		os.Exit(0)
	case "generate":
		content, template, output, errs := parseGenerate(safeSlice(os.Args, 2))

		if len(errs) > 0 {
			for _, err := range errs {
				log.Println(err)
			}
			log.Fatalln("Run `presume help` for additional info.")
		}

		runGenerate(content, template, output)
		os.Exit(0)
	case "help":
		log.Fatalln("Generate or serve CVs using 'generate' or 'serve' commands.")
	default:
		log.Fatalln("Expected 'serve', 'generate', or 'help' command.")
	}
}
