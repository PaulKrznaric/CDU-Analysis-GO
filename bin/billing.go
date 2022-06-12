package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Billing struct {
	count int
}

func (b *Billing) GetBilling() string {
	if b.count == 0 {
		return ""
	}
	return strconv.Itoa(b.count)
}

func (b *Billing) SetBilling(value int) {
	b.count = value
}

func NewBilling(count int) Billing {
	billing := Billing{}
	billing.count = count
	return billing
}

//implement an abstract interface
type IPrintBilling interface {
	PrintBilling() [10]string
}

type IBillingLine struct {
	IPrintBilling
	billingNumber int
	doctorName    string
	date          time.Time
	billingValues [7]Billing
}

func (b *IBillingLine) PrintBilling() [10]string {
	var output [10]string
	output[0] = b.doctorName
	output[1] = strconv.Itoa(b.billingNumber)
	output[2] = b.date.Format("2006-01-02")
	for i := 0; i < 7; i++ {
		output[i+3] = b.billingValues[i].GetBilling()
	}
	return output
}

type MinorBillingLine struct {
	IBillingLine
}

func NewMinorBilling(parentBilling BillingLine, doctorName string, overflow int) MinorBillingLine {
	line := MinorBillingLine{}
	line.billingNumber = parentBilling.billingNumber
	line.doctorName = doctorName
	line.date = parentBilling.date
	line.billingValues = CreateMinorBillingValue(overflow)
	return line
}

func CreateMinorBillingValue(overflow int) [7]Billing {
	values := [7]Billing{NewBilling(0), NewBilling(overflow), NewBilling(0), NewBilling(0), NewBilling(0), NewBilling(0), NewBilling(0)}
	return values
}

type BillingLine struct {
	IBillingLine
	associatedBilling MinorBillingLine
}

func NewBillingLine(billingNumber string, doctorName string, secondDoctor string, timeIn time.Time, timeOut time.Time, timeAdmitted *time.Time) BillingLine {
	line := BillingLine{}
	line.billingNumber = FormatBillingNumber(billingNumber)
	line.doctorName = doctorName
	line.date = timeIn
	line.billingValues, line.associatedBilling = CalculateBillingValue(line, secondDoctor, timeIn, timeOut, timeAdmitted)
	return line
}

func FormatBillingNumber(billingNumber string) int {
	billingNumber = strings.TrimPrefix(strings.ToUpper(billingNumber), "J")
	num, err := strconv.Atoi(billingNumber)
	if err != nil {
		fmt.Println((err))
		fmt.Println("There's been an issue with number {0}", billingNumber)
		os.Exit(2)
	}
	return num
}

func CalculateBillingValue(line BillingLine, secondDoctor string, timeIn time.Time, timeOut time.Time, timeAdmitted *time.Time) ([7]Billing, MinorBillingLine) {
	var billings [7]Billing
	var minor MinorBillingLine
	var timeOutER time.Time
	hasSecondDoctor := secondDoctor != ""

	billings[0] = NewBilling(1)

	if timeAdmitted == nil {
		billings[5] = NewBilling(1)
		billings[6] = NewBilling(0)
		timeOutER = timeOut
	} else {
		billings[5] = NewBilling(0)
		billings[6] = NewBilling(1)
		timeOutER = *timeAdmitted
	}

	for i := 1; i < 5; i++ {
		billings[i] = NewBilling(0)
	}

	totalTime := (timeOutER.Unix() - timeIn.Unix()) / 3600
	billingCount := int(totalTime / 2)

	//anything more than one billing is automatically turned into a CD0R
	if billingCount >= 2 {
		if hasSecondDoctor && billingCount > 3 {
			minor = NewMinorBilling(line, secondDoctor, 3)
			if billingCount >= 6 {
				billingCount = 3
			} else {
				billingCount = billingCount - 3
			}
			billings[1].SetBilling(billingCount)
		} else {
			if billingCount > 3 {
				billingCount = 3
			}
			billings[1].SetBilling(int(billingCount))
		}
	} else if billingCount == 1 {
		if hasSecondDoctor {
			minor = NewMinorBilling(line, secondDoctor, 0)
			minor.billingValues = CalculateIndividualBillingValue(minor.billingValues, timeIn)
		} else {
			billings = CalculateIndividualBillingValue(billings, timeIn)
		}
	}
	return billings, minor
}

func CalculateIndividualBillingValue(billings [7]Billing, timeIn time.Time) [7]Billing {
	if timeIn.Hour() < 7 {
		billings[4].SetBilling(1)
	} else if int(timeIn.Weekday()) > 5 {
		billings[3].SetBilling(1)
	} else if timeIn.Hour() > 17 {
		billings[2].SetBilling(1)
	} else {
		billings[1].SetBilling(1)
	}
	return billings
}

type BillingGroup struct {
	billingLines []IBillingLine
}

//create a BillingLine and append it to the billingLines array
func (b *BillingGroup) AddBillingLine(billingNumber string, doctorName string, secondDoctor string, timeIn time.Time, timeOut time.Time, timeAdmitted *time.Time) {
	line := NewBillingLine(billingNumber, doctorName, secondDoctor, timeIn, timeOut, timeAdmitted)
	b.billingLines = append(b.billingLines, line.IBillingLine)
	if line.associatedBilling.doctorName != "" {
		b.billingLines = append(b.billingLines, line.associatedBilling.IBillingLine)
	}
}

//sort BillingGroup Billinglines by doctor then by date
func (b *BillingGroup) SortBillingLines() {
	sort.Slice(b.billingLines, func(i, j int) bool {
		if b.billingLines[i].doctorName == b.billingLines[j].doctorName {
			return b.billingLines[i].date.Before(b.billingLines[j].date)
		}
		return b.billingLines[i].doctorName < b.billingLines[j].doctorName
	})
}
