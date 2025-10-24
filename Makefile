content_dir="content"
templates_dir="templates"
generated_dir="generated"

generated.css:
	deno task tailwind --input ./templates/styles.css --output ./$(generated_dir)/generated.css

backend.html: generated.css
	go run main.go generate -c ./$(content_dir)/backend.xml -t ./$(templates_dir)/colorful.html -o ./$(generated_dir)/backend.html

qa-engineer.html: generated.css
	go run main.go generate -c ./$(content_dir)/qa-engineer.xml -t ./$(templates_dir)/colorful.html -o ./$(generated_dir)/qa-engineer.html

backend.pdf: backend.html
	deno run -A GeneratePDF.ts --statics ./$(generated_dir) --cv backend.html --output ./$(generated_dir)/backend.pdf

qa-engineer.pdf: qa-engineer.html
	deno run -A GeneratePDF.ts --statics ./$(generated_dir) --cv qa-engineer.html --output ./$(generated_dir)/qa-engineer.pdf
