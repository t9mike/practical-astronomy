package datetime_test

import "testing"
import "fmt"
import "../lib/util"
import "../lib/datetime"

func TestDateOfEaster(t *testing.T) {
	var month, day, year int = datetime.GetDateOfEaster(2009)

	if month == 4 && day == 12 && year == 2009 {
		fmt.Printf("Date of Easter for %d is %d/%d/%d\n", year, month, day, year)
	} else {
		t.Errorf("Expected 4/12/2009, actual is %d/%d/%d", month, day, year)
	}
}

func TestCivilDateToDayNumbers(t *testing.T) {
	CivilDateToDayNumber(t, 1, 1, 2000, 1)
	CivilDateToDayNumber(t, 3, 1, 2000, 61)
	CivilDateToDayNumber(t, 6, 1, 2003, 152)
	CivilDateToDayNumber(t, 11, 27, 2009, 331)
}

func CivilDateToDayNumber(t *testing.T, month int, day int, year int, expectedValue int) {
	dayNumber := datetime.CivilDateToDayNumber(month, day, year)

	if dayNumber == expectedValue {
		fmt.Printf("Day number: [Date] %d/%d/%d = [Day Number] %d\n", month, day, year, dayNumber)
	} else {
		t.Errorf("Expected %d, actual is %d\n", expectedValue, dayNumber)
	}
}

func TestCivilTimeToDecimalHours(t *testing.T) {
	civilHours := 18.0
	civilMinutes := 31.0
	civilSeconds := 27.0

	expectedResult := 18.52416667

	decimalHours := datetime.CivilTimeToDecimalHours(civilHours, civilMinutes, civilSeconds)
	decimalHours = util.RoundFloat64(decimalHours, 8)

	if decimalHours == expectedResult {
		fmt.Printf("Civil time to decimal hours: [Time] %d:%d:%d = [Decimal Hours] %.8f\n", int(civilHours), int(civilMinutes), int(civilSeconds), decimalHours)

	} else {
		t.Errorf("Expected %.8f, actual is %.8f\n", expectedResult, decimalHours)
	}
}

func TestDecimalHoursToCivilTime(t *testing.T) {
	decimalHours := 18.52416667

	expectedCivilHours := 18
	expectedCivilMinutes := 31
	expectedCivilSeconds := 27

	civilHours, civilMinutes, civilSeconds := datetime.DecimalHoursToCivilTime(decimalHours)

	if civilHours == expectedCivilHours && civilMinutes == expectedCivilMinutes && civilSeconds == expectedCivilSeconds {
		fmt.Printf("Decimal hours to civil time: [Decimal Hours] %.8f = [Civil Time] %d:%d:%d\n", decimalHours, civilHours, civilMinutes, civilSeconds)
	} else {
		t.Errorf("Expected %d:%d:%d, actual is %d:%d:%d\n", expectedCivilHours, expectedCivilMinutes, expectedCivilSeconds, civilHours, civilMinutes, civilSeconds)
	}
}
