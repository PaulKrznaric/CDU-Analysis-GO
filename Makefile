build:
	go build bin/*.go"

run:
	go run bin/billing.go bin/read.go bin/write.go bin/go-cdu.go 