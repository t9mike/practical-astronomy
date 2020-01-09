package coordinates

import "testing"
import "fmt"
import "../util"

func TestAngleToDecimalDegrees(t *testing.T) {
	degrees := 182.0
	minutes := 31.0
	seconds := 27.0

	expectedDecimalDegrees := 182.524167

	decimalDegrees := util.RoundFloat64(AngleToDecimalDegrees(degrees, minutes, seconds), 6)

	if decimalDegrees == expectedDecimalDegrees {
		fmt.Printf("Angle to decimal degrees: [Angle] %.0f degrees %.0f minutes %.0f seconds = [Decimal Degrees] %f\n", degrees, minutes, seconds, decimalDegrees)
	} else {
		t.Errorf("Expected %f, actual %f\n", expectedDecimalDegrees, decimalDegrees)
	}

}

func TestDecimalDegreesToAngle(t *testing.T) {
	decimalDegrees := 182.524167

	expectedDegrees := 182.0
	expectedMinutes := 31.0
	expectedSeconds := 27.0

	degrees, minutes, seconds := DecimalDegreesToAngle(decimalDegrees)

	if degrees == expectedDegrees && minutes == expectedMinutes && seconds == expectedSeconds {
		fmt.Printf("Decimal degrees to angle:  [Decimal Degrees] %f = [Angle] %.0f degrees %.0f minutes %.0f seconds\n", decimalDegrees, degrees, minutes, seconds)
	} else {
		t.Errorf("Expected %.0f degrees %.0f minutes %.0f seconds, actual %.0f degrees %.0f minutes %.0f seconds\n", expectedDegrees, expectedMinutes, expectedSeconds, degrees, minutes, seconds)
	}
}
