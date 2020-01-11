package coordinates

import (
	"fmt"
	"testing"

	"../util"
)

func TestAngleToDecimalDegrees(t *testing.T) {
	type args struct {
		degrees float64
		minutes float64
		seconds float64
	}
	tests := []struct {
		name               string
		args               args
		wantDecimalDegrees float64
	}{
		{name: "AngleToDecimalDegrees", args: args{182.0, 31.0, 27.0}, wantDecimalDegrees: 182.524167},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if decimalDegrees := util.RoundFloat64(AngleToDecimalDegrees(tt.args.degrees, tt.args.minutes, tt.args.seconds), 6); decimalDegrees != tt.wantDecimalDegrees {
				t.Errorf("AngleToDecimalDegrees() = %v, want %v", decimalDegrees, tt.wantDecimalDegrees)
			} else {
				fmt.Printf("Angle to decimal degrees: [Angle] %.0f degrees %.0f minutes %.0f seconds = [Decimal Degrees] %f\n", tt.args.degrees, tt.args.minutes, tt.args.seconds, decimalDegrees)
			}
		})
	}
}

func TestDecimalDegreesToAngle(t *testing.T) {
	type args struct {
		decimalDegrees float64
	}
	tests := []struct {
		name        string
		args        args
		wantDegrees float64
		wantMinutes float64
		wantSeconds float64
	}{
		{name: "DecimalDegreesToAngle", args: args{decimalDegrees: 182.524167}, wantDegrees: 182.0, wantMinutes: 31.0, wantSeconds: 27.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			degrees, minutes, seconds := DecimalDegreesToAngle(tt.args.decimalDegrees)
			if degrees != tt.wantDegrees || minutes != tt.wantMinutes || seconds != tt.wantSeconds {
				t.Errorf("DecimalDegreesToAngle() got = %.0f degrees %.0f minutes %.0f seconds, want %.0f degrees %.0f minutes %.0f seconds", degrees, minutes, seconds, tt.wantDegrees, tt.wantMinutes, tt.wantSeconds)
			} else {
				fmt.Printf("Decimal degrees to angle: [Decimal Degrees] %f = [Angle] %.0f degrees %.0f minutes %.0f seconds\n", tt.args.decimalDegrees, degrees, minutes, seconds)
			}
		})
	}
}
