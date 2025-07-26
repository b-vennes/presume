package main

import (
  "fmt"
  "os"
  "flag"

  "github.com/b-vennes/presume/pkg/generate"
  "github.com/b-vennes/presume/pkg/xmldecoding"
)

func main() {
  cvContent := flag.String("c", "", "CV content path")
  cvTemplate := flag.String("t", "", "CV template path")
  cvOutput := flag.String("o", "", "generated CV output")

  flag.Parse()

  errs := []string{}

  if cvContent == nil || *cvContent == "" {
    errs = append(errs, "CV content path is required.")
  }

  if cvTemplate == nil || *cvTemplate == "" {
    errs = append(errs, "CV template path is required.")
  }

  if cvOutput == nil || *cvOutput == "" {
    errs = append(errs, "Generated CV output path is required.")
  }

  if len(errs) > 0 {
    for _, e := range errs {
      fmt.Println(e)
    }
    return
  }

  result, err := xmldecoding.Decode(*cvContent)

  if err != nil {
    fmt.Println("Failed to decode XML file.", err)
  }

  output, err := os.Create(*cvOutput)

  if err != nil {
    fmt.Println("Failed to open output file.")
    fmt.Println(err)
  }

  err = generate.Resume(result, *cvTemplate, output)

  if err != nil {
    fmt.Println("Failed to generate resume.")
    fmt.Println(err)
  }
}
