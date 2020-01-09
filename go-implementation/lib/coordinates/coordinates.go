package coordinates

import (
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
