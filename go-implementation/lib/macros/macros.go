package macros

import "math"
import "../util"

// HMSDH converts a Civil Time (hours,minutes,seconds) to Decimal Hours
//
// Original macro name: HMSDH
func HMSDH(hours float64, minutes float64, seconds float64) float64 {
	fHours := hours
	fMinutes := minutes
	fSeconds := seconds

	a := math.Abs(fSeconds) / 60
	b := (math.Abs(fMinutes) + a) / 60
	c := math.Abs(fHours) + b

	if fHours < 0 || fMinutes < 0 || fSeconds < 0 {
		return -c
	}

	return c
}

// DHHour returns the hour part of a Decimal Hours.
//
// Original macro name: DHHour
func DHHour(decimalHours float64) int {
	a := math.Abs(decimalHours)
	b := a * 3600
	c := util.RoundFloat64(b-60*math.Floor(b/60), 2)
	e := b
	if c == 60 {
		e = b + 60.0
	}

	if decimalHours < 0.0 {
		return int(-math.Floor(e / 3600))
	}

	return int(math.Floor(e / 3600))
}

// DHMin returns the minutes part of a Decimal Hours.
//
// Original macro name: DHMin
func DHMin(decimalHours float64) int {
	a := math.Abs(decimalHours)
	b := a * 3600
	c := util.RoundFloat64(b-60*math.Floor(b/60), 2)
	e := b
	if c == 60 {
		e = b + 60.0
	}

	return int(math.Mod(math.Floor(e/60), 60))
}

// DHSec returns the seconds part of a Decimal Hours.
//
// Original macro name: DHSec
func DHSec(decimalHours float64) float64 {
	a := math.Abs(decimalHours)
	b := a * 3600
	c := util.RoundFloat64(b-60*math.Floor(b/60), 2)
	d := c
	if c == 60.0 {
		d = 0
	}

	return d
}

// CDJD converts a Greenwich Date/Civil Date (day,month,year) to Julian Date
//
// Original macro name: CDJD
func CDJD(day float64, month int, year int) float64 {
	fDay := float64(day)
	fMonth := float64(month)
	fYear := float64(year)

	y := fYear
	if fMonth < 3 {
		y = fYear - 1
	}
	m := fMonth
	if fMonth < 3 {
		m = fMonth + 12
	}

	var b float64
	if fYear > 1582 {
		a := math.Floor(y / 100)
		b = 2 - a + math.Floor(a/4)
	} else {
		if fYear == 1582 && fMonth > 10 {
			a := math.Floor(y / 100)
			b = 2 - a + math.Floor(a/4)
		} else {
			if fYear == 1582 && fMonth == 10 && fDay >= 15 {
				a := math.Floor(y / 100)
				b = 2 - a + math.Floor(a/4)
			} else {
				b = 0
			}
		}
	}

	var c float64
	if y < 0 {
		c = math.Floor((365.25 * y) - 0.75)
	} else {
		c = math.Floor(365.25 * y)
	}

	d := math.Floor(30.6001 * (m + 1))

	return b + c + d + fDay + 1720994.5
}

// JDCDay returns the day part of a Julian Date
//
// Original macro name: JDCDay
func JDCDay(julianDate float64) float64 {
	i := math.Floor(julianDate + 0.5)
	f := julianDate + 0.5 - i
	a := math.Floor((i - 1867216.25) / 36524.25)
	var b float64
	if i > 2299160 {
		b = i + 1 + a - math.Floor(a/4)
	} else {
		b = i
	}
	c := b + 1524
	d := math.Floor((c - 122.1) / 365.25)
	e := math.Floor(365.25 * d)
	g := math.Floor((c - e) / 30.6001)

	return c - e + f - math.Floor(30.6001*g)
}

// JDCMonth returns the month part of a Julian Date
//
// Original macro name: JDCMonth
func JDCMonth(julianDate float64) int {
	i := math.Floor(julianDate + 0.5)
	a := math.Floor((i - 1867216.25) / 36524.25)

	var b float64
	if i > 2299160 {
		b = i + 1.0 + a - math.Floor(a/4)
	} else {
		b = i
	}

	c := b + 1524
	d := math.Floor((c - 122.1) / 365.25)
	e := math.Floor(365.25 * d)
	g := math.Floor((c - e) / 30.6001)

	var returnValue float64
	if g < 13.5 {
		returnValue = g - 1
	} else {
		returnValue = g - 13
	}

	return int(returnValue)
}

// JDCYear returns the year part of a Julian Date
//
// Original macro name: JDCYear
func JDCYear(julianDate float64) int {
	i := math.Floor(julianDate + 0.5)
	a := math.Floor((i - 1867216.25) / 36524.25)

	var b float64
	if i > 2299160 {
		b = i + 1 + a - math.Floor(a/4)
	} else {
		b = i
	}
	c := b + 1524
	d := math.Floor((c - 122.1) / 365.25)
	e := math.Floor(365.25 * d)
	g := math.Floor((c - e) / 30.6001)

	var h float64
	if g < 13.5 {
		h = g - 1
	} else {
		h = g - 13
	}

	var returnValue float64
	if h > 2.5 {
		returnValue = d - 4716
	} else {
		returnValue = d - 4715
	}

	return int(returnValue)
}

// FDOW converts a Julian Date to Day-of-Week (e.g., Sunday)
//
// Original macro name: FDOW
func FDOW(julianDate float64) string {
	j := math.Floor(julianDate-0.5) + 0.5
	n := math.Mod(j+1.5, 7)

	var returnValue string
	switch n {
	case 0:
		returnValue = "Sunday"
	case 1:
		returnValue = "Monday"
	case 2:
		returnValue = "Tuesday"
	case 3:
		returnValue = "Wednesday"
	case 4:
		returnValue = "Thursday"
	case 5:
		returnValue = "Friday"
	case 6:
		returnValue = "Saturday"
	default:
		returnValue = "Unknown"
	}

	return returnValue
}
