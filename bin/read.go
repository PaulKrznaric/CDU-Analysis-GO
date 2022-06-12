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

	//read the file line by line
	for i := 1; i <= file.SheetCount; i++ {
		sheet := file.GetSheetName(i)
		rows, err := file.GetRows(sheet)
		if err != nil {
			log.Fatal(err)
		}
		for _, row := range rows {
			for _, cell := range row {
				log.Print(cell)
			}
		}
	}

	//close the file
	defer file.Close()

}
