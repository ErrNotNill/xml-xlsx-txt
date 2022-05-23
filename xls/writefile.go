package xls

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"xml-txt/xml"
)

func WriteXlsFile(){
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// Set value of a cell.
	//f.SetCellValue("Sheet1", "C5", 100)


			for _,b := range xml.Article {
				f.SetCellValue("Sheet1","B",b)

			}



	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}