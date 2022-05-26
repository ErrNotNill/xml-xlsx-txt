package main

import (
	"fmt"
	"io/ioutil"
	"xml-txt/pkg"
	"xml-txt/xls"
	"xml-txt/xml"
)

func main() {

	if err := pkg.FindAllTypesInDir(); err != nil {
		fmt.Println(err.Error())
	}
	xml.ParseXml("article")  //opts what you need find
	xls.ParseExcel("033536") //todo method for user args

	//xls.GetColumnsExcel()
	//xls.ParseExcel("article")
	//txt.GetArgsFromTxt()

}

func ReadXmlFile() error {
	file, err := ioutil.ReadFile(pkg.ExcelFile)
	if err != nil {
		fmt.Println("cant read")
	}
	fmt.Println(string(file))
	return err
}
func ReadTxtFile() error {
	file, err := ioutil.ReadFile(pkg.TxtFile)
	if err != nil {
		fmt.Println("cant read")
	}
	fmt.Println(string(file))
	return err
}
func ReadExcelFile() error {
	file, err := ioutil.ReadFile(pkg.ExcelFile)
	if err != nil {
		fmt.Println("cant read")
	}
	fmt.Println(string(file))
	return err
}
