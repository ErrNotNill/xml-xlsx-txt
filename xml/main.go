package main

import (
	"fmt"
	"github.com/beevik/etree"
)

func main() {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("./xml/Book.xml"); err != nil {
		panic(err)
	}
	root := doc.SelectElement("CATALOG")
	fmt.Println("ROOT element:", root.Tag)

	for _, book := range root.SelectElements("CD") {
		//fmt.Println("CHILD element:", book.Tag)
		if title := book.SelectElement("TITLE"); title != nil {
			lang := title.SelectAttrValue("lang", "unknown")
			fmt.Printf("  TITLE: %s (%s)\n", title.Text(), lang)
		}
		for _, attr := range book.Attr {
			fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
		}
	}
}
