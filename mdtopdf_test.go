/*
 * Markdown to PDF Converter
 * Available at http://github.com/rickb777/mdtopdf
 *
 * Copyright © 2018 Cecil New <cecil.new@gmail.com>.
 * Distributed under the MIT License.
 * See README.md for details.
 *
 * Dependencies
 * This package depends on two other packages:
 *
 * Blackfriday Markdown Processor
 *   Available at http://github.com/russross/blackfriday
 *
 * gofpdf - a PDF document generator with high level support for
 *   text, drawing and images.
 *   Available at https://github.com/jung-kurt/gofpdf/v2
 */

package mdtopdf

import (
	bf "github.com/russross/blackfriday/v2"
	"io/ioutil"
	"path"
	"strings"
	"testing"
)

// compare the results visually against (e.g.) https://md2pdf.netlify.app/

func testit(name string, t *testing.T) {
	inputDir := "./testdata/"
	input := path.Join(inputDir, name)

	base := strings.TrimSuffix(path.Base(input), ".md")
	pdfFile := path.Join(inputDir, base) + ".pdf"
	htmlFile := path.Join(inputDir, base) + ".html"

	markdown, err := ioutil.ReadFile(input)
	if err != nil {
		t.Fatalf("%v:%v", input, err)
	}

	err = ioutil.WriteFile(htmlFile, bf.Run(markdown), 0644)
	if err != nil {
		t.Fatalf("%v:%v", htmlFile, err)
	}

	r := NewPdfRenderer("", "", "")
	r.TracerFile = path.Join(inputDir, base) + ".log"

	err = r.Process(markdown).ToFile(pdfFile)
	if err != nil {
		t.Fatalf("%v:%v", pdfFile, err)
	}
}

func TestTables(t *testing.T) {
	testit("Tables.md", t)
}

func TestMarkdownDocumenationBasic(t *testing.T) {
	testit("Markdown Documentation - Basics.md", t)
}

func TestMarkdownDocumenationSyntax(t *testing.T) {
	testit("Markdown Documentation - Syntax.md", t)
}

func TestImage(t *testing.T) {
	testit("Image.md", t)
}

func TestAutoLinks(t *testing.T) {
	testit("Auto links.md", t)
}

func TestAmpersandEncoding(t *testing.T) {
	testit("Amps and angle encoding.md", t)
}

func TestInlineLinks(t *testing.T) {
	testit("Links, inline style.md", t)
}

func TestLists(t *testing.T) {
	testit("Ordered and unordered lists.md", t)
}

func TestStringEmph(t *testing.T) {
	testit("Strong and em together.md", t)
}

func TestTabs(t *testing.T) {
	testit("Tabs.md", t)
}

func TestBackslashEscapes(t *testing.T) {
	testit("Backslash escapes.md", t)
}

func TestBackquotes(t *testing.T) {
	testit("Blockquotes with code blocks.md", t)
}

func TestCodeBlocks(t *testing.T) {
	testit("Code Blocks.md", t)
}

func TestCodeSpans(t *testing.T) {
	testit("Code Spans.md", t)
}

func TestHardWrappedPara2(t *testing.T) {
	testit("Hard-wrapped paragraphs with list-like lines.md", t)
}

func TestHorizontalRules(t *testing.T) {
	testit("Horizontal rules.md", t)
}

func TestInlineHtmlSimple(t *testing.T) {
	testit("Inline HTML (Simple).md", t)
}

func TestInlineHtmlAdvanced(t *testing.T) {
	testit("Inline HTML (Advanced).md", t)
}

func TestInlineHtmlComments(t *testing.T) {
	testit("Inline HTML comments.md", t)
}

func TestTitleWithQuotes(t *testing.T) {
	testit("Literal quotes in titles.md", t)
}

func TestNestedBlockquotes(t *testing.T) {
	testit("Nested blockquotes.md", t)
}

func TestLinksReference(t *testing.T) {
	testit("Links, reference style.md", t)
}

func TestLinksShortcut(t *testing.T) {
	testit("Links, shortcut references.md", t)
}

func TestTidyness(t *testing.T) {
	testit("Tidyness.md", t)
}
