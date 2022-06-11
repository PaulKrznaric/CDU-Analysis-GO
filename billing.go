package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Billing struct {
	count int
}

func (b *Billing) GetBilling() string {
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

type MinorBillingLine struct {
	billingNumber int
	doctorName    string
	date          time.Time
	billingValues [7]Billing
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
	values := [7]Billing{NewBilling(0), NewBilling(0), NewBilling(overflow), NewBilling(0), NewBilling(0), NewBilling(0), NewBilling(0)}
	return values
}

type BillingLine struct {
	billingNumber     int
	doctorName        string
	date              time.Time
	billingValues     [7]Billing
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
