package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

	func FindAllTypesInDir() error{
		files, err := ioutil.ReadDir(".")
		if err != nil {
			fmt.Println("cant read current directory")
			log.Fatal(err)
		}

		var sksk = make(map[string]int)

		var count int

		for _, file := range files {
			sksk[file.Name()]++

			xls,_ := regexp.MatchString(".xls",file.Name())
			if xls == true{
				count+=1
				fmt.Println("I found excel")
			}

			txt,_ := regexp.MatchString(".txt",file.Name())
			if txt == true{
				count+=1
				fmt.Println("i found txt")
			}

			xml,_ := regexp.MatchString(".xml",file.Name())
			if xml == true{
				count+=1
				fmt.Println("i found xml")
			}
		}

		if count > 3 || count < 3 {
			fmt.Printf("Error. I found %v file(s), but i need 3",count)
			fmt.Println()
			os.Exit(1)
		} else {
			fmt.Println("Files found, you can work. Status : OK")
		}
		return err
	}
