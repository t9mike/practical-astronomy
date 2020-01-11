package macros

import "math"
import "../util"

// HMSToDH converts a Civil Time (hours,minutes,seconds) to Decimal Hours
//
// Original macro name: HMSDH
func HMSToDH(hours float64, minutes float64, seconds float64) float64 {
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

// CDToJD converts a Greenwich Date/Civil Date (day,month,year) to Julian Date
//
// Original macro name: CDJD
func CDToJD(day float64, month int, year int) float64 {
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

// RAToHA converts Right Ascension to Hour Angle
//
// Original macro name: RAHA
func RAToHA(raHours float64, raMinutes float64, raSeconds float64, lctHours float64, lctMinutes float64, lctSeconds float64, daylightSaving int, zoneCorrection int, localDay float64, localMonth int, localYear int, geographicalLongitude float64) float64 {
	a := LCTToUT(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	b := LCTGreenwichDay(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	c := LCTGreenwichMonth(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	d := LCTGreenwichYear(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	e := UTToGST(a, 0.0, 0.0, b, c, d)
	f := GSTToLST(e, 0.0, 0.0, geographicalLongitude)
	g := HMSToDH(raHours, raMinutes, raSeconds)
	h := f - g

	if h < 0.0 {
		return 24.0 + h
	}
	return h
}

// HAToRA converts Hour Angle to Right Ascension
//
// Original macro name: HARA
func HAToRA(hourAngleHours float64, hourAngleMinutes float64, hourAngleSeconds float64, lctHours float64, lctMinutes float64, lctSeconds float64, daylightSaving int, zoneCorrection int, localDay float64, localMonth int, localYear int, geographicalLongitude float64) float64 {
	a := LCTToUT(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	b := LCTGreenwichDay(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	c := LCTGreenwichMonth(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	d := LCTGreenwichYear(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	e := UTToGST(a, 0.0, 0.0, b, c, d)
	f := GSTToLST(e, 0.0, 0.0, geographicalLongitude)
	g := HMSToDH(hourAngleHours, hourAngleMinutes, hourAngleSeconds)
	h := f - g

	if h < 0.0 {
		return 24.0 + h
	}
	return h
}

// LCTToUT converts Local Civil Time to Universal Time
//
// Original macro name: LctUT
func LCTToUT(lctHours float64, lctMinutes float64, lctSeconds float64, daylightSaving int, zoneCorrection int, localDay float64, localMonth int, localYear int) float64 {
	a := HMSToDH(lctHours, lctMinutes, lctSeconds)
	b := a - float64(daylightSaving) - float64(zoneCorrection)
	c := localDay + (b / 24.0)
	d := CDToJD(c, localMonth, localYear)
	e := JDCDay(d)
	e1 := math.Floor(e)

	return 24.0 * (e - e1)
}

// UTToLCT converts Universal Time to Local Civil Time
//
// Original macro name: UTLct
func UTToLCT(uHours float64, uMinutes float64, uSeconds float64, daylightSaving int, zoneCorrection int, greenwichDay float64, greenwichMonth int, greenwichYear int) float64 {
	a := HMSToDH(uHours, uMinutes, uSeconds)
	b := a + float64(zoneCorrection)
	c := b + float64(daylightSaving)
	d := CDToJD(greenwichDay, greenwichMonth, greenwichYear) + (c / 24.0)
	e := JDCDay(d)
	e1 := math.Floor(e)

	return 24.0 * (e - e1)
}

// LCTGreenwichDay determines Greenwich Day for Local Time
//
// Original macro name: LctGDay
func LCTGreenwichDay(lctHours float64, lctMinutes float64, lctSeconds float64, daylightSaving int, zoneCorrection int, localDay float64, localMonth int, localYear int) float64 {
	a := HMSToDH(lctHours, lctMinutes, lctSeconds)
	b := a - float64(daylightSaving) - float64(zoneCorrection)
	c := localDay + (b / 24.0)
	d := CDToJD(c, localMonth, localYear)
	e := JDCDay(d)

	return math.Floor(e)
}

// LCTGreenwichMonth determines Greenwich Month for Local Time
//
// Original macro name: LctGMonth
func LCTGreenwichMonth(lctHours float64, lctMinutes float64, lctSeconds float64, daylightSaving int, zoneCorrection int, localDay float64, localMonth int, localYear int) int {
	a := HMSToDH(lctHours, lctMinutes, lctSeconds)
	b := a - float64(daylightSaving) - float64(zoneCorrection)
	c := localDay + (b / 24.0)
	d := CDToJD(c, localMonth, localYear)

	return JDCMonth(d)
}

// LCTGreenwichYear determines Greenwich Year for Local Time
//
// Original macro name: LctGYear
func LCTGreenwichYear(lctHours float64, lctMinutes float64, lctSeconds float64, daylightSaving int, zoneCorrection int, localDay float64, localMonth int, localYear int) int {
	a := HMSToDH(lctHours, lctMinutes, lctSeconds)
	b := a - float64(daylightSaving) - float64(zoneCorrection)
	c := localDay + (b / 24.0)
	d := CDToJD(c, localMonth, localYear)

	return JDCYear(d)
}

// UTToGST converts Universal Time to Greenwich Sidereal Time
//
// Original macro name: UTGST
func UTToGST(uHours float64, uMinutes float64, uSeconds float64, greenwichDay float64, greenwichMonth int, greenwichYear int) float64 {
	a := CDToJD(greenwichDay, greenwichMonth, greenwichYear)
	b := a - 2451545.0
	c := b / 36525.0
	d := 6.697374558 + (2400.051336 * c) + (0.000025862 * c * c)
	e := d - (24.0 * math.Floor(d/24.0))
	f := HMSToDH(uHours, uMinutes, uSeconds)
	g := f * 1.002737909
	h := e + g

	return h - (24.0 * math.Floor(h/24.0))
}

// GSTToLST converts Greenwich Sidereal Time to Local Sidereal Time
//
// Original macro name: GSTLST
func GSTToLST(
	greenwichHours float64,
	greenwichMinutes float64,
	greenwichSeconds float64,
	geographicalLongitude float64,
) float64 {
	a := HMSToDH(greenwichHours, greenwichMinutes, greenwichSeconds)
	b := geographicalLongitude / 15.0
	c := a + b

	return c - (24.0 * math.Floor(c/24.0))
}