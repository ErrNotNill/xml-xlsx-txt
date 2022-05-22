package main

import (
	"fmt"
	"os"
)

func main() {
	check, _ := os.ReadFile("./txt/Booking.txt")
	fmt.Println(string(check))

}
