# Presume

Generate a CV using an HTML Go template and typed XML content.

The [Go toolchain](https://go.dev/dl/) is required to build and run the application.

The `cv-content` directory contains some examples of my CVs. The `cv-content/cv-schema.xsd` file contains the defined XML schema.

## Build

`make presume`

## Run

### Generate

First, set arguments to the generate command in the `GENERATE_ARGS` environment variable.

`export GENERATE_ARGS="-c ./cv-content/general.xml -t ./cv-template/colorful.html -o ./generated/general.html"`

Then,

`make generate`

### Serve

First, set arguments to the serve command in the `SERVE_ARGS` environment variable.

`export SERVE_ARGS="-dir ./generated"`

Then,

`make serve`
