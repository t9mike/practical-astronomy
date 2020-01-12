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
