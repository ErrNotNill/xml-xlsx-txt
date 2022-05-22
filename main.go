package main

import (
	"fmt"
	"io/ioutil"
	"xml-txt/pkg"
)

//1 - FindAllTypes (added all files)
//2 - ParseXml (added Article array with articles)
//3 -

func main() {
	if err := pkg.FindAllTypesInDir(); err != nil {
		fmt.Println(err.Error())
	}

	/*check, _ := os.ReadFile("./txt/Booking.txt")
	fmt.Println(string(check))*/

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
