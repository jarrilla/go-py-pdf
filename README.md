# PDF to Text Demo

This is a demo of a microservice that converts a PDF file to text.

## Usage
Start the Python microservice:
```
cd pdf2txt
uv install
uv run
```

Then, in another terminal, run the Go client:
```bash
go run main.go <path_to_pdf_file>
```