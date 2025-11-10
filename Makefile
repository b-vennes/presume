content_dir="content"
templates_dir="templates"
generated_dir="generated"

.PHONY: generated.css
generated.css:
	deno task tailwind --input ./templates/styles.css --output ./$(generated_dir)/generated.css

.PHONY: backend.html
backend.html: generated.css
	go run main.go generate -c ./$(content_dir)/backend.xml -t ./$(templates_dir)/colorful.html -o ./$(generated_dir)/backend.html

.PHONY: qa-engineer.html
qa-engineer.html: generated.css
	go run main.go generate -c ./$(content_dir)/qa-engineer.xml -t ./$(templates_dir)/colorful.html -o ./$(generated_dir)/qa-engineer.html

.PHONY: backend.pdf
backend.pdf: backend.html
	deno run -A GeneratePDF.ts --statics ./$(generated_dir) --cv backend.html --output ./$(generated_dir)/backend.pdf

.PHONY: qa-engineer.pdf
qa-engineer.pdf: qa-engineer.html
	deno run -A GeneratePDF.ts --statics ./$(generated_dir) --cv qa-engineer.html --output ./$(generated_dir)/qa-engineer.pdf

.PHONY: customer-service.html
customer-service.html: generated.css
	go run main.go generate -c ./$(content_dir)/customer-service.xml -t ./$(templates_dir)/colorful.html -o ./$(generated_dir)/customer-service.html

.PHONY: customer-service.pdf
customer-service.pdf: customer-service.html
	deno run -A GeneratePDF.ts --statics ./$(generated_dir) --cv customer-service.html --output ./$(generated_dir)/customer-service.pdf
