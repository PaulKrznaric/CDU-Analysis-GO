package main

func main() {
	//get data
	billingGroup := readFile()

	billingGroup.SortBillingLines()
	//output data
	export(billingGroup)
}
