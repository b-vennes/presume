/*
Package generate provides functions for generating CVs from models.
*/
package generate

import (
  "html/template"
  "io"

  "github.com/b-vennes/presume/pkg/models"
)

func Resume(data *models.ResumeContentData, templatePath string, output io.Writer) error {
  t, err := template.ParseFiles(templatePath)

  if err != nil {
    return err
  }

  resumeView := models.MakeResumeContentView(*data)

  err = t.Execute(output, resumeView)

  return err
}
