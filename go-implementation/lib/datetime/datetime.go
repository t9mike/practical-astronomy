package datetime

import (
	"../macros"
	"../util"
	"math"
)

// GetDateOfEaster calculates the date upon which Easter falls for a given year.
func GetDateOfEaster(inputYear int) (int, int, int) {
	year := float64(inputYear)

	a := math.Mod(year, 19)
	b := math.Floor(year / 100)
	c := math.Mod(year, 100)
	d := math.Floor(b / 4)
	e := math.Mod(b, 4)
	f := math.Floor((b + 8) / 25)
	g := math.Floor((b - f + 1) / 3)
	h := math.Mod(((19 * a) + b - d - g + 15), 30)
	i := math.Floor(c / 4)
	k := math.Mod(c, 4)
	l := math.Mod((32 + 2*(e+i) - h - k), 7)
	m := math.Floor((a + (11 * h) + (22 * l)) / 451)
	n := math.Floor((h + l - (7 * m) + 114) / 31)
	p := math.Mod((h + l - (7 * m) + 114), 31)

	day := p + 1
	month := n

	return int(month), int(day), int(year)
}

// CivilDateToDayNumber determines the day number for a given civil date.
func CivilDateToDayNumber(month int, day int, year int) int {
	if month <= 2 {
		month = month - 1

		if util.IsLeapYear(year) {
			month = month * 62
		} else {
			month = month * 63
		}

		month = int(math.Floor(float64(month) / 2))
	} else {
		month = int(math.Floor((float64(month) + 1) * 30.6))

		if util.IsLeapYear(year) {
			month = month - 62
		} else {
			month = month - 63
		}
	}

	return month + day
}

// CivilTimeToDecimalHours converts a Civil Time (hours,minutes,seconds) to Decimal Hours
func CivilTimeToDecimalHours(hours float64, minutes float64, seconds float64) float64 {
	return macros.HmsDh(hours, minutes, seconds)
}

// DecimalHoursToCivilTime converts decimal hours to civil time.
//
// Returns hours, minutes, and seconds.
func DecimalHoursToCivilTime(decimalHours float64) (int, int, int) {
	hours := macros.DhHour(decimalHours)
	minutes := macros.DhMin(decimalHours)
	seconds := macros.DhSec(decimalHours)

	return hours, minutes, int(seconds)
}
