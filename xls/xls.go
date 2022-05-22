package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"xml-txt/pkg"
	"xml-txt/xml"
)

//Sheet1 = лист1, Sheet2 = лист2 и так далее

func main() {
	if err := pkg.FindAllTypesInDir(); err != nil {
		fmt.Println(err.Error())
	}
	var article = xml.Article
	fmt.Println(article)

	f, err := excelize.OpenFile(pkg.ExcelFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	/*sheet = "Остатки"
	column = "Артикул" (F++)
	for r := range column {
		f.GetCellValue("")
	}*/
	var cellArr = make([]string, 0)
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Остатки", "F3") //get all F

	/*stroka = "F0"
	strconv.Atoi()
	F = F+1++
	fmt.Println(cell)*/
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Остатки")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range article {
		fmt.Println(v)
	}
	for _, row := range rows {

		//columnA := row[len(row)-6 : len(row)-4] //todo something with this (slice out of range [-1:]
		//columnB := row[len(row)-5 : len(row)-4] //Предупреждения (2 столбец)
		//columnC := row[len(row)-4 : len(row)-3]
		/*columnB := row[len(row)-5 : len(row)-4] //Article
		columnC := row[len(row)-4 : len(row)-3] //Article
		columnD := row[len(row)-3 : len(row)-2] //Article
		columnE := row[len(row)-2 : len(row)-1] //Article*/
		//columnF := row[len(row)-1:] //Article (F 6)

		//fmt.Println(columnA)
		//fmt.Println(columnB)
		//fmt.Println(columnC)
		/*fmt.Println(columnB)
		fmt.Println(columnC)
		fmt.Println(columnD)
		fmt.Println(columnE)*/
		//fmt.Println(columnF)

		for _, colCell := range row {
			//fmt.Println(row)
			//fmt.Println(colCell)

			cellArr = append(cellArr, colCell)
			//fmt.Print(colCell, "\t")
		}
		//fmt.Println()
	}

}
