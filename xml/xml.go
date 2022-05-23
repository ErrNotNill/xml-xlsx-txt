package xml

import (
	"fmt"
	"github.com/beevik/etree"
	"xml-txt/pkg"
)

//var TryImport string

var Article []string //THERE ALL Articles
var UserArgsArray []string

func ParseXml(scan string) {

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(pkg.XmlFile); err != nil {
		panic(err)
	}
	root := doc.SelectElement("data")
	fmt.Println("ROOT element:", root.Tag)

	//todo NEXT WORK HERE

	for _, book := range root.SelectElements("product") {
		//fmt.Println("CHILD element:", book.Tag)
	//	if title := book.SelectElement("article"); title != nil { //todo return this
			//lang := title.SelectAttrValue("lang", "unknown")
			//fmt.Printf("  Article: %s \n", title.Text())
		//	Article = append(Article, title.Text()) //todo look at this
		if userargs := book.SelectElement(scan);userargs != nil{ //todo User Args
			Article = append(Article, userargs.Text())
			fmt.Println(userargs.Text())
		}

	//UserArgsArray = append(UserArgsArray,userargs.Text())
}
		/*for _, attr := range book.Attr {
			fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
		}
	}
	for _,v := range Article {
		if scan == v {
			fmt.Printf("Article: %v\n", v)
		}

	}*/
}
