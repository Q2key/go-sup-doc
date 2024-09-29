package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Document struct {
	XMLName xml.Name `xml:"w:document"`
	XMLNS   string   `xml:"xmlns:w,attr"`
	Body    Body     `xml:"w:body"`
}

func NewDocument() *Document {
	return &Document{
		XMLNS: "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
		Body:  *NewBody(),
	}
}

type Body struct {
	XMLName    xml.Name `xml:"w:body"`
	Paragraphs *[]Paragraph
}

func NewBody() *Body {
	return &Body{
		Paragraphs: nil,
	}
}

func (b *Body) AppendParagraph(paragraph Paragraph) {
	if b.Paragraphs == nil {
		b.Paragraphs = &[]Paragraph{}
	}
	*b.Paragraphs = append(*b.Paragraphs, paragraph)
}

type Paragraph struct {
	XMLName xml.Name `xml:"w:p"`
	Regular *RegularText
}

type RegularText struct {
	XMLName   xml.Name `xml:"w:r"`
	TextNodes *[]Text
}

type Text struct {
	XMLName xml.Name `xml:"w:t"`
	Text    string   `xml:",chardata"`
}

func (p *Paragraph) AppendRegularText(str string) {
	if p.Regular == nil {
		p.Regular = &RegularText{}
	}

	if p.Regular.TextNodes == nil {
		p.Regular.TextNodes = &[]Text{}
	}

	*p.Regular.TextNodes = append(*p.Regular.TextNodes, Text{Text: str})
}

/*
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:body>
    <w:p>
      <w:r>
        <w:t>Hello, World!</w:t>
      </w:r>
    </w:p>
  </w:body>
</w:document>
*/

func main() {

	doc := NewDocument()

	p1 := Paragraph{}

	p1.AppendRegularText("Hello World")
	p1.AppendRegularText("Hello World2")

	doc.Body.AppendParagraph(p1)
	out, err := xml.MarshalIndent(doc, " ", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))

	file, err := os.Create("test.docx")
	if err != nil {
		log.Fatal(err)
	}

	_, _ = file.Write([]byte(string(out)))
}
