package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rickb777/mdtopdf"
)

var (
	input, output string
)

func main() {
	flag.StringVar(&input, "i", "", "Input text filename; default is os.Stdin")
	flag.StringVar(&output, "o", "", "Output PDF filename; required")
	var help = flag.Bool("help", false, "Show usage message")

	flag.Parse()

	if *help {
		usage("Help Message")
	}

	if output == "" {
		if input == "" {
			usage("Output PDF filename is required")
		}
		output = strings.TrimSuffix(input, filepath.Ext(input)) + ".pdf"
	}

	// get text for PDF
	var content []byte
	var err error
	if input == "" {
		content, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		content, err = ioutil.ReadFile(input)
		if err != nil {
			log.Fatal(err)
		}
	}

	pf := mdtopdf.NewPdfRenderer("", "", fontDir)
	pf.TracerFile = "trace.log"

	if fileExists(fontDir + "/" + fontFile) {
		fmt.Println(fontDir + "/" + fontFile)
		pf.Pdf.AddUTF8Font(fontName, "", fontFile)
		pf.Pdf.SetFont(fontName, "", 10)
		if pf.Pdf.Err() {
			fmt.Println(pf.Pdf.Error())
			os.Exit(1)
		}
		pf.Normal = mdtopdf.Styler{Font: fontName, Style: "", Size: 10, Spacing: 4, TextColor: mdtopdf.Black, FillColor: mdtopdf.White}
	}

	pf.Pdf.SetSubject("How to convert markdown to PDF", true)
	pf.Pdf.SetTitle("Example PDF converted from Markdown", true)

	err = pf.Process(content).ToFile(output)
	if err != nil {
		log.Fatalf("pdf.ToFile() error:%v", err)
	}
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: convert [options]\n")
	flag.PrintDefaults()
	os.Exit(1)
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
