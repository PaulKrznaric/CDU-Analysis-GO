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

func NewBilling(count int) Billing {
	billing := Billing{}
	billing.count = count
	return billing
}

type BillingLine struct {
	billingNumber int
	doctorName    string
	date          time.Time
	billingValues [7]Billing
}

func NewBillingLine(billingNumber string, doctorName string, timeIn time.Time, timeOut time.Time, timeAdmitted time.Time) BillingLine {
	line := BillingLine{}
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
