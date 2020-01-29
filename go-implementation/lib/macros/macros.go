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

// SunAngularDiameter calculates Sun's angular diameter in decimal degrees
//
// Original macro name: SunDia
func SunAngularDiameter(lch float64, lcm float64, lcs float64, ds int, zc int, ld float64, lm int, ly int) float64 {
	a := SunDistance(lch, lcm, lcs, ds, zc, ld, lm, ly)

	return 0.533128 / a
}

// SunDistance calculates Sun's distance from the Earth in astronomical units
//
// Original macro name: SunDist
func SunDistance(lch float64, lcm float64, lcs float64, ds int, zc int, ld float64, lm int, ly int) float64 {
	aa := LCTGreenwichDay(lch, lcm, lcs, ds, zc, ld, lm, ly)
	bb := LCTGreenwichMonth(lch, lcm, lcs, ds, zc, ld, lm, ly)
	cc := LCTGreenwichYear(lch, lcm, lcs, ds, zc, ld, lm, ly)
	ut := LCTToUT(lch, lcm, lcs, ds, zc, ld, lm, ly)
	dj := CDToJD(aa, bb, cc) - 2415020.0

	t := (dj / 36525.0) + (ut / 876600.0)
	t2 := t * t

	a := 100.0021359 * t
	b := 360.0 * (a - math.Floor(a))
	a = 99.99736042 * t
	b = 360.0 * (a - math.Floor(a))
	m1 := 358.47583 - (0.00015+0.0000033*t)*t2 + b
	ec := 0.01675104 - 0.0000418*t - 0.000000126*t2

	am := util.DegreesToRadians(m1)
	ae := EccentricAnomaly(am, ec)

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
	a = 183.1353208 * t
	b = 360.0 * (a - math.Floor(a))
	h1 := util.DegreesToRadians(353.4 + b)

	d3 := (0.00000543*math.Sin(a1) + 0.00001575*math.Sin(b1)) + (0.00001627*math.Sin(c1) + 0.00003076*math.Cos(d1)) + (0.00000927 * math.Sin(h1))

	return 1.0000002*(1.0-ec*math.Cos(ae)) + d3
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

// MoonGeocentricEclipticLongitude calculates geocentric ecliptic longitude for the Moon
//
// Original macro name: MoonLong
func MoonGeocentricEclipticLongitude(lh float64, lm float64, ls float64, ds int, zc int, dy float64, mn int, yr int) float64 {
	ut := LCTToUT(lh, lm, ls, ds, zc, dy, mn, yr)
	gd := LCTGreenwichDay(lh, lm, ls, ds, zc, dy, mn, yr)
	gm := LCTGreenwichMonth(lh, lm, ls, ds, zc, dy, mn, yr)
	gy := LCTGreenwichYear(lh, lm, ls, ds, zc, dy, mn, yr)
	t := ((CDToJD(gd, gm, gy) - 2415020.0) / 36525.0) + (ut / 876600.0)
	t2 := t * t

	m1 := 27.32158213
	m2 := 365.2596407
	m3 := 27.55455094
	m4 := 29.53058868
	m5 := 27.21222039
	m6 := 6798.363307
	q := CDToJD(gd, gm, gy) - 2415020.0 + (ut / 24.0)
	m1 = q / m1
	m2 = q / m2
	m3 = q / m3
	m4 = q / m4
	m5 = q / m5
	m6 = q / m6
	m1 = 360.0 * (m1 - math.Floor(m1))
	m2 = 360.0 * (m2 - math.Floor(m2))
	m3 = 360.0 * (m3 - math.Floor(m3))
	m4 = 360.0 * (m4 - math.Floor(m4))
	m5 = 360.0 * (m5 - math.Floor(m5))
	m6 = 360.0 * (m6 - math.Floor(m6))

	ml := 270.434164 + m1 - (0.001133-0.0000019*t)*t2
	ms := 358.475833 + m2 - (0.00015+0.0000033*t)*t2
	md := 296.104608 + m3 + (0.009192+0.0000144*t)*t2
	me1 := 350.737486 + m4 - (0.001436-0.0000019*t)*t2
	mf := 11.250889 + m5 - (0.003211+0.0000003*t)*t2
	na := 259.183275 - m6 + (0.002078+0.0000022*t)*t2
	a := util.DegreesToRadians(51.2 + 20.2*t)
	s1 := math.Sin(a)
	s2 := math.Sin(util.DegreesToRadians(na))
	b := 346.56 + (132.87-0.0091731*t)*t
	s3 := 0.003964 * math.Sin(util.DegreesToRadians(b))
	c := util.DegreesToRadians(na + 275.05 - 2.3*t)
	s4 := math.Sin(c)
	ml = ml + 0.000233*s1 + s3 + 0.001964*s2
	ms = ms - 0.001778*s1
	md = md + 0.000817*s1 + s3 + 0.002541*s2
	mf = mf + s3 - 0.024691*s2 - 0.004328*s4
	me1 = me1 + 0.002011*s1 + s3 + 0.001964*s2
	e := 1.0 - (0.002495+0.00000752*t)*t
	e2 := e * e
	ml = util.DegreesToRadians(ml)
	ms = util.DegreesToRadians(ms)
	me1 = util.DegreesToRadians(me1)
	mf = util.DegreesToRadians(mf)
	md = util.DegreesToRadians(md)

	l := 6.28875*math.Sin(md) + 1.274018*math.Sin(2.0*me1-md)
	l = l + 0.658309*math.Sin(2.0*me1) + 0.213616*math.Sin(2.0*md)
	l = l - e*0.185596*math.Sin(ms) - 0.114336*math.Sin(2.0*mf)
	l = l + 0.058793*math.Sin(2.0*(me1-md))
	l = l + 0.057212*e*math.Sin(2.0*me1-ms-md) + 0.05332*math.Sin(2.0*me1+md)
	l = l + 0.045874*e*math.Sin(2.0*me1-ms) + 0.041024*e*math.Sin(md-ms)
	l = l - 0.034718*math.Sin(me1) - e*0.030465*math.Sin(ms+md)
	l = l + 0.015326*math.Sin(2.0*(me1-mf)) - 0.012528*math.Sin(2.0*mf+md)
	l = l - 0.01098*math.Sin(2.0*mf-md) + 0.010674*math.Sin(4.0*me1-md)
	l = l + 0.010034*math.Sin(3.0*md) + 0.008548*math.Sin(4.0*me1-2.0*md)
	l = l - e*0.00791*math.Sin(ms-md+2.0*me1) - e*0.006783*math.Sin(2.0*me1+ms)
	l = l + 0.005162*math.Sin(md-me1) + e*0.005*math.Sin(ms+me1)
	l = l + 0.003862*math.Sin(4.0*me1) + e*0.004049*math.Sin(md-ms+2.0*me1)
	l = l + 0.003996*math.Sin(2.0*(md+me1)) + 0.003665*math.Sin(2.0*me1-3.0*md)
	l = l + e*0.002695*math.Sin(2.0*md-ms) + 0.002602*math.Sin(md-2.0*(mf+me1))
	l = l + e*0.002396*math.Sin(2.0*(me1-md)-ms) - 0.002349*math.Sin(md+me1)
	l = l + e2*0.002249*math.Sin(2.0*(me1-ms)) - e*0.002125*math.Sin(2.0*md+ms)
	l = l - e2*0.002079*math.Sin(2.0*ms) + e2*0.002059*math.Sin(2.0*(me1-ms)-md)
	l = l - 0.001773*math.Sin(md+2.0*(me1-mf)) - 0.001595*math.Sin(2.0*(mf+me1))
	l = l + e*0.00122*math.Sin(4.0*me1-ms-md) - 0.00111*math.Sin(2.0*(md+mf))
	l = l + 0.000892*math.Sin(md-3.0*me1) - e*0.000811*math.Sin(ms+md+2.0*me1)
	l = l + e*0.000761*math.Sin(4.0*me1-ms-2.0*md)
	l = l + e2*0.000704*math.Sin(md-2.0*(ms+me1))
	l = l + e*0.000693*math.Sin(ms-2.0*(md-me1))
	l = l + e*0.000598*math.Sin(2.0*(me1-mf)-ms)
	l = l + 0.00055*math.Sin(md+4.0*me1) + 0.000538*math.Sin(4.0*md)
	l = l + e*0.000521*math.Sin(4.0*me1-ms) + 0.000486*math.Sin(2.0*md-me1)
	l = l + e2*0.000717*math.Sin(md-2.0*ms)
	mm := Unwind(ml + util.DegreesToRadians(l))

	return Degrees(mm)
}

// MoonGeocentricEclipticLatitude calculates geocentric ecliptic latitude for the Moon
//
// Original macro name: MoonLat
func MoonGeocentricEclipticLatitude(lh float64, lm float64, ls float64, ds int, zc int, dy float64, mn int, yr int) float64 {
	ut := LCTToUT(lh, lm, ls, ds, zc, dy, mn, yr)
	gd := LCTGreenwichDay(lh, lm, ls, ds, zc, dy, mn, yr)
	gm := LCTGreenwichMonth(lh, lm, ls, ds, zc, dy, mn, yr)
	gy := LCTGreenwichYear(lh, lm, ls, ds, zc, dy, mn, yr)
	t := ((CDToJD(gd, gm, gy) - 2415020.0) / 36525.0) + (ut / 876600.0)
	t2 := t * t

	m1 := 27.32158213
	m2 := 365.2596407
	m3 := 27.55455094
	m4 := 29.53058868
	m5 := 27.21222039
	m6 := 6798.363307
	q := CDToJD(gd, gm, gy) - 2415020.0 + (ut / 24.0)
	m1 = q / m1
	m2 = q / m2
	m3 = q / m3
	m4 = q / m4
	m5 = q / m5
	m6 = q / m6
	m1 = 360.0 * (m1 - math.Floor(m1))
	m2 = 360.0 * (m2 - math.Floor(m2))
	m3 = 360.0 * (m3 - math.Floor(m3))
	m4 = 360.0 * (m4 - math.Floor(m4))
	m5 = 360.0 * (m5 - math.Floor(m5))
	m6 = 360.0 * (m6 - math.Floor(m6))

	ml := 270.434164 + m1 - (0.001133-0.0000019*t)*t2
	ms := 358.475833 + m2 - (0.00015+0.0000033*t)*t2
	md := 296.104608 + m3 + (0.009192+0.0000144*t)*t2
	me1 := 350.737486 + m4 - (0.001436-0.0000019*t)*t2
	mf := 11.250889 + m5 - (0.003211+0.0000003*t)*t2
	na := 259.183275 - m6 + (0.002078+0.0000022*t)*t2
	a := util.DegreesToRadians(51.2 + 20.2*t)
	s1 := math.Sin(a)
	s2 := math.Sin(util.DegreesToRadians(na))
	b := 346.56 + (132.87-0.0091731*t)*t
	s3 := 0.003964 * math.Sin(util.DegreesToRadians(b))
	c := util.DegreesToRadians(na + 275.05 - 2.3*t)
	s4 := math.Sin(c)
	ml = ml + 0.000233*s1 + s3 + 0.001964*s2
	ms = ms - 0.001778*s1
	md = md + 0.000817*s1 + s3 + 0.002541*s2
	mf = mf + s3 - 0.024691*s2 - 0.004328*s4
	me1 = me1 + 0.002011*s1 + s3 + 0.001964*s2
	e := 1.0 - (0.002495+0.00000752*t)*t
	e2 := e * e
	ms = util.DegreesToRadians(ms)
	na = util.DegreesToRadians(na)
	me1 = util.DegreesToRadians(me1)
	mf = util.DegreesToRadians(mf)
	md = util.DegreesToRadians(md)

	g := 5.128189*math.Sin(mf) + 0.280606*math.Sin(md+mf)
	g = g + 0.277693*math.Sin(md-mf) + 0.173238*math.Sin(2.0*me1-mf)
	g = g + 0.055413*math.Sin(2.0*me1+mf-md) + 0.046272*math.Sin(2.0*me1-mf-md)
	g = g + 0.032573*math.Sin(2.0*me1+mf) + 0.017198*math.Sin(2.0*md+mf)
	g = g + 0.009267*math.Sin(2.0*me1+md-mf) + 0.008823*math.Sin(2.0*md-mf)
	g = g + e*0.008247*math.Sin(2.0*me1-ms-mf) + 0.004323*math.Sin(2.0*(me1-md)-mf)
	g = g + 0.0042*math.Sin(2.0*me1+mf+md) + e*0.003372*math.Sin(mf-ms-2.0*me1)
	g = g + e*0.002472*math.Sin(2.0*me1+mf-ms-md)
	g = g + e*0.002222*math.Sin(2.0*me1+mf-ms)
	g = g + e*0.002072*math.Sin(2.0*me1-mf-ms-md)
	g = g + e*0.001877*math.Sin(mf-ms+md) + 0.001828*math.Sin(4.0*me1-mf-md)
	g = g - e*0.001803*math.Sin(mf+ms) - 0.00175*math.Sin(3.0*mf)
	g = g + e*0.00157*math.Sin(md-ms-mf) - 0.001487*math.Sin(mf+me1)
	g = g - e*0.001481*math.Sin(mf+ms+md) + e*0.001417*math.Sin(mf-ms-md)
	g = g + e*0.00135*math.Sin(mf-ms) + 0.00133*math.Sin(mf-me1)
	g = g + 0.001106*math.Sin(mf+3.0*md) + 0.00102*math.Sin(4.0*me1-mf)
	g = g + 0.000833*math.Sin(mf+4.0*me1-md) + 0.000781*math.Sin(md-3.0*mf)
	g = g + 0.00067*math.Sin(mf+4.0*me1-2.0*md) + 0.000606*math.Sin(2.0*me1-3.0*mf)
	g = g + 0.000597*math.Sin(2.0*(me1+md)-mf)
	g = g + e*0.000492*math.Sin(2.0*me1+md-ms-mf) + 0.00045*math.Sin(2.0*(md-me1)-mf)
	g = g + 0.000439*math.Sin(3.0*md-mf) + 0.000423*math.Sin(mf+2.0*(me1+md))
	g = g + 0.000422*math.Sin(2.0*me1-mf-3.0*md) - e*0.000367*math.Sin(ms+mf+2.0*me1-md)
	g = g - e*0.000353*math.Sin(ms+mf+2.0*me1) + 0.000331*math.Sin(mf+4.0*me1)
	g = g + e*0.000317*math.Sin(2.0*me1+mf-ms+md)
	g = g + e2*0.000306*math.Sin(2.0*(me1-ms)-mf) - 0.000283*math.Sin(md+3.0*mf)
	w1 := 0.0004664 * math.Cos(na)
	w2 := 0.0000754 * math.Cos(c)
	bm := util.DegreesToRadians(g) * (1.0 - w1 - w2)

	return Degrees(bm)
}

// MoonHorizontalParallax calculates horizontal parallax for the Moon
//
// Original macro name: MoonHP
func MoonHorizontalParallax(lh float64, lm float64, ls float64, ds int, zc int, dy float64, mn int, yr int) float64 {
	ut := LCTToUT(lh, lm, ls, ds, zc, dy, mn, yr)
	gd := LCTGreenwichDay(lh, lm, ls, ds, zc, dy, mn, yr)
	gm := LCTGreenwichMonth(lh, lm, ls, ds, zc, dy, mn, yr)
	gy := LCTGreenwichYear(lh, lm, ls, ds, zc, dy, mn, yr)
	t := ((CDToJD(gd, gm, gy) - 2415020.0) / 36525.0) + (ut / 876600.0)
	t2 := t * t

	m1 := 27.32158213
	m2 := 365.2596407
	m3 := 27.55455094
	m4 := 29.53058868
	m5 := 27.21222039
	m6 := 6798.363307
	q := CDToJD(gd, gm, gy) - 2415020.0 + (ut / 24.0)
	m1 = q / m1
	m2 = q / m2
	m3 = q / m3
	m4 = q / m4
	m5 = q / m5
	m6 = q / m6
	m1 = 360.0 * (m1 - math.Floor(m1))
	m2 = 360.0 * (m2 - math.Floor(m2))
	m3 = 360.0 * (m3 - math.Floor(m3))
	m4 = 360.0 * (m4 - math.Floor(m4))
	m5 = 360.0 * (m5 - math.Floor(m5))
	m6 = 360.0 * (m6 - math.Floor(m6))

	ml := 270.434164 + m1 - (0.001133-0.0000019*t)*t2
	ms := 358.475833 + m2 - (0.00015+0.0000033*t)*t2
	md := 296.104608 + m3 + (0.009192+0.0000144*t)*t2
	me1 := 350.737486 + m4 - (0.001436-0.0000019*t)*t2
	mf := 11.250889 + m5 - (0.003211+0.0000003*t)*t2
	na := 259.183275 - m6 + (0.002078+0.0000022*t)*t2
	a := util.DegreesToRadians(51.2 + 20.2*t)
	s1 := math.Sin(a)
	s2 := math.Sin(util.DegreesToRadians(na))
	b := 346.56 + (132.87-0.0091731*t)*t
	s3 := 0.003964 * math.Sin(util.DegreesToRadians(b))
	c := util.DegreesToRadians(na + 275.05 - 2.3*t)
	s4 := math.Sin(c)
	ml = ml + 0.000233*s1 + s3 + 0.001964*s2
	ms = ms - 0.001778*s1
	md = md + 0.000817*s1 + s3 + 0.002541*s2
	mf = mf + s3 - 0.024691*s2 - 0.004328*s4
	me1 = me1 + 0.002011*s1 + s3 + 0.001964*s2
	e := 1.0 - (0.002495+0.00000752*t)*t
	e2 := e * e
	ms = util.DegreesToRadians(ms)
	me1 = util.DegreesToRadians(me1)
	mf = util.DegreesToRadians(mf)
	md = util.DegreesToRadians(md)

	pm := 0.950724 + 0.051818*math.Cos(md) + 0.009531*math.Cos(2.0*me1-md)
	pm = pm + 0.007843*math.Cos(2.0*me1) + 0.002824*math.Cos(2.0*md)
	pm = pm + 0.000857*math.Cos(2.0*me1+md) + e*0.000533*math.Cos(2.0*me1-ms)
	pm = pm + e*0.000401*math.Cos(2.0*me1-md-ms)
	pm = pm + e*0.00032*math.Cos(md-ms) - 0.000271*math.Cos(me1)
	pm = pm - e*0.000264*math.Cos(ms+md) - 0.000198*math.Cos(2.0*mf-md)
	pm = pm + 0.000173*math.Cos(3.0*md) + 0.000167*math.Cos(4.0*me1-md)
	pm = pm - e*0.000111*math.Cos(ms) + 0.000103*math.Cos(4.0*me1-2.0*md)
	pm = pm - 0.000084*math.Cos(2.0*md-2.0*me1) - e*0.000083*math.Cos(2.0*me1+ms)
	pm = pm + 0.000079*math.Cos(2.0*me1+2.0*md) + 0.000072*math.Cos(4.0*me1)
	pm = pm + e*0.000064*math.Cos(2.0*me1-ms+md) - e*0.000063*math.Cos(2.0*me1+ms-md)
	pm = pm + e*0.000041*math.Cos(ms+me1) + e*0.000035*math.Cos(2.0*md-ms)
	pm = pm - 0.000033*math.Cos(3.0*md-2.0*me1) - 0.00003*math.Cos(md+me1)
	pm = pm - 0.000029*math.Cos(2.0*(mf-me1)) - e*0.000029*math.Cos(2.0*md+ms)
	pm = pm + e2*0.000026*math.Cos(2.0*(me1-ms)) - 0.000023*math.Cos(2.0*(mf-me1)+md)
	pm = pm + e*0.000019*math.Cos(4.0*me1-ms-md)

	return pm
}

// Unwind converts angle in radians to equivalent angle in degrees.
//
// Original macro name: Unwind
func Unwind(w float64) float64 {
	return UnwindRad(w)
}

// UnwindDeg converts angle in degrees to equivalent angle in the range 0 to 360 degrees.
//
// Original macro name: UnwindDeg
func UnwindDeg(w float64) float64 {
	return w - 360.0*math.Floor(w/360.0)
}

// UnwindRad converts angle in radians to equivalent angle in degrees.
//
// Original macro name: UnwindRad
func UnwindRad(w float64) float64 {
	return w - 6.283185308*math.Floor(w/6.283185308)
}

// EclipticRightAscension calculates right ascension (in degrees) for the ecliptic.
//
// Original macro name: ECRA
func EclipticRightAscension(eld float64, elm float64, els float64, bd float64, bm float64, bs float64, gd float64, gm int, gy int) float64 {
	a := util.DegreesToRadians(DMSToDD(eld, elm, els))
	b := util.DegreesToRadians(DMSToDD(bd, bm, bs))
	c := util.DegreesToRadians(Obliq(gd, gm, gy))
	d := math.Sin(a)*math.Cos(c) - math.Tan(b)*math.Sin(c)
	e := math.Cos(a)
	f := Degrees(math.Atan2(d, e))

	return f - 360.0*math.Floor(f/360.0)
}

// EclipticDeclination calculates declination (in degrees) for the ecliptic.
//
// Original macro name: ECDec
func EclipticDeclination(eld float64, elm float64, els float64, bd float64, bm float64, bs float64, gd float64, gm int, gy int) float64 {
	a := util.DegreesToRadians(DMSToDD(eld, elm, els))
	b := util.DegreesToRadians(DMSToDD(bd, bm, bs))
	c := util.DegreesToRadians(Obliq(gd, gm, gy))
	d := math.Sin(b)*math.Cos(c) + math.Cos(b)*math.Sin(c)*math.Sin(a)

	return Degrees(math.Asin(d))
}

// SunTrueAnomaly calculates Sun's true anomaly, i.e., how much its orbit deviates from a true circle to an ellipse.
//
// Original macro name: SunTrueAnomaly
func SunTrueAnomaly(lch float64, lcm float64, lcs float64, ds int, zc int, ld float64, lm int, ly int) float64 {
	aa := LCTGreenwichDay(lch, lcm, lcs, ds, zc, ld, lm, ly)
	bb := LCTGreenwichMonth(lch, lcm, lcs, ds, zc, ld, lm, ly)
	cc := LCTGreenwichYear(lch, lcm, lcs, ds, zc, ld, lm, ly)
	ut := LCTToUT(lch, lcm, lcs, ds, zc, ld, lm, ly)
	dj := CDToJD(aa, bb, cc) - 2415020.0

	t := (dj / 36525.0) + (ut / 876600.0)
	t2 := t * t

	a := 100.0021359 * t
	b := 360.0 * (a - math.Floor(a))

	a = 99.99736042 * t
	b = 360.0 * (a - math.Floor(a))

	m1 := 358.47583 - (0.00015+0.0000033*t)*t2 + b
	ec := 0.01675104 - 0.0000418*t - 0.000000126*t2

	am := util.DegreesToRadians(m1)

	return Degrees(TrueAnomaly(am, ec))
}

// SunMeanEclipticLongitude calculates mean ecliptic longitude of the Sun at the epoch
//
// Original macro name: SunElong
func SunMeanEclipticLongitude(gd float64, gm int, gy int) float64 {
	t := (CDToJD(gd, gm, gy) - 2415020.0) / 36525.0
	t2 := t * t
	x := 279.6966778 + 36000.76892*t + 0.0003025*t2

	return x - 360.0*math.Floor(x/360.0)
}

// SunPerigee calculates longitude of the Sun at perigee
//
// Original macro name: SunPeri
func SunPerigee(gd float64, gm int, gy int) float64 {
	t := (CDToJD(gd, gm, gy) - 2415020.0) / 36525.0
	t2 := t * t
	x := 281.2208444 + 1.719175*t + 0.000452778*t2

	return x - 360.0*math.Floor(x/360.0)
}

// SunEccentricity calculates eccentricity of the Sun-Earth orbit
//
// Original macro name: SunEcc
func SunEccentricity(gd float64, gm int, gy int) float64 {
	t := (CDToJD(gd, gm, gy) - 2415020.0) / 36525.0
	t2 := t * t

	return 0.01675104 - 0.0000418*t - 0.000000126*t2
}
