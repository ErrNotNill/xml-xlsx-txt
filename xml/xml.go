package xml

import (
	"fmt"
	"github.com/beevik/etree"
	"xml-txt/pkg"
)

const Takeme = "Hi"

var TryImport string
var Article []string //THERE ALL Articles

func ParseXml() {
	if err := pkg.FindAllTypesInDir(); err != nil {
		fmt.Println(err.Error())
	}

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(pkg.XmlFile); err != nil {
		panic(err)
	}
	root := doc.SelectElement("data")
	fmt.Println("ROOT element:", root.Tag)

	//todo NEXT WORK HERE

	for _, book := range root.SelectElements("product") {
		//fmt.Println("CHILD element:", book.Tag)
		if title := book.SelectElement("article"); title != nil {
			//lang := title.SelectAttrValue("lang", "unknown")
			//fmt.Printf("  Article: %s \n", title.Text())
			Article = append(Article, title.Text())
		}
		for _, attr := range book.Attr {
			fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
		}
	}
	for i := range Article {
		fmt.Printf("Article: %v\n", Article[i])
	}
}
