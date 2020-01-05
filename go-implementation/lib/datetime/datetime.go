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
	return macros.HMSDH(hours, minutes, seconds)
}

// DecimalHoursToCivilTime converts decimal hours to civil time.
//
// Returns hours, minutes, and seconds.
func DecimalHoursToCivilTime(decimalHours float64) (int, int, int) {
	hours := macros.DHHour(decimalHours)
	minutes := macros.DHMin(decimalHours)
	seconds := macros.DHSec(decimalHours)

	return hours, minutes, int(seconds)
}

// LocalCivilTimeToUniversalTime converts local Civil Time to Universal Time
//
// Returns UT hours, UT mins, UT secs, GW day, GW month, GW year
func LocalCivilTimeToUniversalTime(
	lctHours float64,
	lctMinutes float64,
	lctSeconds float64,
	isDayLightSavings bool,
	zoneCorrection int,
	localDay float64,
	localMonth int,
	localYear int) (int, int, int, int, int, int) {
	lct := CivilTimeToDecimalHours(lctHours, lctMinutes, lctSeconds)

	daylightSavingsOffset := 0
	if isDayLightSavings == true {
		daylightSavingsOffset = 1
	}

	utInterim := lct - float64(daylightSavingsOffset) - float64(zoneCorrection)
	gdayInterim := float64(localDay) + (utInterim / 24)

	jd := macros.CDJD(gdayInterim, localMonth, localYear)

	gDay := float64(macros.JDCDay(jd))
	gMonth := macros.JDCMonth(jd)
	gYear := macros.JDCYear(jd)

	ut := 24 * (gDay - math.Floor(gDay))

	return macros.DHHour(ut), macros.DHMin(ut), int(macros.DHSec(ut)), int(math.Floor(gDay)), gMonth, gYear
}

// UniversalTimeToLocalCivilTime converts Universal Time to local Civil Time
//
// Returns LCT hours, LCT minutes, LCT seconds, day, month, year
func UniversalTimeToLocalCivilTime(utHours float64, utMinutes float64, utSeconds float64, isDaylightSavings bool, zoneCorrection int, gwDay int, gwMonth int, gwYear int) (int, int, int, int, int, int) {
	var dstValue int
	if isDaylightSavings == true {
		dstValue = 1
	} else {
		dstValue = 0
	}

	ut := macros.HMSDH(utHours, utMinutes, utSeconds)
	zoneTime := ut + float64(zoneCorrection)
	localTime := zoneTime + float64(dstValue)
	localJDPlusLocalTime := macros.CDJD(float64(gwDay), gwMonth, gwYear) + (localTime / 24)
	localDay := macros.JDCDay(localJDPlusLocalTime)
	integerDay := math.Floor(localDay)
	localMonth := macros.JDCMonth(localJDPlusLocalTime)
	localYear := macros.JDCYear(localJDPlusLocalTime)

	lct := 24.0 * (localDay - integerDay)

	return macros.DHHour(lct), macros.DHMin(lct), int(macros.DHSec(lct)), int(integerDay), localMonth, localYear
}
