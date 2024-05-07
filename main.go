package main

import (
	//"github.com/gin-gonic/gin"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc(`/`, GetPdfHandler)

	http.ListenAndServe(":8080", r)

}

func GetPdfHandler(w http.ResponseWriter, r *http.Request) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Create a table
	pdf.SetFont("Arial", "", 12)
	pdf.SetFillColor(255, 255, 255) // White fill color
	pdf.CellFormat(190, 7, "Table Title", "0", 1, "C", false, 0, "")
	pdf.Ln(10)

	// Define table columns
	colWidths := []float64{40, 50, 50, 50}

	// Create table cells
	rows := [][]string{
		{"Name", "Age", "Email", "Country"},
		{"John Doe", "31", "john@example.com", "USA"},
		{"Jane Smith", "25", "jane@example.com", "UK"},
	}

	// Create table
	pdf.SetFont("Arial", "", 12)
	for _, row := range rows {
		for _, cell := range row {
			pdf.CellFormat(colWidths[0], 7, cell, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
	}
	pdf.ImageOptions("image.jpg", 55, 55, 100, 0, false, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "")

	pdf.OutputFileAndClose("table.pdf")

	http.ServeFile(w, r, "table.pdf")
}
