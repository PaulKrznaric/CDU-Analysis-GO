package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

//read an excel file
func readFile() BillingGroup {
	//open the file
	file, err := excelize.OpenFile("/Users/paulkrznaric/Documents/Work/CDU/CDU Details May 2022.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Opened sheet")
	//create a billinggroup
	var billingGroup BillingGroup
	//read the file line by line
	sheet := file.GetSheetName(0)
	rows, err := file.GetRows(sheet, excelize.Options{RawCellValue: true})
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range rows {
		if row[0] == "MRN" || strings.Contains(row[0], "***") {
			continue
		}
		id := row[0]
		timeIn := createTime(row[4])
		primaryDoctor := row[3]
		secondaryDoctor := row[7]
		if secondaryDoctor == "." || secondaryDoctor == primaryDoctor {
			secondaryDoctor = ""
		}
		primaryDoctor = formatDoctor(primaryDoctor)
		secondaryDoctor = formatDoctor(secondaryDoctor)

		timeOut := createTime(row[8])
		timeAdmitted := createTime(row[10])
		if len(row) != 13 {
			fmt.Println(len(row))
			fmt.Println(id, " ", timeIn, " ", primaryDoctor, " ", secondaryDoctor)
		}
		billingGroup.AddBillingLine(id, primaryDoctor, secondaryDoctor, *timeIn, *timeOut, timeAdmitted)
		log.Println(billingGroup.billingLines[len(billingGroup.billingLines)-1].PrintBilling())
	}
	//close the file
	defer file.Close()

	return billingGroup

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

func formatDoctor(original string) string {
	hasPrefix := strings.HasPrefix(original, "N.")
	if hasPrefix && original != "" {
		return "DR." + strings.TrimPrefix(original, "N.")
	}
	return original
}
