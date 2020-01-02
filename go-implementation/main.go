package main

import "fmt"
import "./lib/datetime"

func main() {
	var month, day, year int = datetime.GetDateOfEaster(2009)

	fmt.Printf("Date of Easter for %d is %d/%d/%d\n", year, month, day, year)
}
