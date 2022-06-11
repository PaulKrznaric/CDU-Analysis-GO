package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xuri/excelize/v2"
)

type ExcelTracker struct {
	row    int
	column int
}

func NewExcelTracker(row int, col int) ExcelTracker {
	excelTracker := ExcelTracker{}
	excelTracker.row = row
	excelTracker.column = col
	return excelTracker
}

func (m *ExcelTracker) getCurrentRow() int {
	return m.row
}

func (m *ExcelTracker) setRow(row int) {
	m.row = row
}

func (m *ExcelTracker) getCurrentColumn() int {
	return m.column
}

func (m *ExcelTracker) incramentCol() {
	m.column++
}

var (
	tracker ExcelTracker
	file    *excelize.File
)

func init() {
	tracker = NewExcelTracker(0, 0)
	file = excelize.NewFile()
}

func writeLine() {
	tracker.setRow(0)

	tracker.incramentCol()
}

func dontFormat() {
	fmt.Println("asdf")
	excelize.CoordinatesToCellName(tracker.getCurrentRow(), tracker.getCurrentColumn())
}

func writeOut() {
	f := excelize.NewFile()
	file.SetCellValue("Sheet1", "B2", 100)
	file.SetCellValue("Sheet1", "A1", 50)
	now := time.Now()
	file.SetCellValue("Sheet1", "A4", now.Format(time.ANSIC))

	if err := f.SaveAs("test.xlsx"); err != nil {
		log.Fatal(err)
	}
}