presume:
	go build -o presume main.go

.PHONY: generate
generate:
	go run main.go generate ${GENERATE_ARGS}

.PHONY: serve
serve:
	go run main.go serve ${SERVE_ARGS}
