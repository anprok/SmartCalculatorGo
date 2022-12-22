package main

import (
	"fmt"
	"log"
	"time"
)

const (
	HoursPerDay = 24
	DateLayout  = "2006-01-02"
)

func main() {
	var startDateString, endDateString string
	var salaryPerHour, workingTime uint64

	fmt.Scan(&startDateString, &endDateString) // Read the start date and the end date
	fmt.Scan(&salaryPerHour, &workingTime)     // Read the salary per hour and the working time

	startDate, err := time.Parse(DateLayout, startDateString)
	if err != nil {
		log.Fatal(err)
	}

	endDate, err := time.Parse(DateLayout, endDateString) // Parse the end date of the period
	if err != nil {
		log.Fatal(err)
	}

	// The calculateSalary() function outputs the total salary for the given period:
	calculateSalary(startDate, endDate, salaryPerHour, workingTime)
}

// Write the required code to finish the implementation of the calculateSalary() function:
func calculateSalary(startDate time.Time, endDate time.Time, salaryPerHour uint64, workingTime uint64) {
	var totalSalary uint64

	for !startDate.Equal(endDate) {
		weekday := startDate.Weekday() // What function returns the day of the week?
		startDate = startDate.Add(time.Hour * HoursPerDay)

		salaryPerDay := salaryPerHour * workingTime // Calculate the salary per day

		if weekday == time.Wednesday || weekday == time.Saturday || weekday == time.Sunday {
			continue
		}

		if weekday == time.Friday {
			salaryPerDay -= salaryPerHour
		}

		totalSalary += salaryPerDay // Add the daily salary to the total salary
	}
	fmt.Println("Total:", totalSalary)
}
