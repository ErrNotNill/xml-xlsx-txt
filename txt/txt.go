package txt

import (
	"fmt"
	"os"
	"strings"
	"xml-txt/pkg"
)

/*4;2 = article
4 = brand
5 = stock
6 = cost
....*/

//1-row (choose excel row) 2-column (choose excel column) //row/column
//3-column (new parse xml argument) row = previous, 4-name column = (choose xml parse argument)

var Args []string

func GetArgsFromTxt(){
	read,_ :=os.ReadFile(pkg.TxtFile)
	//fmt.Println(string(read))

	split := strings.Split(string(read),";")
	Args=append(split)
	/*for _,v := range Args{
		fmt.Println(v)
	}*/
	fmt.Println(Args[0])
	//fmt.Println(Args)
	/*split[0] = ExcelRow()
	split[1] = ExcelColumn()
	split[2] = NextColumn()*/

fmt.Println(split[2])
fmt.Println(split[6])
}

func ExcelRow() string{
	return ""
}
func ExcelColumn() string{
	return ""
}
func NextColumn() string{
	return ""
}