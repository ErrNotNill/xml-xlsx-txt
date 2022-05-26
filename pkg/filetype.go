package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var ExcelFile string
var XmlFile string
var TxtFile string

//todo fix a rules with search (example .xls) in file name, not in type

func CheckExcel(filename string) bool {
	exceltypes := []string{".xls", ".xlsx", ".xlam", ".xlsm", ".xltm", ".xltx", ".xlt", ".xlr"}
	for _, v := range exceltypes {
		if strings.HasSuffix(filename, v) {
			return true
		}
	}
	return false
}

func FindAllTypesInDir() error {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("cant read current directory")
		log.Fatal(err)
	}

	var sksk = make(map[string]int)

	var count int

	for _, file := range files {
		sksk[file.Name()]++

		xls, _ := regexp.MatchString(".xl", file.Name())
		if xls == true {
			if CheckExcel(file.Name()) == true {
				count += 1
				fmt.Println("I found excel...OK")
				ExcelFile = file.Name()
			}
		}

		txt, _ := regexp.MatchString(".txt", file.Name())
		if txt == true {
			if strings.HasSuffix(file.Name(), ".txt") {
				count += 1
				fmt.Println("i found txt...OK")
				TxtFile = file.Name()
			}
		}

		xml, _ := regexp.MatchString(".xml", file.Name())
		if xml == true {
			if strings.HasSuffix(file.Name(), ".xml") {
				count += 1
				fmt.Println("i found xml...OK")
				XmlFile = file.Name()
			}
		}
	}

	if count > 3 || count < 3 {
		fmt.Printf("Error. I found %v file(s), but i need 3", count)
		fmt.Println()
		os.Exit(1)
	} else {
		fmt.Println("Files found, you can work. Status : OK")
	}
	return err
}
