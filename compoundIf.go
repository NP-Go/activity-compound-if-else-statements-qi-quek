package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Input year for leap year check")
	var yearInput int
	var message string
	fmt.Scan(&yearInput)

	if yearInput%4 == 0 {
		if yearInput%100 == 0 {
			message = strconv.Itoa(yearInput) + " is a leap year"
		} else {
			message = strconv.Itoa(yearInput) + " is not a leap year"
		}
	} else {
		message = strconv.Itoa(yearInput) + " is not a leap year"
	}
	fmt.Println(message)
}
