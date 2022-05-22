package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)


func main() {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var sksk = make(map[string]int)

	for _, file := range files {
sksk[file.Name()]++
	}

	var count int

	for k := range sksk{

		switch strings.Contains(k,".txt") {
		case true:
			fmt.Println(k)
			count+=1
		case false:
			break
		}
		switch strings.Contains(k,".xml") {
		case true:
			fmt.Println(k)
			count+=1
		case false:
			break
		}
		switch strings.Contains(k,".xlsx") {
		case true:
			fmt.Println(k)
			count+=1
		case false:
			break
		}
		}

		if count > 3 || count < 3 {
			fmt.Printf("Error. I found %v file(s), but i need 3",count)
			fmt.Println()
			os.Exit(1)
		} else {
			fmt.Println("Files found, you can work. Status : OK")
		}
	}
