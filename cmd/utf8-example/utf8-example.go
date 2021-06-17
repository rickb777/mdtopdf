package main

// This example manipulates gofpdf directly and shows off how UTF-8 characters are handled.

import (
	"fmt"
	"github.com/phpdave11/gofpdf"
)

func main() {
	// on Ubuntu, DejaVuSerif is expected here
	pdf := gofpdf.New("P", "pt", "A4", "/usr/share/fonts")
	pdf.AddUTF8Font("DejaVuSerif", "", "truetype/dejavu/DejaVuSerif.ttf")
	pdf.AddPage()
	for i := 1; i < len(txt); i += 2 {
		pdf.SetFont("DejaVuSerif", "", 14)
		pdf.CellFormat(0, 32, txt[i-1], "1", 1, "L", false, 0, "")
		pdf.SetFont("DejaVuSerif", "", 12)
		pdf.CellFormat(0, 20, txt[i], "", 1, "L", false, 0, "")
	}
	err := pdf.OutputFileAndClose("utf8-example.pdf")
	fmt.Println(err)
}

var txt = []string{
	`Russian ‘Молоко и творог’`,
	`- means “milk and cottage cheese”`,
	`Welsh ‘Côf a lithr, llythyrau a geidw’`,
	`- means “memory slips, letters remain”`,
	`Danish ‘Så er den ged barberet’`,
	`- means “now that goat has been shaved” (the work is done)`,
	`Icelandic ‘Árinni kennir illur ræðari‘`,
	`- means “a bad rower blames his oars”`,
	`Greek ‘Όταν λείπει η γάτα, χορεύουν τα ποντίκια’`,
	`- means “when the cat’s away, the mice dance”`,
	//`Japanese ‘猿も木から落ちる’`, -- does not render with this font
	//`- means “even monkeys fall from trees” (everyone makes mistakes)`,
}
