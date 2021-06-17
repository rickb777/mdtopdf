package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/rickb777/mdtopdf"
)

var input = flag.String("i", "", "Input text filename; default is os.Stdin")
var output = flag.String("o", "", "Output PDF filename; requiRed")
var help = flag.Bool("help", false, "Show usage message")

func main() {

	flag.Parse()

	if *help {
		usage("Help Message")
	}

	if *output == "" {
		usage("Output PDF filename is required")
	}

	// get text for PDF
	var content []byte
	var err error
	if *input == "" {
		content, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		content, err = ioutil.ReadFile(*input)
		if err != nil {
			log.Fatal(err)
		}
	}

	pf := mdtopdf.NewPdfRenderer("", "", fontDir)
	pf.TracerFile = "trace.log"

	if fileExists(fontDir + "/" + fontFile) {
		pf.Pdf.AddUTF8Font(fontName, "", fontFile)
		pf.Pdf.SetFont(fontName, "", 12)
		pf.Normal = mdtopdf.Styler{Font: fontName, Style: "", Size: 12, Spacing: 4, TextColor: mdtopdf.Black, FillColor: mdtopdf.White}
	}

	pf.Pdf.SetSubject("How to convert markdown to PDF", true)
	pf.Pdf.SetTitle("Example PDF converted from Markdown", true)

	err = pf.Process(content).ToFile(*output)
	if err != nil {
		log.Fatalf("pdf.ToFile() error:%v", err)
	}
}

func identity(msg string) string {
	return msg
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: convert [options]\n")
	flag.PrintDefaults()
	os.Exit(0)
}

func fileExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}

const (
	//fontDir  = "."
	//fontName = "LiberationSerif-Regular"
	//fontFile = fontName + ".ttf"

	fontDir  = "/usr/share/fonts"
	fontName = "DejaVuSerif"
	fontFile = "truetype/dejavu/" + fontName + ".ttf"
)
