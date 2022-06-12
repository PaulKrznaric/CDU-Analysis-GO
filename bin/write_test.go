package main

import (
	"log"
	"testing"
	"time"
)

func TestWriteLine(t *testing.T) {

	billingNumber := "J00123"
	doctorName := "Dr. John Doe"
	timeIn := time.Date(2022, time.June, 4, 20, 10, 10, 10, time.Local)
	timeOut := time.Date(2022, time.June, 4, 22, 10, 10, 10, time.Local)
	secondDoctor := "Dr. Jane Doe"
	line := NewBillingLine(billingNumber, doctorName, secondDoctor, timeIn, timeOut, nil)
	writeLine(line.IBillingLine)
	if err := file.SaveAs("../test.xlsx"); err != nil {
		log.Fatal(err)
	}
}
