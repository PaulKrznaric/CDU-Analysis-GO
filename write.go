package main

import (
	"log"

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
	tracker = NewExcelTracker(1, 1)
	file = excelize.NewFile()
}

func writeLine(line IBillingLine) {
	tracker.setRow(1)
	values := line.PrintBilling()
	for i := 0; i < 4; i++ {
		currentCell, err := excelize.CoordinatesToCellName(tracker.getCurrentRow(), tracker.getCurrentColumn())
		if err != nil {
			log.Fatal(err)
		}
		file.SetCellValue("Sheet1", currentCell, values[i])
		tracker.setRow(tracker.getCurrentRow() + 1)
	}
	tracker.setRow(tracker.getCurrentRow() + 1)
	for i := 4; i < 10; i++ {
		currentCell, err := excelize.CoordinatesToCellName(tracker.getCurrentRow(), tracker.getCurrentColumn())
		if err != nil {
			log.Fatal(err)
		}
		file.SetCellValue("Sheet1", currentCell, values[i])
		tracker.setRow(tracker.row + 2)
	}
	tracker.incramentCol()
}
