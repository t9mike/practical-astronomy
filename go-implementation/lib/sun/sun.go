package sun

import (
	"../macros"
	"../util"
	"math"
)

// ApproximatePositionOfSun calculates approximate position of the sun for a local date and time.
//
// Arguments
//	lctHours -- Local civil time, in hours.
//	lctMinutes -- Local civil time, in minutes.
//	lctSeconds -- Local civil time, in seconds.
//	localDay -- Local date, day part.
//	localMonth -- Local date, month part.
//	localYear -- Local date, year part.
//	isDaylightSaving -- Is daylight savings in effect?
//	zoneCorrection -- Time zone correction, in hours.
//
// Returns
//	sunRAHour -- Right Ascension of Sun, hour part
//	sunRAMin -- Right Ascension of Sun, minutes part
//	sunRASec -- Right Ascension of Sun, seconds part
//	sunDecDeg -- Declination of Sun, degrees part
//	sunDecMin -- Declination of Sun, minutes part
//	sunDecSec -- Declination of Sun, seconds part
func ApproximatePositionOfSun(lctHours float64, lctMinutes float64, lctSeconds float64, localDay float64, localMonth int, localYear int, isDaylightSaving bool, zoneCorrection int) (float64, float64, float64, float64, float64, float64) {
	var daylightSaving int
	if isDaylightSaving == true {
		daylightSaving = 1
	} else {
		daylightSaving = 0
	}

	greenwichDateDay := macros.LCTGreenwichDay(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	greenwichDateMonth := macros.LCTGreenwichMonth(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	greenwichDateYear := macros.LCTGreenwichYear(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	utHours := macros.LCTToUT(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	utDays := utHours / 24.0
	jdDays := macros.CDToJD(greenwichDateDay, greenwichDateMonth, greenwichDateYear) + utDays
	dDays := jdDays - macros.CDToJD(0, 1, 2010)
	nDeg := 360.0 * dDays / 365.242191
	mDeg1 := nDeg + macros.SunMeanEclipticLongitude(0, 1, 2010) - macros.SunPerigee(0, 1, 2010)
	mDeg2 := mDeg1 - 360.0*math.Floor(mDeg1/360.0)
	eCDeg := 360.0 * macros.SunEccentricity(0, 1, 2010) * math.Sin(util.DegreesToRadians(mDeg2)) / math.Pi
	lSDeg1 := nDeg + eCDeg + macros.SunMeanEclipticLongitude(0, 1, 2010)
	lSDeg2 := lSDeg1 - 360.0*math.Floor(lSDeg1/360.0)
	raDeg := macros.EclipticRightAscension(lSDeg2, 0, 0, 0, 0, 0, greenwichDateDay, greenwichDateMonth, greenwichDateYear)
	raHours := macros.DDToDH(raDeg)
	decDeg := macros.EclipticDeclination(lSDeg2, 0, 0, 0, 0, 0, greenwichDateDay, greenwichDateMonth, greenwichDateYear)

	sunRAHour := macros.DHHour(raHours)
	sunRAMin := macros.DHMin(raHours)
	sunRASec := macros.DHSec(raHours)
	sunDecDeg := macros.DDDeg(decDeg)
	sunDecMin := macros.DDMin(decDeg)
	sunDecSec := macros.DDSec(decDeg)

	return float64(sunRAHour), float64(sunRAMin), sunRASec, sunDecDeg, sunDecMin, sunDecSec
}

// PrecisePositionOfSun calculates precise position of the sun for a local date and time.
//
// Arguments
//	lctHours -- Local civil time, in hours.
//	lctMinutes -- Local civil time, in minutes.
//	lctSeconds -- Local civil time, in seconds.
//	localDay -- Local date, day part.
//	localMonth -- Local date, month part.
//	localYear -- Local date, year part.
//	isDaylightSaving -- Is daylight savings in effect?
//	zoneCorrection -- Time zone correction, in hours.
//
// ## Returns
//	sunRAHour -- Right Ascension of Sun, hour part
//	sunRAMin -- Right Ascension of Sun, minutes part
//	sunRASec -- Right Ascension of Sun, seconds part
//	sunDecDeg -- Declination of Sun, degrees part
//	sunDecMin -- Declination of Sun, minutes part
//	sunDecSec -- Declination of Sun, seconds part
func PrecisePositionOfSun(lctHours float64, lctMinutes float64, lctSeconds float64, localDay float64, localMonth int, localYear int, isDaylightSaving bool, zoneCorrection int) (float64, float64, float64, float64, float64, float64) {
	var daylightSaving int
	if isDaylightSaving == true {
		daylightSaving = 1
	} else {
		daylightSaving = 0
	}

	gDay := macros.LCTGreenwichDay(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	gMonth := macros.LCTGreenwichMonth(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	gYear := macros.LCTGreenwichYear(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	sunEclipticLongitudeDeg := macros.SunEclipticLongitude(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	raDeg := macros.EclipticRightAscension(sunEclipticLongitudeDeg, 0.0, 0.0, 0.0, 0.0, 0.0, gDay, gMonth, gYear)
	raHours := macros.DDToDH(raDeg)
	decDeg := macros.EclipticDeclination(sunEclipticLongitudeDeg, 0.0, 0.0, 0.0, 0.0, 0.0, gDay, gMonth, gYear)

	sunRAHour := macros.DHHour(raHours)
	sunRAMin := macros.DHMin(raHours)
	sunRASec := macros.DHSec(raHours)
	sunDecDeg := macros.DDDeg(decDeg)
	sunDecMin := macros.DDMin(decDeg)
	sunDecSec := macros.DDSec(decDeg)

	return float64(sunRAHour), float64(sunRAMin), sunRASec, sunDecDeg, sunDecMin, sunDecSec
}
