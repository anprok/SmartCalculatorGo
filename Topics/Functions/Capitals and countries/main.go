package main

import "fmt"

// Update the parameter list of the countryCapital function below:
func countryCapital(capital string, country string) string {
	return capital + " is the capital of " + country
}

// DO NOT modify or delete the contents within the main function!
func main() {
	var capital, country string
	fmt.Scanln(&capital, &country)

	fmt.Println(countryCapital(capital, country))
}
