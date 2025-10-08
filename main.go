// Package main contains the CLI to generate and serve resumes.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/b-vennes/presume/pkg/generate"
	"github.com/b-vennes/presume/pkg/xmldecoding"
)

// Safely retrieves the value at the given index with the given fallback
// alternative.
func safeGet[A any](slice []A, index int, fallback A) A {
	if len(slice) <= index {
		return fallback
	} else {
		return slice[index]
	}
}

// Safely slices the given slice starting at the given index to the end.
// If the index is beyond the slice then it returns an empty slice.
func safeSlice[A any](slice []A, start int) []A {
	if len(slice) <= start {
		return []A{}
	} else {
		return slice[start:]
	}
}

// Checks if help was requested by the user.
func helpRequested(args []string) bool {
	return len(args) > 0 && args[0] == "help"
}

type ServeArgs struct {
	// The directory of files to serve.
	Directory string
	// The port to serve files on.
	Port string
}

// Creates a static server at the given directory path.
func runServe(args ServeArgs) int {
	log.Println("Serving", args.Directory)

	fs := http.FileServer(http.Dir(args.Directory))

	log.Println("Listening on", ":"+args.Port)

	err := http.ListenAndServe(":"+args.Port, fs)

	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

const SERVE_HELP = "presume serve -dir <path to directory to serve over HTTP> -p <port : defaults to 5050>\n" +
	"Serves all the files contained in the given directory over HTTP on port 5050."

// Parses serve command arguments.
// Returns a tuple of the serve args and a possible list of errors.
func parseServe(args []string) (*ServeArgs, []string) {
	if helpRequested(args) {
		errs := []string{SERVE_HELP}
		return nil, errs
	}

	cmd := flag.NewFlagSet("serve", flag.ExitOnError)

	parsedDirectory := cmd.String("dir", "", "Path to directory to serve.")
	parsedPort := cmd.Int("p", 5050, "Port to serve files on.")

	err := cmd.Parse(args)

	if err != nil {
		fmt.Println("Failed to parse arguments: " + err.Error())
		os.Exit(1)
	}

	errs := []string{}

	if parsedDirectory == nil || *parsedDirectory == "" {
		errs = append(errs, "No 'dir' argument provided.")
	}

	if len(errs) > 0 {
		return nil, errs
	}

	return &ServeArgs{
			Directory: *parsedDirectory,
			Port:      strconv.Itoa(*parsedPort),
		},
		nil
}

type GenerateArgs struct {
	ContentPath  string
	TemplatePath string
	OutputPath   string
}

// Runs the CV generator using the given options.
func runGenerate(args GenerateArgs) int {
	result, err := xmldecoding.Decode(args.ContentPath)

	if err != nil {
		log.Println("Failed to decode XML file.", err)
		return 1
	}

	output, err := os.Create(args.OutputPath)

	if err != nil {
		log.Println("Failed to open output file.")
		log.Println(err)
		return 1
	}

	err = generate.Resume(result, args.TemplatePath, output)

	if err != nil {
		log.Println("Failed to generate resume.")
		log.Println(err)
		return 1
	}

	return 0
}

const GENERATE_HELP = "presume generate -c <CV content path> -t <CV template path> -o <generated output path>"

// Parses arguments to the generate command.
// Returns a tuple of a pointer to the generated args (possible nil)
// and a list of errors (if there were any).
func parseGenerate(args []string) (*GenerateArgs, []string) {
	if helpRequested(args) {
		errs := []string{GENERATE_HELP}
		return nil, errs
	}

	cmd := flag.NewFlagSet("generate", flag.ExitOnError)

	if cmd == nil {
		err := "Failed to create 'generate' command processor."
		log.Println(err)
		return nil, []string{err}
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

	if len(errs) != 0 {
		return nil, errs
	}

	generateArgs := GenerateArgs{
		ContentPath:  content,
		TemplatePath: template,
		OutputPath:   output,
	}

	return &generateArgs, nil
}

func main() {
	exitCode := 0

	switch safeGet(os.Args, 1, "") {
	case "serve":
		serveArgs, errs := parseServe(safeSlice(os.Args, 2))

		if errs != nil {
			for _, err := range errs {
				log.Println(err)
			}

			fmt.Println("Run `presume help` for additional info.")
			exitCode = 1
			break
		}

		exitCode = runServe(*serveArgs)
	case "generate":
		generateArgs, errs := parseGenerate(safeSlice(os.Args, 2))

		if errs != nil {
			for _, err := range errs {
				fmt.Println(err)
			}
			fmt.Println("Run `presume help` for additional info.")

			exitCode = 1
			break
		}

		exitCode = runGenerate(*generateArgs)
	default:
		message :=
			"Generate or serve CVs using the 'generate' or 'serve' commands.\n" +
				"Also give 'presume generate help' or 'presume serve help' a try!"
		fmt.Println(message)
		exitCode = 1
	}

	os.Exit(exitCode)
}
