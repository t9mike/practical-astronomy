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
