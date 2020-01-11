package util

import "math"
import "strconv"
import "fmt"

// IsLeapYear determines if a given year is a leap year.
func IsLeapYear(inputYear int) bool {
	year := float64(inputYear)

	if math.Mod(year, 4) == 0 {
		if math.Mod(year, 100) == 0 {
			if math.Mod(year, 400) == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

// RoundFloat64 rounds a float to a specified number of decimal places.
func RoundFloat64(inputValue float64, places int) float64 {
	returnValue, _ := strconv.ParseFloat(fmt.Sprintf("%.*f", places, inputValue), 64)

	return returnValue
}

// DegreesToRadians converts units from degrees to radians
func DegreesToRadians(inputDegrees float64) float64 {
	return inputDegrees * (math.Pi / 180)
}
