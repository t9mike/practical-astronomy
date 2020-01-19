package coordinates

import (
	"../macros"
	"../util"
	"math"
)

// AngleToDecimalDegrees converts an Angle (degrees, minutes, and seconds) to Decimal Degrees
func AngleToDecimalDegrees(degrees float64, minutes float64, seconds float64) float64 {
	a := math.Abs(seconds) / 60
	b := (math.Abs(minutes) + a) / 60
	c := math.Abs(degrees) + b
	var d float64
	if degrees < 0 || minutes < 0 || seconds < 0 {
		d = -c
	} else {
		d = c
	}

	return d
}

// DecimalDegreesToAngle converts Decimal Degrees to an Angle (degrees, minutes, and seconds)
//
// Returns degrees, minutes, seconds
func DecimalDegreesToAngle(decimalDegrees float64) (float64, float64, float64) {
	unsignedDecimal := math.Abs(decimalDegrees)
	totalSeconds := unsignedDecimal * 3600
	seconds2Dp := util.RoundFloat64(math.Mod(totalSeconds, 60), 2)

	var correctedSeconds float64
	if seconds2Dp == 60 {
		correctedSeconds = 0.0
	} else {
		correctedSeconds = seconds2Dp
	}

	var correctedRemainder float64
	if seconds2Dp == 60 {
		correctedRemainder = totalSeconds + 60
	} else {
		correctedRemainder = totalSeconds
	}
	minutes := math.Mod(math.Floor(correctedRemainder/60), 60)
	unsignedDegrees := math.Floor(correctedRemainder / 3600)

	var signedDegrees float64
	if decimalDegrees < 0 {
		signedDegrees = -1.0 * unsignedDegrees
	} else {
		signedDegrees = unsignedDegrees
	}

	return signedDegrees, minutes, math.Floor(correctedSeconds)
}

// RightAscensionToHourAngle converts Right Ascension to Hour Angle
func RightAscensionToHourAngle(raHours float64, raMinutes float64, raSeconds float64, lctHours float64, lctMinutes float64, lctSeconds float64, isDaylightSaving bool, zoneCorrection int, localDay float64, localMonth int, localYear int, geographicalLongitude float64) (float64, float64, float64) {
	var daylightSaving int
	if isDaylightSaving == true {
		daylightSaving = 1
	} else {
		daylightSaving = 0
	}

	hourAngle := macros.RAToHA(raHours, raMinutes, raSeconds, lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear, geographicalLongitude)

	hourAngleHours := macros.DHHour(hourAngle)
	hourAngleMinutes := macros.DHMin(hourAngle)
	hourAngleSeconds := macros.DHSec(hourAngle)

	return float64(hourAngleHours), float64(hourAngleMinutes), hourAngleSeconds
}

// HourAngleToRightAscension converts Hour Angle to Right Ascension
func HourAngleToRightAscension(hourAngleHours float64, hourAngleMinutes float64, hourAngleSeconds float64, lctHours float64, lctMinutes float64, lctSeconds float64, isDaylightSaving bool, zoneCorrection int, localDay float64, localMonth int, localYear int, geographicalLongitude float64) (float64, float64, float64) {
	var daylightSaving int
	if isDaylightSaving == true {
		daylightSaving = 1
	} else {
		daylightSaving = 0
	}

	rightAscension := macros.HAToRA(hourAngleHours, hourAngleMinutes, hourAngleSeconds, lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear, geographicalLongitude)

	rightAscensionHours := macros.DHHour(rightAscension)
	rightAscensionMinutes := macros.DHMin(rightAscension)
	rightAscensionSeconds := macros.DHSec(rightAscension)

	return float64(rightAscensionHours), float64(rightAscensionMinutes), rightAscensionSeconds
}

// EquatorialCoordinatesToHorizonCoordinates converts Equatorial Coordinates to Horizon Coordinates
func EquatorialCoordinatesToHorizonCoordinates(hourAngleHours float64, hourAngleMinutes float64, hourAngleSeconds float64, declinationDegrees float64, declinationMinutes float64, declinationSeconds float64, geographicalLatitude float64) (float64, float64, float64, float64, float64, float64) {
	azimuthInDecimalDegrees := macros.EqToAz(hourAngleHours, hourAngleMinutes, hourAngleSeconds, declinationDegrees, declinationMinutes, declinationSeconds, geographicalLatitude)

	altitudeInDecimalDegrees := macros.EqToAlt(hourAngleHours, hourAngleMinutes, hourAngleSeconds, declinationDegrees, declinationMinutes, declinationSeconds, geographicalLatitude)

	azimuthDegrees := macros.DDDeg(azimuthInDecimalDegrees)
	azimuthMinutes := macros.DDMin(azimuthInDecimalDegrees)
	azimuthSeconds := macros.DDSec(azimuthInDecimalDegrees)

	altitudeDegrees := macros.DDDeg(altitudeInDecimalDegrees)
	altitudeMinutes := macros.DDMin(altitudeInDecimalDegrees)
	altitudeSeconds := macros.DDSec(altitudeInDecimalDegrees)

	return azimuthDegrees, azimuthMinutes, azimuthSeconds, altitudeDegrees, altitudeMinutes, altitudeSeconds
}

// HorizonCoordinatesToEquatorialCoordinates converts Horizon Coordinates to Equatorial Coordinates
func HorizonCoordinatesToEquatorialCoordinates(azimuthDegrees float64, azimuthMinutes float64, azimuthSeconds float64, altitudeDegrees float64, altitudeMinutes float64, altitudeSeconds float64, geographicalLatitude float64) (float64, float64, float64, float64, float64, float64) {
	hourAngleInDecimalDegrees := macros.HorToHA(azimuthDegrees, azimuthMinutes, azimuthSeconds, altitudeDegrees, altitudeMinutes, altitudeSeconds, geographicalLatitude)

	declinationInDecimalDegrees := macros.HorToDec(azimuthDegrees, azimuthMinutes, azimuthSeconds, altitudeDegrees, altitudeMinutes, altitudeSeconds, geographicalLatitude)

	hourAngleHours := macros.DHHour(hourAngleInDecimalDegrees)
	hourAngleMinutes := macros.DHMin(hourAngleInDecimalDegrees)
	hourAngleSeconds := macros.DHSec(hourAngleInDecimalDegrees)

	declinationDegrees := macros.DDDeg(declinationInDecimalDegrees)
	declinationMinutes := macros.DDMin(declinationInDecimalDegrees)
	declinationSeconds := macros.DDSec(declinationInDecimalDegrees)

	return float64(hourAngleHours), float64(hourAngleMinutes), hourAngleSeconds, declinationDegrees, declinationMinutes, declinationSeconds
}

// MeanObliquityOfTheEcliptic calculates Mean Obliquity of the Ecliptic for a Greenwich Date
func MeanObliquityOfTheEcliptic(greenwichDay float64, greenwichMonth int, greenwichYear int) float64 {
	jd := macros.CDToJD(greenwichDay, greenwichMonth, greenwichYear)
	mjd := jd - 2451545.0
	t := mjd / 36525.0
	de1 := t * (46.815 + t*(0.0006-(t*0.00181)))
	de2 := de1 / 3600.0

	return 23.439292 - de2
}

// EclipticCoordinateToEquatorialCoordinate converts Ecliptic Coordinates to Equatorial Coordinates
func EclipticCoordinateToEquatorialCoordinate(eclipticLongitudeDegrees float64, eclipticLongitudeMinutes float64, eclipticLongitudeSeconds float64, eclipticLatitudeDegrees float64, eclipticLatitudeMinutes float64, eclipticLatitudeSeconds float64, greenwichDay float64, greenwichMonth int, greenwichYear int) (float64, float64, float64, float64, float64, float64) {
	eclonDeg := macros.DMSToDD(eclipticLongitudeDegrees, eclipticLongitudeMinutes, eclipticLongitudeSeconds)
	eclatDeg := macros.DMSToDD(eclipticLatitudeDegrees, eclipticLatitudeMinutes, eclipticLatitudeSeconds)
	eclonRad := util.DegreesToRadians(eclonDeg)
	eclatRad := util.DegreesToRadians(eclatDeg)
	obliqDeg := macros.Obliq(greenwichDay, greenwichMonth, greenwichYear)
	obliqRad := util.DegreesToRadians(obliqDeg)
	sinDec := math.Sin(eclatRad)*math.Cos(obliqRad) + math.Cos(eclatRad)*math.Sin(obliqRad)*math.Sin(eclonRad)
	decRad := math.Asin(sinDec)
	decDeg := macros.Degrees(decRad)
	y := math.Sin(eclonRad)*math.Cos(obliqRad) - math.Tan(eclatRad)*math.Sin(obliqRad)
	x := math.Cos(eclonRad)
	raRad := math.Atan2(y, x)
	raDeg1 := macros.Degrees(raRad)
	raDeg2 := raDeg1 - 360.0*math.Floor(raDeg1/360.0)
	raHours := macros.DDToDH(raDeg2)

	outRaHours := macros.DHHour(raHours)
	outRaMinutes := macros.DHMin(raHours)
	outRaSeconds := macros.DHSec(raHours)
	outDecDegrees := macros.DDDeg(decDeg)
	outDecMinutes := macros.DDMin(decDeg)
	outDecSeconds := macros.DDSec(decDeg)

	return float64(outRaHours), float64(outRaMinutes), outRaSeconds, outDecDegrees, outDecMinutes, outDecSeconds
}

// EquatorialCoordinateToEclipticCoordinate converts Equatorial Coordinates to Ecliptic Coordinates
func EquatorialCoordinateToEclipticCoordinate(raHours float64, raMinutes float64, raSeconds float64, decDegrees float64, decMinutes float64, decSeconds float64, gwDay float64, gwMonth int, gwYear int) (float64, float64, float64, float64, float64, float64) {
	raDeg := macros.DHToDD(macros.HMSToDH(raHours, raMinutes, raSeconds))
	decDeg := macros.DMSToDD(decDegrees, decMinutes, decSeconds)
	raRad := util.DegreesToRadians(raDeg)
	decRad := util.DegreesToRadians(decDeg)
	obliqDeg := macros.Obliq(gwDay, gwMonth, gwYear)
	obliqRad := util.DegreesToRadians(obliqDeg)
	sinEclLat := math.Sin(decRad)*math.Cos(obliqRad) - math.Cos(decRad)*math.Sin(obliqRad)*math.Sin(raRad)
	eclLatRad := math.Asin(sinEclLat)
	eclLatDeg := macros.Degrees(eclLatRad)
	y := math.Sin(raRad)*math.Cos(obliqRad) + math.Tan(decRad)*math.Sin(obliqRad)
	x := math.Cos(raRad)
	eclLongRad := math.Atan2(y, x)
	eclLongDeg1 := macros.Degrees(eclLongRad)
	eclLongDeg2 := eclLongDeg1 - 360.0*math.Floor(eclLongDeg1/360.0)

	outEclLongDeg := macros.DDDeg(eclLongDeg2)
	outEclLongMin := macros.DDMin(eclLongDeg2)
	outEclLongSec := macros.DDSec(eclLongDeg2)
	outEclLatDeg := macros.DDDeg(eclLatDeg)
	outEclLatMin := macros.DDMin(eclLatDeg)
	outEclLatSec := macros.DDSec(eclLatDeg)

	return outEclLongDeg, outEclLongMin, outEclLongSec, outEclLatDeg, outEclLatMin, outEclLatSec
}

// EquatorialCoordinateToGalacticCoordinate converts Equatorial Coordinates to Galactic Coordinates
func EquatorialCoordinateToGalacticCoordinate(raHours float64, raMinutes float64, raSeconds float64, decDegrees float64, decMinutes float64, decSeconds float64) (float64, float64, float64, float64, float64, float64) {
	raDeg := macros.DHToDD(macros.HMSToDH(raHours, raMinutes, raSeconds))
	decDeg := macros.DMSToDD(decDegrees, decMinutes, decSeconds)
	raRad := util.DegreesToRadians(raDeg)
	decRad := util.DegreesToRadians(decDeg)
	sinB := math.Cos(decRad)*math.Cos(util.DegreesToRadians(27.4))*math.Cos(raRad-util.DegreesToRadians(192.25)) + math.Sin(decRad)*math.Sin(util.DegreesToRadians(27.4))
	bRadians := math.Asin(sinB)
	bDeg := macros.Degrees(bRadians)
	y := math.Sin(decRad) - sinB*math.Sin(util.DegreesToRadians(27.4))
	x := math.Cos(decRad) * math.Sin(raRad-util.DegreesToRadians(192.25)) * math.Cos(util.DegreesToRadians(27.4))
	longDeg1 := macros.Degrees(math.Atan2(y, x)) + 33.0
	longDeg2 := longDeg1 - 360.0*math.Floor(longDeg1/360.0)

	galLongDeg := macros.DDDeg(longDeg2)
	galLongMin := macros.DDMin(longDeg2)
	galLongSec := macros.DDSec(longDeg2)
	galLatDeg := macros.DDDeg(bDeg)
	galLatMin := macros.DDMin(bDeg)
	galLatSec := macros.DDSec(bDeg)

	return galLongDeg, galLongMin, galLongSec, galLatDeg, galLatMin, galLatSec
}

// GalacticCoordinateToEquatorialCoordinate converts Galactic Coordinates to Equatorial Coordinates
func GalacticCoordinateToEquatorialCoordinate(galLongDeg float64, galLongMin float64, galLongSec float64, galLatDeg float64, galLatMin float64, galLatSec float64) (float64, float64, float64, float64, float64, float64) {
	glongDeg := macros.DMSToDD(galLongDeg, galLongMin, galLongSec)
	glatDeg := macros.DMSToDD(galLatDeg, galLatMin, galLatSec)
	glongRad := util.DegreesToRadians(glongDeg)
	glatRad := util.DegreesToRadians(glatDeg)
	sinDec := math.Cos(glatRad)*math.Cos(util.DegreesToRadians(27.4))*math.Sin(glongRad-util.DegreesToRadians(33)) + math.Sin(glatRad)*math.Sin(util.DegreesToRadians(27.4))
	decRadians := math.Asin(sinDec)
	decDeg := macros.Degrees(decRadians)

	y := math.Cos(glatRad) * math.Cos(glongRad-util.DegreesToRadians(33))
	x := math.Sin(glatRad)*math.Cos(util.DegreesToRadians(27.4)) - math.Cos(glatRad)*math.Sin(util.DegreesToRadians(27.4))*math.Sin(glongRad-util.DegreesToRadians(33))

	raDeg1 := macros.Degrees(math.Atan2(y, x)) + 192.25
	raDeg2 := raDeg1 - 360.0*math.Floor(raDeg1/360.0)
	raHours1 := macros.DDToDH(raDeg2)

	raHours := macros.DHHour(raHours1)
	raMinutes := macros.DHMin(raHours1)
	raSeconds := macros.DHSec(raHours1)
	decDegrees := macros.DDDeg(decDeg)
	decMinutes := macros.DDMin(decDeg)
	decSeconds := macros.DDSec(decDeg)

	return float64(raHours), float64(raMinutes), raSeconds, decDegrees, decMinutes, decSeconds
}

// AngleBetweenTwoObjects calculates the angle between two celestial objects
func AngleBetweenTwoObjects(raLong1HourDeg float64, raLong1Min float64, raLong1Sec float64, decLat1Deg float64, decLat1Min float64, decLat1Sec float64, raLong2HourDeg float64, raLong2Min float64, raLong2Sec float64, decLat2Deg float64, decLat2Min float64, decLat2Sec float64, hourOrDegree string) (float64, float64, float64) {
	var raLong1Decimal float64
	if hourOrDegree == "H" {
		raLong1Decimal = macros.HMSToDH(raLong1HourDeg, raLong1Min, raLong1Sec)
	} else {
		raLong1Decimal = macros.DMSToDD(raLong1HourDeg, raLong1Min, raLong1Sec)
	}

	var raLong1Deg float64
	if hourOrDegree == "H" {
		raLong1Deg = macros.DHToDD(raLong1Decimal)
	} else {
		raLong1Deg = raLong1Decimal
	}

	raLong1Rad := util.DegreesToRadians(raLong1Deg)
	decLat1Deg1 := macros.DMSToDD(decLat1Deg, decLat1Min, decLat1Sec)
	decLat1Rad := util.DegreesToRadians(decLat1Deg1)

	var raLong2Decimal float64
	if hourOrDegree == "H" {
		raLong2Decimal = macros.HMSToDH(raLong2HourDeg, raLong2Min, raLong2Sec)
	} else {
		raLong2Decimal = macros.DMSToDD(raLong2HourDeg, raLong2Min, raLong2Sec)
	}

	var raLong2Deg float64
	if hourOrDegree == "H" {
		raLong2Deg = macros.DHToDD(raLong2Decimal)
	} else {
		raLong2Deg = raLong2Decimal
	}

	raLong2Rad := util.DegreesToRadians(raLong2Deg)
	decLat2Deg1 := macros.DMSToDD(decLat2Deg, decLat2Min, decLat2Sec)
	decLat2Rad := util.DegreesToRadians(decLat2Deg1)

	cosD := math.Sin(decLat1Rad)*math.Sin(decLat2Rad) + math.Cos(decLat1Rad)*math.Cos(decLat2Rad)*math.Cos(raLong1Rad-raLong2Rad)

	dRad := math.Acos(cosD)
	dDeg := macros.Degrees(dRad)

	angleDeg := macros.DDDeg(dDeg)
	angleMin := macros.DDMin(dDeg)
	angleSec := macros.DDSec(dDeg)

	return angleDeg, angleMin, angleSec
}

// RisingAndSetting calculates rising and setting times
//
// Arguments
//	raHours -- Right Ascension, in hours.
// 	raMinutes -- Right Ascension, in minutes.
//	raSeconds -- Right Ascension, in seconds.
//	decDeg -- Declination, in degrees.
//	decMin -- Declination, in minutes.
//	decSec -- Declination, in seconds.
//	gwDateDay -- Greenwich Date, day part.
//	gwDateMonth -- Greenwich Date, month part.
//	gwDateYear -- Greenwich Date, year part.
//	geogLongDeg -- Geographical Longitude, in degrees.
//	geogLatDeg -- Geographical Latitude, in degrees.
//	vertShiftDeg -- Vertical Shift, in degrees.
//
// Returns
//	riseSetStatus -- "Never Rises", "Circumpolar", or "OK".
//	utRiseHour -- Rise time, UT, hour part.
//	utRiseMin -- Rise time, UT, minute part.
//	utSetHour -- Set time, UT, hour part.
//	utSetMin -- Set time, UT, minute part.
//	azRise -- Azimuth angle, at rise.
//	azSet -- Azimuth angle, at set.
func RisingAndSetting(raHours float64, raMinutes float64, raSeconds float64, decDeg float64, decMin float64, decSec float64, gwDateDay float64, gwDateMonth int, gwDateYear int, geogLongDeg float64, geogLatDeg float64, vertShiftDeg float64) (string, float64, float64, float64, float64, float64, float64) {
	raHours1 := macros.HMSToDH(raHours, raMinutes, raSeconds)
	decRad := util.DegreesToRadians(macros.DMSToDD(decDeg, decMin, decSec))
	verticalDisplRadians := util.DegreesToRadians(vertShiftDeg)
	geoLatRadians := util.DegreesToRadians(geogLatDeg)
	cosH := -(math.Sin(verticalDisplRadians) + math.Sin(geoLatRadians)*math.Sin(decRad)) / (math.Cos(geoLatRadians) * math.Cos(decRad))
	hHours := macros.DDToDH(macros.Degrees(math.Acos(cosH)))
	lstRiseHours := (raHours1 - hHours) - 24.0*math.Floor((raHours1-hHours)/24.0)
	lstSetHours := (raHours1 + hHours) - 24.0*math.Floor((raHours1+hHours)/24.0)
	aDeg := macros.Degrees(math.Acos((math.Sin(decRad) + math.Sin(verticalDisplRadians)*math.Sin(geoLatRadians)) / (math.Cos(verticalDisplRadians) * math.Cos(geoLatRadians))))
	azRiseDeg := aDeg - 360.0*math.Floor(aDeg/360.0)
	azSetDeg := (360.0 - aDeg) - 360.0*math.Floor((360.0-aDeg)/360.0)
	utRiseHours1 := macros.GSTToUT(macros.LSTToGST(lstRiseHours, 0.0, 0.0, geogLongDeg), 0.0, 0.0, gwDateDay, gwDateMonth, gwDateYear)
	utSetHours1 := macros.GSTToUT(macros.LSTToGST(lstSetHours, 0.0, 0.0, geogLongDeg), 0.0, 0.0, gwDateDay, gwDateMonth, gwDateYear)
	utRiseAdjustedHours := utRiseHours1 + 0.008333
	utSetAdjustedHours := utSetHours1 + 0.008333

	riseSetStatus := "OK"
	if cosH > 1.0 {
		riseSetStatus = "never rises"
	}
	if cosH < -1.0 {
		riseSetStatus = "circumpolar"
	}

	var utRiseHour float64
	if riseSetStatus == "OK" {
		utRiseHour = float64(macros.DHHour(utRiseAdjustedHours))
	} else {
		utRiseHour = 0.0
	}

	var utRiseMin float64
	if riseSetStatus == "OK" {
		utRiseMin = float64(macros.DHMin(utRiseAdjustedHours))
	} else {
		utRiseMin = 0.0
	}

	var utSetHour float64
	if riseSetStatus == "OK" {
		utSetHour = float64(macros.DHHour(utSetAdjustedHours))
	} else {
		utSetHour = 0.0
	}

	var utSetMin float64
	if riseSetStatus == "OK" {
		utSetMin = float64(macros.DHMin(utSetAdjustedHours))
	} else {
		utSetMin = 0.0
	}

	var azRise float64
	if riseSetStatus == "OK" {
		azRise = util.RoundFloat64(azRiseDeg, 2)
	} else {
		azRise = 0.0
	}

	var azSet float64
	if riseSetStatus == "OK" {
		azSet = util.RoundFloat64(azSetDeg, 2)
	} else {
		azSet = 0.0
	}

	return riseSetStatus, utRiseHour, utRiseMin, utSetHour, utSetMin, azRise, azSet
}

// CorrectForPrecession calculates precession (corrected coordinates between two epochs)
//
// Returns
//	corrected RA hour
//	corrected RA minutes
//	corrected RA seconds
//	corrected Declination degrees
//	corrected Declination minutes
//	corrected Declination seconds
func CorrectForPrecession(raHour float64, raMinutes float64, raSeconds float64, decDeg float64, decMinutes float64, decSeconds float64, epoch1Day float64, epoch1Month int, epoch1Year int, epoch2Day float64, epoch2Month int, epoch2Year int) (float64, float64, float64, float64, float64, float64) {
	ra1Rad := util.DegreesToRadians(macros.DHToDD(macros.HMSToDH(raHour, raMinutes, raSeconds)))
	dec1Rad := util.DegreesToRadians(macros.DMSToDD(decDeg, decMinutes, decSeconds))
	tCenturies := (macros.CDToJD(epoch1Day, epoch1Month, epoch1Year) - 2415020.0) / 36525.0
	mSec := 3.07234 + (0.00186 * tCenturies)
	nArcsec := 20.0468 - (0.0085 * tCenturies)
	nYears := (macros.CDToJD(epoch2Day, epoch2Month, epoch2Year) - macros.CDToJD(epoch1Day, epoch1Month, epoch1Year)) / 365.25
	s1Hours := ((mSec + (nArcsec * math.Sin(ra1Rad) * math.Tan(dec1Rad) / 15.0)) * nYears) / 3600.0
	ra2Hours := macros.HMSToDH(raHour, raMinutes, raSeconds) + s1Hours
	s2Deg := (nArcsec * math.Cos(ra1Rad) * nYears) / 3600.0
	dec2Deg := macros.DMSToDD(decDeg, decMinutes, decSeconds) + s2Deg

	correctedRAHour := macros.DHHour(ra2Hours)
	correctedRAMinutes := macros.DHMin(ra2Hours)
	correctedRASeconds := macros.DHSec(ra2Hours)
	correctedDecDeg := macros.DDDeg(dec2Deg)
	correctedDecMinutes := macros.DDMin(dec2Deg)
	correctedDecSeconds := macros.DDSec(dec2Deg)

	return float64(correctedRAHour), float64(correctedRAMinutes), correctedRASeconds, correctedDecDeg, correctedDecMinutes, correctedDecSeconds
}

// NutationInEclipticLongitudeAndObliquity calculates nutation for two values: ecliptic longitude and obliquity,
// for a Greenwich date.
//
// Returns
//	nutation in ecliptic longitude (degrees)
//	nutation in obliquity (degrees)
func NutationInEclipticLongitudeAndObliquity(greenwichDay float64, greenwichMonth int, greenwichYear int) (float64, float64) {
	jdDays := macros.CDToJD(greenwichDay, greenwichMonth, greenwichYear)
	tCenturies := (jdDays - 2415020.0) / 36525.0
	aDeg := 100.0021358 * tCenturies
	l1Deg := 279.6967 + (0.000303 * tCenturies * tCenturies)
	lDeg1 := l1Deg + 360.0*(aDeg-math.Floor(aDeg))
	lDeg2 := lDeg1 - 360.0*math.Floor(lDeg1/360.0)
	lRad := util.DegreesToRadians(lDeg2)
	bDeg := 5.372617 * tCenturies
	nDeg1 := 259.1833 - 360.0*(bDeg-math.Floor(bDeg))
	nDeg2 := nDeg1 - 360.0*(math.Floor(nDeg1/360.0))
	nRad := util.DegreesToRadians(nDeg2)
	nutInLongArcsec := -17.2*math.Sin(nRad) - 1.3*math.Sin(2.0*lRad)
	nutInOblArcsec := 9.2*math.Cos(nRad) + 0.5*math.Cos(2.0*lRad)

	nutInLongDeg := nutInLongArcsec / 3600.0
	nutInOblDeg := nutInOblArcsec / 3600.0

	return nutInLongDeg, nutInOblDeg
}

// CorrectForAberration corrects ecliptic coordinates for the effects of aberration.
//
// Returns
//	apparent ecliptic longitude (degrees, minutes, seconds)
//	apparent ecliptic latitude (degrees, minutes, seconds)
func CorrectForAberration(utHour float64, utMinutes float64, utSeconds float64, gwDay float64, gwMonth int, gwYear int, trueEclLongDeg float64, trueEclLongMin float64, trueEclLongSec float64, trueEclLatDeg float64, trueEclLatMin float64, trueEclLatSec float64) (float64, float64, float64, float64, float64, float64) {
	trueLongDeg := macros.DMSToDD(trueEclLongDeg, trueEclLongMin, trueEclLongSec)
	trueLatDeg := macros.DMSToDD(trueEclLatDeg, trueEclLatMin, trueEclLatSec)
	sunTrueLongDeg := macros.SunEclipticLongitude(utHour, utMinutes, utSeconds, 0, 0, gwDay, gwMonth, gwYear)
	dlongArcsec := -20.5 * math.Cos(util.DegreesToRadians(sunTrueLongDeg-trueLongDeg)) / math.Cos(util.DegreesToRadians(trueLatDeg))
	dlatArcsec := -20.5 * math.Sin(util.DegreesToRadians(sunTrueLongDeg-trueLongDeg)) * math.Sin(util.DegreesToRadians(trueLatDeg))
	apparentLongDeg := trueLongDeg + (dlongArcsec / 3600.0)
	apparentLatDeg := trueLatDeg + (dlatArcsec / 3600.0)

	apparentEclLongDeg := macros.DDDeg(apparentLongDeg)
	apparentEclLongMin := macros.DDMin(apparentLongDeg)
	apparentEclLongSec := macros.DDSec(apparentLongDeg)
	apparentEclLatDeg := macros.DDDeg(apparentLatDeg)
	apparentEclLatMin := macros.DDMin(apparentLatDeg)
	apparentEclLatSec := macros.DDSec(apparentLatDeg)

	return apparentEclLongDeg, apparentEclLongMin, apparentEclLongSec, apparentEclLatDeg, apparentEclLatMin, apparentEclLatSec
}

// AtmosphericRefraction calculates corrected RA/Dec, accounting for atmospheric refraction.
//
// NOTE: Valid values for coordinate_type are "TRUE" and "APPARENT".
//
// Returns
//	corrected RA hours,minutes,seconds
//	corrected Declination degrees,minutes,seconds
func AtmosphericRefraction(trueRAHour float64, trueRAMin float64, trueRASec float64, trueDecDeg float64, trueDecMin float64, trueDecSec float64, coordinateType string, geogLongDeg float64, geogLatDeg float64, daylightSavingHours int, timezoneHours int, lcdDay float64, lcdMonth int, lcdYear int, lctHour float64, lctMin float64, lctSec float64, atmosphericPressureMbar float64, atmosphericTemperatureCelsius float64) (float64, float64, float64, float64, float64, float64) {
	haHour := macros.RAToHA(trueRAHour, trueRAMin, trueRASec, lctHour, lctMin, lctSec, daylightSavingHours, timezoneHours, lcdDay, lcdMonth, lcdYear, geogLongDeg)
	azimuthDeg := macros.EqToAz(haHour, 0.0, 0.0, trueDecDeg, trueDecMin, trueDecSec, geogLatDeg)
	altitudeDeg := macros.EqToAlt(haHour, 0.0, 0.0, trueDecDeg, trueDecMin, trueDecSec, geogLatDeg)
	correctedAltitudeDeg := macros.Refraction(altitudeDeg, coordinateType, atmosphericPressureMbar, atmosphericTemperatureCelsius)

	correctedHaHour := macros.HorToHA(azimuthDeg, 0.0, 0.0, correctedAltitudeDeg, 0.0, 0.0, geogLatDeg)
	correctedRAHour1 := macros.HAToRA(correctedHaHour, 0.0, 0.0, lctHour, lctMin, lctSec, daylightSavingHours, timezoneHours, lcdDay, lcdMonth, lcdYear, geogLongDeg)
	correctedDecDeg1 := macros.HorToDec(azimuthDeg, 0.0, 0.0, correctedAltitudeDeg, 0.0, 0.0, geogLatDeg)

	correctedRAHour := macros.DHHour(correctedRAHour1)
	correctedRAMin := macros.DHMin(correctedRAHour1)
	correctedRASec := macros.DHSec(correctedRAHour1)
	correctedDecDeg := macros.DDDeg(correctedDecDeg1)
	correctedDecMin := macros.DDMin(correctedDecDeg1)
	correctedDecSec := macros.DDSec(correctedDecDeg1)

	return float64(correctedRAHour), float64(correctedRAMin), correctedRASec, correctedDecDeg, correctedDecMin, correctedDecSec
}
