package main

import (
	"html/template"
	"net/http"
	"os"
)

func handleRequests(w http.ResponseWriter, r *http.Request) {
	// Read the content of the HTML file using os.ReadFile
	htmlContent, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Error reading HTML file", http.StatusInternalServerError)
		return
	}

	// Parse the HTML template with the content read from the file
	tmpl, err := template.New("index").Parse(string(htmlContent))
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Execute the template without passing any data
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Start the web server on port 8081
	http.ListenAndServe(":8081", http.HandlerFunc(handleRequests))
}
