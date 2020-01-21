package macros

import "math"
import "strings"
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

// GSTToUT converts Greenwich Sidereal Time to Universal Time
//
// Original macro name: GSTUT
func GSTToUT(greenwichSiderealHours float64, greenwichSiderealMinutes float64, greenwichSiderealSeconds float64, greenwichDay float64, greenwichMonth int, greenwichYear int) float64 {
	a := CDToJD(greenwichDay, greenwichMonth, greenwichYear)
	b := a - 2451545.0
	c := b / 36525.0
	d := 6.697374558 + (2400.051336 * c) + (0.000025862 * c * c)
	e := d - (24.0 * math.Floor(d/24.0))
	f := HMSToDH(greenwichSiderealHours, greenwichSiderealMinutes, greenwichSiderealSeconds)
	g := f - e
	h := g - (24.0 * math.Floor(g/24.0))
	return h * 0.9972695663
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

// LSTToGST converts Local Sidereal Time to Greenwich Sidereal Time
//
// Original macro name: LSTGST
func LSTToGST(localHours float64, localMinutes float64, localSeconds float64, longitude float64) float64 {
	a := HMSToDH(localHours, localMinutes, localSeconds)
	b := longitude / 15.0
	c := a - b
	return c - (24.0 * math.Floor(c/24.0))
}

// EqToAz converts Equatorial Coordinates to Azimuth (in decimal degrees)
//
// Original macro name: EQAz
func EqToAz(hourAngleHours float64, hourAngleMinutes float64, hourAngleSeconds float64, declinationDegrees float64, declinationMinutes float64, declinationSeconds float64, geographicalLatitude float64) float64 {
	a := HMSToDH(hourAngleHours, hourAngleMinutes, hourAngleSeconds)
	b := a * 15.0
	c := util.DegreesToRadians(b)
	d := DMSToDD(declinationDegrees, declinationMinutes, declinationSeconds)
	e := util.DegreesToRadians(d)
	f := util.DegreesToRadians(geographicalLatitude)
	g := math.Sin(e)*math.Sin(f) + math.Cos(e)*math.Cos(f)*math.Cos(c)
	h := -math.Cos(e) * math.Cos(f) * math.Sin(c)
	i := math.Sin(e) - (math.Sin(f) * g)
	j := Degrees(math.Atan2(h, i))

	return j - 360.0*math.Floor(j/360.0)
}

// EqToAlt converts Equatorial Coordinates to Altitude (in decimal degrees)
//
// Original macro name: EQAlt
func EqToAlt(hourAngleHours float64, hourAngleMinutes float64, hourAngleSeconds float64, declinationDegrees float64, declinationMinutes float64, declinationSeconds float64, geographicalLatitude float64) float64 {
	a := HMSToDH(hourAngleHours, hourAngleMinutes, hourAngleSeconds)
	b := a * 15.0
	c := util.DegreesToRadians(b)
	d := DMSToDD(declinationDegrees, declinationMinutes, declinationSeconds)
	e := util.DegreesToRadians(d)
	f := util.DegreesToRadians(geographicalLatitude)
	g := math.Sin(e)*math.Sin(f) + math.Cos(e)*math.Cos(f)*math.Cos(c)

	return Degrees(math.Asin(g))
}

// DMSToDD converts Degrees Minutes Seconds to Decimal Degrees
//
// Original macro name: DMSDD
func DMSToDD(degrees float64, minutes float64, seconds float64) float64 {
	a := math.Abs(seconds) / 60.0
	b := (math.Abs(minutes) + a) / 60.0
	c := math.Abs(degrees) + b

	if degrees < 0.0 || minutes < 0.0 || seconds < 0.0 {
		return -c
	}
	return c
}

// Degrees converts W to Degrees
//
// Original macro name: Degrees
func Degrees(w float64) float64 {
	return w * 57.29577951
}

// DDDeg returns Degrees part of Decimal Degrees
//
// Original macro name: DDDeg
func DDDeg(decimalDegrees float64) float64 {
	a := math.Abs(decimalDegrees)
	b := a * 3600.0
	c := util.RoundFloat64(b-60.0*math.Floor(b/60.0), 2)

	var e float64
	if c == 60.0 {
		e = 60.0
	} else {
		e = b
	}

	if decimalDegrees < 0.0 {
		return -math.Floor(e / 3600.0)
	}
	return math.Floor(e / 3600.0)
}

// DDMin returns Minutes part of Decimal Degrees
//
// Original macro name: DDMin
func DDMin(decimalDegrees float64) float64 {
	a := math.Abs(decimalDegrees)
	b := a * 3600.0
	c := util.RoundFloat64(b-60.0*math.Floor(b/60.0), 2)

	var e float64
	if c == 60.0 {
		e = b + 60.0
	} else {
		e = b
	}

	return math.Mod(math.Floor(e/60.0), 60.0)
}

// DDSec returns Seconds part of Decimal Degrees
//
// Original macro name: DDSec
func DDSec(decimalDegrees float64) float64 {
	a := math.Abs(decimalDegrees)
	b := a * 3600.0

	c := util.RoundFloat64(b-60.0*math.Floor(b/60.0), 2)

	var d float64
	if c == 60.0 {
		d = 0.0
	} else {
		d = c
	}

	return d
}

// DDToDH converts Decimal Degrees to Degree-Hours
//
// Original macro name: DDDH
func DDToDH(decimalDegrees float64) float64 {
	return decimalDegrees / 15.0
}

// DHToDD converts Degree-Hours to Decimal Degrees
//
// Original macro name: DHDD
func DHToDD(degreeHours float64) float64 {
	return degreeHours * 15.0
}

// HorToDec converts Horizon Coordinates to Declination (in decimal degrees)
//
// Original macro name: HORDec
func HorToDec(azimuthDegrees float64, azimuthMinutes float64, azimuthSeconds float64, altitudeDegrees float64, altitudeMinutes float64, altitudeSeconds float64, geographicalLatitude float64) float64 {
	a := DMSToDD(azimuthDegrees, azimuthMinutes, azimuthSeconds)
	b := DMSToDD(altitudeDegrees, altitudeMinutes, altitudeSeconds)
	c := util.DegreesToRadians(a)
	d := util.DegreesToRadians(b)
	e := util.DegreesToRadians(geographicalLatitude)
	f := math.Sin(d)*math.Sin(e) + math.Cos(d)*math.Cos(e)*math.Cos(c)

	return Degrees(math.Asin(f))
}

// HorToHA converts Horizon Coordinates to Hour Angle (in decimal degrees)
//
// Original macro name: HORHa
func HorToHA(azimuthDegrees float64, azimuthMinutes float64, azimuthSeconds float64, altitudeDegrees float64, altitudeMinutes float64, altitudeSeconds float64, geographicalLatitude float64) float64 {
	a := DMSToDD(azimuthDegrees, azimuthMinutes, azimuthSeconds)
	b := DMSToDD(altitudeDegrees, altitudeMinutes, altitudeSeconds)
	c := util.DegreesToRadians(a)
	d := util.DegreesToRadians(b)
	e := util.DegreesToRadians(geographicalLatitude)
	f := math.Sin(d)*math.Sin(e) + math.Cos(d)*math.Cos(e)*math.Cos(c)
	g := -math.Cos(d) * math.Cos(e) * math.Sin(c)
	h := math.Sin(d) - math.Sin(e)*f
	i := DDToDH(Degrees(math.Atan2(g, h)))

	return i - 24.0*math.Floor(i/24.0)
}

// Obliq returns Obliquity of the Ecliptic for a Greenwich Date
//
// Original macro name: Obliq
func Obliq(greenwichDay float64, greenwichMonth int, greenwichYear int) float64 {
	a := CDToJD(greenwichDay, greenwichMonth, greenwichYear)
	b := a - 2415020.0
	c := (b / 36525.0) - 1.0
	d := c * (46.815 + c*(0.0006-(c*0.00181)))
	e := d / 3600.0

	return 23.43929167 - e + NutationOfObliquity(greenwichDay, greenwichMonth, greenwichYear)
}

// NutationOfObliquity returns Nutation of Obliquity
//
// Original macro name: NutatObl
func NutationOfObliquity(greenwichDay float64, greenwichMonth int, greenwichYear int) float64 {
	dj := CDToJD(greenwichDay, greenwichMonth, greenwichYear) - 2415020.0
	t := dj / 36525.0
	t2 := t * t

	a := 100.0021358 * t
	b := 360.0 * (a - math.Floor(a))

	l1 := 279.6967 + 0.000303*t2 + b
	l2 := 2.0 * util.DegreesToRadians(l1)

	a = 1336.855231 * t
	b = 360.0 * (a - math.Floor(a))

	d1 := 270.4342 - 0.001133*t2 + b
	d2 := 2.0 * util.DegreesToRadians(d1)

	a = 99.99736056 * t
	b = 360.0 * (a - math.Floor(a))

	m1 := util.DegreesToRadians(358.4758 - 0.00015*t2 + b)

	a = 1325.552359 * t
	b = 360.0 * (a - math.Floor(a))

	m2 := util.DegreesToRadians(296.1046 + 0.009192*t2 + b)

	a = 5.372616667 * t
	b = 360.0 * (a - math.Floor(a))

	n1 := util.DegreesToRadians(259.1833 + 0.002078*t2 - b)

	n2 := 2.0 * n1

	ddo := (9.21 + 0.00091*t) * math.Cos(n1)
	ddo = ddo + (0.5522-0.00029*t)*math.Cos(l2) - 0.0904*math.Cos(n2)
	ddo = ddo + 0.0884*math.Cos(d2) + 0.0216*math.Cos(l2+m1)
	ddo = ddo + 0.0183*math.Cos(d2-n1) + 0.0113*math.Cos(d2+m2)
	ddo = ddo - 0.0093*math.Cos(l2-m1) - 0.0066*math.Cos(l2-n1)

	return ddo / 3600.0
}

// SunEclipticLongitude calculates the Sun's ecliptic longitude
//
// Original macro name: SunLong
func SunEclipticLongitude(lch float64, lcm float64, lcs float64, ds int, zc int, ld float64, lm int, ly int) float64 {
	aa := LCTGreenwichDay(lch, lcm, lcs, ds, zc, ld, lm, ly)
	bb := LCTGreenwichMonth(lch, lcm, lcs, ds, zc, ld, lm, ly)
	cc := LCTGreenwichYear(lch, lcm, lcs, ds, zc, ld, lm, ly)
	ut := LCTToUT(lch, lcm, lcs, ds, zc, ld, lm, ly)
	dj := CDToJD(aa, bb, cc) - 2415020.0
	t := (dj / 36525.0) + (ut / 876600.0)
	t2 := t * t
	a := 100.0021359 * t
	b := 360.0 * (a - math.Floor(a))

	l := 279.69668 + 0.0003025*t2 + b
	a = 99.99736042 * t
	b = 360.0 * (a - math.Floor(a))

	m1 := 358.47583 - (0.00015+0.0000033*t)*t2 + b
	ec := 0.01675104 - 0.0000418*t - 0.000000126*t2

	am := util.DegreesToRadians(m1)
	at := TrueAnomaly(am, ec)

	a = 62.55209472 * t
	b = 360.0 * (a - math.Floor(a))

	a1 := util.DegreesToRadians(153.23 + b)
	a = 125.1041894 * t
	b = 360.0 * (a - math.Floor(a))

	b1 := util.DegreesToRadians(216.57 + b)
	a = 91.56766028 * t
	b = 360.0 * (a - math.Floor(a))

	c1 := util.DegreesToRadians(312.69 + b)
	a = 1236.853095 * t
	b = 360.0 * (a - math.Floor(a))

	d1 := util.DegreesToRadians(350.74 - 0.00144*t2 + b)
	e1 := util.DegreesToRadians(231.19 + 20.2*t)
	a = 183.1353208 * t
	b = 360.0 * (a - math.Floor(a))

	d2 := 0.00134*math.Cos(a1) + 0.00154*math.Cos(b1) + 0.002*math.Cos(c1)
	d2 = d2 + 0.00179*math.Sin(d1) + 0.00178*math.Sin(e1)
	d3 := 0.00000543*math.Sin(a1) + 0.00001575*math.Sin(b1)
	d3 = d3 + 0.00001627*math.Sin(c1) + 0.00003076*math.Cos(d1)

	sr := at + util.DegreesToRadians(l-m1+d2)
	tp := 6.283185308

	sr = sr - tp*math.Floor(sr/tp)
	return Degrees(sr)
}

// TrueAnomaly solves Kepler's equation, and returns the value of the true anomaly in radians.
//
// Original macro name: TrueAnomaly
func TrueAnomaly(am float64, ec float64) float64 {
	tp := 6.283185308
	m := am - tp*math.Floor(am/tp)

	ae := m

	for true {
		d := ae - (ec * math.Sin(ae)) - m
		if math.Abs(d) < 0.000001 {
			break
		}
		d = d / (1.0 - (ec * math.Cos(ae)))
		ae = ae - d
	}

	a := math.Sqrt((1.0+ec)/(1.0-ec)) * math.Tan(ae/2.0)
	at := 2.0 * math.Atan(a)

	return at
}

// EccentricAnomaly solves Kepler's equation, and returns the value of the eccentric anomaly in radians.
//
// Original macro name: EccentricAnomaly
func EccentricAnomaly(am float64, ec float64) float64 {
	tp := 6.283185308
	m := am - tp*math.Floor(am/tp)

	ae := m

	for true {
		d := ae - (ec * math.Sin(ae)) - m

		if math.Abs(d) < 0.000001 {
			break
		}

		d = d / (1.0 - (ec * math.Cos(ae)))
		ae = ae - d
	}

	return ae
}

// Refraction calculates effects of refraction.
//
// Original macro name: Refract
func Refraction(y2 float64, sw string, pr float64, tr float64) float64 {
	y := util.DegreesToRadians(y2)

	var d float64
	if strings.ToLower(sw[0:1]) == "t" {
		d = -1.0
	} else {
		d = 1.0
	}

	if d == -1.0 {
		y3 := y
		y1 := y
		r1 := 0.0

		for true {
			y := y1 + r1
			rf := RefractL3035(pr, tr, y, d)
			if y < -0.087 {
				return 0.0
			}
			r2 := rf

			if (r2 == 0.0) || (math.Abs(r2-r1) < 0.000001) {
				q := y3
				return Degrees(q + rf)
			}

			r1 = r2
		}
	}

	rf := RefractL3035(pr, tr, y, d)

	if y < -0.087 {
		return 0.0
	}

	q := y

	return Degrees(q + rf)
}

// RefractL3035 is a helper function for Refraction()
func RefractL3035(pr float64, tr float64, y float64, d float64) float64 {
	if y < 0.2617994 {
		if y < -0.087 {
			return 0.0
		}

		yd := Degrees(y)
		a := ((0.00002*yd+0.0196)*yd + 0.1594) * pr
		b := (273.0 + tr) * ((0.0845*yd+0.505)*yd + 1.0)

		return util.DegreesToRadians(-(a / b) * d)
	}

	return -d * 0.00007888888 * pr / ((273.0 + tr) * math.Tan(y))
}

// ParallaxHA calculates corrected hour angle in decimal hours.
//
// Original macro name: ParallaxHA
func ParallaxHA(hh float64, hm float64, hs float64, dd float64, dm float64, ds float64, sw string, gp float64, ht float64, hp float64) float64 {
	a := util.DegreesToRadians(gp)
	c1 := math.Cos(a)
	s1 := math.Sin(a)

	u := math.Atan(0.996647 * s1 / c1)
	c2 := math.Cos(u)
	s2 := math.Sin(u)
	b := ht / 6378160.0

	rs := (0.996647 * s2) + (b * s1)

	rc := c2 + (b * c1)
	tp := 6.283185308

	rp := 1.0 / math.Sin(util.DegreesToRadians(hp))

	x := util.DegreesToRadians(DHToDD(HMSToDH(hh, hm, hs)))
	x1 := x
	y := util.DegreesToRadians(DMSToDD(dd, dm, ds))
	y1 := y

	var d float64
	if strings.ToLower(sw[0:1]) == "t" {
		d = 1.0
	} else {
		d = -1.0
	}

	if d == 1.0 {
		p, _ := ParallaxHAL2870(x, y, rc, rp, rs, tp)
		return DDToDH(Degrees(p))
	}

	p1 := 0.0
	q1 := 0.0
	xLoop := x
	yLoop := y
	for true {
		p, q := ParallaxHAL2870(xLoop, yLoop, rc, rp, rs, tp)
		p2 := p - xLoop
		q2 := q - yLoop

		aa := math.Abs(p2 - p1)
		bb := math.Abs(q2 - q1)

		if (aa < 0.000001) && (bb < 0.000001) {
			p = x1 - p2

			return DDToDH(Degrees(p))
		}
		xLoop = x1 - p2
		yLoop = y1 - q2
		p1 = p2
		q1 = q2
	}

	return DDToDH(Degrees(0.0))
}

// ParallaxHAL2870 is a helper function for ParallaxHA
func ParallaxHAL2870(x float64, y float64, rc float64, rp float64, rs float64, tp float64) (float64, float64) {
	cx := math.Cos(x)
	sy := math.Sin(y)
	cy := math.Cos(y)

	aa := (rc * math.Sin(x)) / ((rp * cy) - (rc * cx))

	dx := math.Atan(aa)
	p := x + dx
	cp := math.Cos(p)

	p = p - tp*math.Floor(p/tp)
	q := math.Atan(cp * (rp*sy - rs) / (rp*cy*cx - rc))

	return p, q
}

// ParallaxDec calculates corrected declination in decimal degrees.
//
// Original macro name: ParallaxDec
func ParallaxDec(hh float64, hm float64, hs float64, dd float64, dm float64, ds float64, sw string, gp float64, ht float64, hp float64) float64 {
	a := util.DegreesToRadians(gp)
	c1 := math.Cos(a)
	s1 := math.Sin(a)

	u := math.Atan(0.996647 * s1 / c1)

	c2 := math.Cos(u)
	s2 := math.Sin(u)
	b := ht / 6378160.0
	rs := (0.996647 * s2) + (b * s1)

	rc := c2 + (b * c1)
	tp := 6.283185308

	rp := 1.0 / math.Sin(util.DegreesToRadians(hp))

	x := util.DegreesToRadians(DHToDD(HMSToDH(hh, hm, hs)))
	x1 := x

	y := util.DegreesToRadians(DMSToDD(dd, dm, ds))
	y1 := y

	var d float64
	if strings.ToLower(sw[0:1]) == "t" {
		d = 1.0
	} else {
		d = -1.0
	}

	if d == 1.0 {
		_, q := ParallaxDecL2870(x, y, rc, rp, rs, tp)
		return Degrees(q)
	}

	p1 := 0.0

	xLoop := x
	yLoop := y
	for true {
		p, q := ParallaxDecL2870(xLoop, yLoop, rc, rp, rs, tp)
		p2 := p - xLoop
		q2 := q - yLoop
		aa := math.Abs(p2 - p1)
		if (aa < 0.000001) && (b < 0.000001) {
			q := y1 - q2
			return Degrees(q)
		}
		xLoop = x1 - p2
		yLoop = y1 - q2
		p1 = p2
	}

	return Degrees(0.0)
}

// ParallaxDecL2870 is a helper function for ParallaxDec.
func ParallaxDecL2870(x float64, y float64, rc float64, rp float64, rs float64, tp float64) (float64, float64) {
	cx := math.Cos(x)
	sy := math.Sin(y)
	cy := math.Cos(y)

	aa := (rc * math.Sin(x)) / ((rp * cy) - (rc * cx))
	dx := math.Atan(aa)
	p := x + dx
	cp := math.Cos(p)

	p = p - tp*math.Floor(p/tp)
	q := math.Atan(cp * (rp*sy - rs) / (rp*cy*cx - rc))

	return p, q
}
