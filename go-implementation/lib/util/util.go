package util

import "math"

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
