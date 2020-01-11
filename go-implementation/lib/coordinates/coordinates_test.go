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
				fmt.Printf("Angle to decimal degrees: [Angle] %v degrees %v minutes %v seconds = [Decimal Degrees] %v\n", tt.args.degrees, tt.args.minutes, tt.args.seconds, decimalDegrees)
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
				t.Errorf("DecimalDegreesToAngle() got = %v degrees %v minutes %v seconds, want %v degrees %v minutes %v seconds", degrees, minutes, seconds, tt.wantDegrees, tt.wantMinutes, tt.wantSeconds)
			} else {
				fmt.Printf("Decimal degrees to angle: [Decimal Degrees] %v = [Angle] %v degrees %v minutes %v seconds\n", tt.args.decimalDegrees, degrees, minutes, seconds)
			}
		})
	}
}

func TestRightAscensionToHourAngle(t *testing.T) {
	type args struct {
		raHours               float64
		raMinutes             float64
		raSeconds             float64
		lctHours              float64
		lctMinutes            float64
		lctSeconds            float64
		isDaylightSaving      bool
		zoneCorrection        int
		localDay              float64
		localMonth            int
		localYear             int
		geographicalLongitude float64
	}
	tests := []struct {
		name                 string
		args                 args
		wantHourAngleHours   float64
		wantHourAngleMinutes float64
		wantHourAngleSeconds float64
	}{
		{name: "RightAscensionToHourAngle", args: args{raHours: 18, raMinutes: 32, raSeconds: 21, lctHours: 14, lctMinutes: 36, lctSeconds: 51.67, isDaylightSaving: false, zoneCorrection: -4, localDay: 22, localMonth: 4, localYear: 1980, geographicalLongitude: -64.0}, wantHourAngleHours: 9, wantHourAngleMinutes: 52, wantHourAngleSeconds: 23.66},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hourAngleHours, hourAngleMinutes, hourAngleSeconds := RightAscensionToHourAngle(tt.args.raHours, tt.args.raMinutes, tt.args.raSeconds, tt.args.lctHours, tt.args.lctMinutes, tt.args.lctSeconds, tt.args.isDaylightSaving, tt.args.zoneCorrection, tt.args.localDay, tt.args.localMonth, tt.args.localYear, tt.args.geographicalLongitude)

			if hourAngleHours != tt.wantHourAngleHours || hourAngleMinutes != tt.wantHourAngleMinutes || hourAngleSeconds != tt.wantHourAngleSeconds {
				t.Errorf("RightAscensionToHourAngle() got = %v:%v:%v, want %v:%v:%v", hourAngleHours, hourAngleMinutes, hourAngleSeconds, tt.wantHourAngleHours, tt.wantHourAngleMinutes, tt.wantHourAngleSeconds)
			} else {
				fmt.Printf("Right Ascension to Hour Angle: [RA] %v:%v:%v [LCT] %v:%v:%v [DST?] %t [ZC] %d [Local Day] %d/%.0f/%d [Geog Long] %v = [Hour Angle] %v:%v:%v\n", tt.args.raHours, tt.args.raMinutes, tt.args.raSeconds, tt.args.lctHours, tt.args.lctMinutes, tt.args.lctSeconds, tt.args.isDaylightSaving, tt.args.zoneCorrection, tt.args.localMonth, tt.args.localDay, tt.args.localYear, tt.args.geographicalLongitude, hourAngleHours, hourAngleMinutes, hourAngleSeconds)
			}
		})
	}
}

func TestHourAngleToRightAscension(t *testing.T) {
	type args struct {
		hourAngleHours        float64
		hourAngleMinutes      float64
		hourAngleSeconds      float64
		lctHours              float64
		lctMinutes            float64
		lctSeconds            float64
		isDaylightSaving      bool
		zoneCorrection        int
		localDay              float64
		localMonth            int
		localYear             int
		geographicalLongitude float64
	}
	tests := []struct {
		name                      string
		args                      args
		wantRightAscensionHours   float64
		wantRightAscensionMinutes float64
		wantRightAscensionSeconds float64
	}{
		{name: "HourAngleToRightAscension", args: args{hourAngleHours: 9, hourAngleMinutes: 52, hourAngleSeconds: 23.66, lctHours: 14, lctMinutes: 36, lctSeconds: 51.67, isDaylightSaving: false, zoneCorrection: -4, localDay: 22, localMonth: 4, localYear: 1980, geographicalLongitude: -64}, wantRightAscensionHours: 18, wantRightAscensionMinutes: 32, wantRightAscensionSeconds: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rightAscensionHours, rightAscensionMinutes, rightAscensionSeconds := HourAngleToRightAscension(tt.args.hourAngleHours, tt.args.hourAngleMinutes, tt.args.hourAngleSeconds, tt.args.lctHours, tt.args.lctMinutes, tt.args.lctSeconds, tt.args.isDaylightSaving, tt.args.zoneCorrection, tt.args.localDay, tt.args.localMonth, tt.args.localYear, tt.args.geographicalLongitude)

			if rightAscensionHours != tt.wantRightAscensionHours || rightAscensionMinutes != tt.wantRightAscensionMinutes || rightAscensionSeconds != tt.wantRightAscensionSeconds {
				t.Errorf("HourAngleToRightAscension() got = %v:%v:%v, want %v:%v:%v", rightAscensionHours, rightAscensionMinutes, rightAscensionSeconds, tt.wantRightAscensionHours, tt.wantRightAscensionMinutes, tt.wantRightAscensionSeconds)
			} else {
				fmt.Printf("Hour Angle to Right Ascension: [Hour Angle] %v:%v:%v [LCT] %v:%v:%v [DST?] %v [ZC] %v [Local Day] %v/%v/%v [Geog Longitude] %v = [Right Ascension] %v:%v:%v\n", tt.args.hourAngleHours, tt.args.hourAngleMinutes, tt.args.hourAngleSeconds, tt.args.lctHours, tt.args.lctMinutes, tt.args.lctSeconds, tt.args.isDaylightSaving, tt.args.zoneCorrection, tt.args.localMonth, tt.args.localDay, tt.args.localYear, tt.args.geographicalLongitude, rightAscensionHours, rightAscensionMinutes, rightAscensionSeconds)
			}
		})
	}
}
