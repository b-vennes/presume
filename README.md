# Presume

Generate a CV using an HTML Go template and typed XML data.

The [Go toolchain](https://go.dev/dl/) is required to build and run the application.

## Build

`go build .`

## Run

```bash
go run main.go \
  -c [content-xml-path] \
  -t [view-template-html-path] \
  -o [output-file-html-path]
```
