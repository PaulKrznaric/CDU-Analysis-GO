package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

//read an excel file
func readFile() {
	//open the file
	file, err := excelize.OpenFile("/Users/paulkrznaric/Documents/Work/CDU/CDU Details April 2022.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	//close the file
	defer file.Close()

}
