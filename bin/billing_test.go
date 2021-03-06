package main

import (
	"testing"
	"time"
)

func TestFormatBillingNumber(t *testing.T) {
	got := FormatBillingNumber("J00123")
	want := 123
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func RunTestCalculateBilling(t *testing.T, line BillingLine, secondDoctorName string, timeIn time.Time, timeOut time.Time, timeAdmitted *time.Time, expectedBilling [7]Billing, expectedMinor MinorBillingLine) {
	gotBilling, gotMinor := CalculateBillingValue(line, secondDoctorName, line.date, timeOut, timeAdmitted)
	if gotBilling != expectedBilling {
		t.Errorf("Got Billing is not the same as expected")
		for i := 0; i < 7; i++ {
			t.Errorf("Line: %d", i)
			t.Errorf(gotBilling[i].GetBilling())
			t.Errorf(expectedBilling[i].GetBilling())
		}
	}
	if expectedMinor != gotMinor {
		t.Errorf("Got minor is not the same as expected")
		for i := 0; i < 7; i++ {
			t.Errorf("Line: %d", i)
			t.Errorf(gotMinor.billingValues[i].GetBilling())
			t.Errorf(expectedMinor.billingValues[i].GetBilling())
		}
	}
}

func TestCalculateBillingNotAdmitted(t *testing.T) {
	var line BillingLine
	secondDoctorName := "Dr. Gupta"
	line.doctorName = "Dr. Joe"
	expectedBilling := [7]Billing{NewBilling(1), NewBilling(3), NewBilling(0), NewBilling(0), NewBilling(0), NewBilling(1), NewBilling(0)}
	timeIn := time.Date(2022, time.June, 4, 20, 10, 10, 10, time.Local)
	timeOut := time.Date(2022, time.June, 5, 10, 10, 10, 10, time.Local)
	line.date = timeIn
	expectedMinor := NewMinorBilling(line, secondDoctorName, 3)
	RunTestCalculateBilling(t, line, secondDoctorName, timeIn, timeOut, nil, expectedBilling, expectedMinor)
}

func TestCalculatedBillingAdmitted(t *testing.T) {
	var line BillingLine
	secondDoctorName := "Dr. Gupta"
	line.doctorName = "Dr. Joe"
	expectedBilling := [7]Billing{NewBilling(1), NewBilling(3), NewBilling(0), NewBilling(0), NewBilling(0), NewBilling(0), NewBilling(1)}
	timeIn := time.Date(2022, time.June, 4, 20, 10, 10, 10, time.Local)
	timeOut := time.Date(2022, time.June, 5, 10, 10, 10, 10, time.Local)
	timeAdmitted := time.Date(2022, time.June, 5, 10, 10, 10, 10, time.Local)
	line.date = timeIn
	expectedMinor := NewMinorBilling(line, secondDoctorName, 3)
	RunTestCalculateBilling(t, line, secondDoctorName, timeIn, timeOut, &timeAdmitted, expectedBilling, expectedMinor)
}

func TestCalculatedBillingNoSecondDoctor(t *testing.T) {
	var line BillingLine
	line.doctorName = "Dr. Joe"
	expectedBilling := [7]Billing{NewBilling(1), NewBilling(3), NewBilling(0), NewBilling(0), NewBilling(0), NewBilling(1), NewBilling(0)}
	timeIn := time.Date(2022, time.June, 4, 20, 10, 10, 10, time.Local)
	timeOut := time.Date(2022, time.June, 5, 10, 10, 10, 10, time.Local)
	line.date = timeIn
	var expectedMinorBilling MinorBillingLine
	RunTestCalculateBilling(t, line, "", timeIn, timeOut, nil, expectedBilling, expectedMinorBilling)
}

func TestCalculatedBillingOneBillingSecondDoctor(t *testing.T) {
	var line BillingLine
	line.doctorName = "Dr. Joe"
	secondDoctorName := "Dr. Gupta"
	expectedBilling := [7]Billing{NewBilling(1), NewBilling(0), NewBilling(0), NewBilling(0), NewBilling(0), NewBilling(1), NewBilling(0)}
	timeIn := time.Date(2022, time.June, 4, 20, 10, 10, 10, time.Local)
	timeOut := time.Date(2022, time.June, 4, 22, 10, 10, 10, time.Local)
	line.date = timeIn
	expectedMinor := NewMinorBilling(line, secondDoctorName, 0)
	expectedMinor.billingValues[3].SetBilling(1)
	RunTestCalculateBilling(t, line, secondDoctorName, timeIn, timeOut, nil, expectedBilling, expectedMinor)
}

func TestCalculatedBillingOneBillingEarlyMorning(t *testing.T) {
	var line BillingLine
	line.doctorName = "Dr. Joe"
	expectedBilling := [7]Billing{NewBilling(1), NewBilling(0), NewBilling(0), NewBilling(0), NewBilling(1), NewBilling(1), NewBilling(0)}
	timeIn := time.Date(2022, time.June, 4, 6, 10, 10, 10, time.Local)
	timeOut := time.Date(2022, time.June, 4, 8, 10, 10, 10, time.Local)
	line.date = timeIn
	var expectedMinorBilling MinorBillingLine
	RunTestCalculateBilling(t, line, "", timeIn, timeOut, nil, expectedBilling, expectedMinorBilling)
}

func TestCalculatedBillingOneBillingNight(t *testing.T) {
	var line BillingLine
	line.doctorName = "Dr. Joe"
	expectedBilling := [7]Billing{NewBilling(1), NewBilling(0), NewBilling(0), NewBilling(1), NewBilling(0), NewBilling(1), NewBilling(0)}
	timeIn := time.Date(2022, time.June, 3, 20, 10, 10, 10, time.Local)
	timeOut := time.Date(2022, time.June, 3, 23, 10, 10, 10, time.Local)
	line.date = timeIn
	var expectedMinorBilling MinorBillingLine
	RunTestCalculateBilling(t, line, "", timeIn, timeOut, nil, expectedBilling, expectedMinorBilling)
}

func TestCalculatedBillingOneBilling(t *testing.T) {
	var line BillingLine
	line.doctorName = "Dr. Joe"
	expectedBilling := [7]Billing{NewBilling(1), NewBilling(1), NewBilling(0), NewBilling(0), NewBilling(0), NewBilling(1), NewBilling(0)}
	timeIn := time.Date(2022, time.June, 3, 12, 10, 10, 10, time.Local)
	timeOut := time.Date(2022, time.June, 3, 14, 10, 10, 10, time.Local)
	line.date = timeIn
	var expectedMinorBilling MinorBillingLine
	RunTestCalculateBilling(t, line, "", timeIn, timeOut, nil, expectedBilling, expectedMinorBilling)
}

func TestPrintMinorBilling(t *testing.T) {
	var line BillingLine
	line.billingNumber = 123
	line.doctorName = "Dr. Joe"
	line.date = time.Date(2022, time.June, 3, 12, 10, 10, 10, time.Local)
	minorBilling := NewMinorBilling(line, "Dr. Joe", 1)
	expected := [10]string{"Dr. Joe", "123", line.date.Format("2006-01-02"), "", "1", "", "", "", "", ""}
	got := minorBilling.PrintBilling()
	if expected != got {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestPrintBilling(t *testing.T) {
	var line BillingLine
	line.billingNumber = 123
	line.doctorName = "Dr. Joe"
	line.date = time.Date(2022, time.June, 3, 12, 10, 10, 10, time.Local)
	line.billingValues = [7]Billing{NewBilling(1), NewBilling(1), NewBilling(0), NewBilling(0), NewBilling(0), NewBilling(1), NewBilling(0)}
	expected := [10]string{"Dr. Joe", "123", line.date.Format("2006-01-02"), "1", "1", "", "", "", "1", ""}
	got := line.PrintBilling()
	if expected != got {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
