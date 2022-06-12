package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

//read an excel file
func readFile() {
	//open the file
	file, err := excelize.OpenFile("/Users/paulkrznaric/Documents/Work/CDU/CDU Details April 2022.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Opened sheet")

	//read the file line by line
	sheet := file.GetSheetName(0)
	rows, err := file.GetRows(sheet, excelize.Options{RawCellValue: true})
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range rows {
		if row[0] == "MRN" {
			continue
		}
		id := row[0]
		timeIn := createTime(row[2])
		primaryDoctor := row[3]
		secondaryDoctor := row[7]
		if secondaryDoctor == "." {
			secondaryDoctor = ""
		}
		if len(row) != 13 {
			fmt.Println(len(row))
			fmt.Println(id, " ", timeIn, " ", primaryDoctor, " ", secondaryDoctor)
		}
	}
	//close the file
	defer file.Close()

}

func createTime(timeString string) *time.Time {
	if timeString == "" {
		return nil
	}
	//convert the string to an int
	timeInt, err := strconv.ParseFloat(timeString, 64)
	if err != nil {
		log.Fatal(err)
	}
	excelEpoch := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
	time := excelEpoch.Add(time.Duration(timeInt * float64(24*time.Hour)))
	return &time
}
