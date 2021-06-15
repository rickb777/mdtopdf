package main

// This example manipulates gofpdf directly and shows off how UTF-8 characters are handled.

import (
	"fmt"
	"github.com/jung-kurt/gofpdf/v2"
)

func main() {
	// on Ubuntu, DejaVuSerif is
	pdf := gofpdf.New("P", "mm", "A4", "/usr/share/fonts")
	pdf.AddUTF8Font("DejaVuSerif", "", "truetype/dejavu/DejaVuSerif.ttf")
	pdf.AddPage()
	pdf.SetFont("DejaVuSerif", "", 16)
	pdf.Cell(40, 10, "Hello, Félicité Sørina Siân Øyvind François.  «µ» !¡ 1¹ 2² 3³")
	err := pdf.OutputFileAndClose("utf8-example.pdf")
	fmt.Println(err)
}
