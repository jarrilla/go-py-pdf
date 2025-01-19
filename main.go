package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Check if PDF path is provided
	if len(os.Args) < 2 {
		log.Fatal("Please provide the path to a PDF file as an argument")
	}

	pdfPath := os.Args[1]

	// Create a pipe to stream the data
	pipeReader, pipeWriter := io.Pipe()

	// Create the request with the pipe reader
	req, err := http.NewRequest("POST", "http://localhost:5000/parse_pdf", pipeReader)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/pdf")

	// Start a goroutine to write the file to the pipe
	go func() {
		file, err := os.Open(pdfPath)
		if err != nil {
			pipeWriter.CloseWithError(err)
			return
		}
		defer file.Close()

		_, err = io.Copy(pipeWriter, file)
		pipeWriter.CloseWithError(err) // Close with error if any, or nil
	}()

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	fmt.Println("Response from microservice:", string(body))
}
