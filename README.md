# Presume

Generate a CV using an HTML Go template and typed XML content.

The [Go toolchain](https://go.dev/dl/) is required to build and run the application.

The `content` directory contains some examples of my CVs. The `content/cv-schema.xsd` file contains the defined XML schema.

## Run Hydration Tool

### Generate

Given an XML CV content file located in `content.xml`, an HTML CV template located in `template.html`, and an expected output at `result.html` run:

`go run main.go -c content.xml -t template.html -o result.html`

## Full PDF Builds

For my own CVs, I've set up a Makefile which runs the entire process for building some particular CVs I've developed.

### Backend CV

`make backend.pdf`

The hydrated PDF file will be saved to `./generated/backend.pdf`.

### QA Engineer CV

`make qa-engineer.pdf`

The hydrated PDF file will be saved to `./generated/qa-engineer.pdf`.
