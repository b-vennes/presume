presume:
	go build -o presume main.go

.PHONY: generate
generate: presume
	./presume generate ${GENERATE_ARGS}

.PHONY: serve
serve: presume
	./presume serve ${SERVE_ARGS}
