package main

import (
	"fmt"
	"os"
)

func main(){

	file,_ := os.ReadFile("rules.txt") //todo filename correct TxtFile
	fmt.Println(string(file))
}

