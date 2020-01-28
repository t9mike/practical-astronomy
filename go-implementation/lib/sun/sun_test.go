package sun

import (
	"fmt"
	"testing"
)

func TestApproximatePositionOfSun(t *testing.T) {
	type args struct {
		lctHours         float64
		lctMinutes       float64
		lctSeconds       float64
		localDay         float64
		localMonth       int
		localYear        int
		isDaylightSaving bool
		zoneCorrection   int
	}
	tests := []struct {
		name              string
		args              args
		wantSunRAHour     float64
		wantSunRAMinutes  float64
		wantSunRASeconds  float64
		wantSunDecDeg     float64
		wantSunDecMinutes float64
		wantSunDecSeconds float64
	}{
		{name: "ApproximatePositionOfSun", args: args{lctHours: 0, lctMinutes: 0, lctSeconds: 0, localDay: 27, localMonth: 7, localYear: 2003, isDaylightSaving: false, zoneCorrection: 0}, wantSunRAHour: 8, wantSunRAMinutes: 23, wantSunRASeconds: 33.73, wantSunDecDeg: 19, wantSunDecMinutes: 21, wantSunDecSeconds: 14.33},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sunRAHour, sunRAMin, sunRASec, sunDecDeg, sunDecMin, sunDecSec := ApproximatePositionOfSun(tt.args.lctHours, tt.args.lctMinutes, tt.args.lctSeconds, tt.args.localDay, tt.args.localMonth, tt.args.localYear, tt.args.isDaylightSaving, tt.args.zoneCorrection)

			if sunRAHour != tt.wantSunRAHour || sunRAMin != tt.wantSunRAMinutes || sunRASec != tt.wantSunRASeconds || sunDecDeg != tt.wantSunDecDeg || sunDecMin != tt.wantSunDecMinutes || sunDecSec != tt.wantSunDecSeconds {
				t.Errorf("ApproximatePositionOfSun() got = [RA] %v hours %v minutes %v seconds [Declination] %v degrees %v minutes %v seconds, want [RA] %v hours %v minutes %v seconds [Declination] %v degrees %v minutes %v seconds", sunRAHour, sunRAMin, sunRASec, sunDecDeg, sunDecMin, sunDecSec, tt.wantSunRAHour, tt.wantSunRAMinutes, tt.wantSunRASeconds, tt.wantSunDecDeg, tt.wantSunDecMinutes, tt.wantSunDecSeconds)
			} else {
				fmt.Printf("Approximate position of sun: [Local Civil Time] %v:%v:%v [Local Date] %v/%v/%v [DST?] %v [Zone Correction] %v = [RA] %v hours %v minutes %v seconds [Declination] %v degrees %v minutes %v seconds\n", tt.args.lctHours, tt.args.lctMinutes, tt.args.lctSeconds, tt.args.localMonth, tt.args.localDay, tt.args.localYear, tt.args.isDaylightSaving, tt.args.zoneCorrection, sunRAHour, sunRAMin, sunRASec, sunDecDeg, sunDecMin, sunDecSec)
			}
		})
	}
}

func TestPrecisePositionOfSun(t *testing.T) {
	type args struct {
		lctHours         float64
		lctMinutes       float64
		lctSeconds       float64
		localDay         float64
		localMonth       int
		localYear        int
		isDaylightSaving bool
		zoneCorrection   int
	}
	tests := []struct {
		name              string
		args              args
		wantSunRAHour     float64
		wantSunRAMinutes  float64
		wantSunRASeconds  float64
		wantSunDecDeg     float64
		wantSunDecMinutes float64
		wantSunDecSeconds float64
	}{
		{name: "PrecisePositionOfSun", args: args{lctHours: 0, lctMinutes: 0, lctSeconds: 0, localDay: 27, localMonth: 7, localYear: 1988, isDaylightSaving: false, zoneCorrection: 0}, wantSunRAHour: 8, wantSunRAMinutes: 26, wantSunRASeconds: 3.83, wantSunDecDeg: 19, wantSunDecMinutes: 12, wantSunDecSeconds: 49.72},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sunRAHour, sunRAMin, sunRASec, sunDecDeg, sunDecMin, sunDecSec := PrecisePositionOfSun(tt.args.lctHours, tt.args.lctMinutes, tt.args.lctSeconds, tt.args.localDay, tt.args.localMonth, tt.args.localYear, tt.args.isDaylightSaving, tt.args.zoneCorrection)

			if sunRAHour != tt.wantSunRAHour || sunRAMin != tt.wantSunRAMinutes || sunRASec != tt.wantSunRASeconds || sunDecDeg != tt.wantSunDecDeg || sunDecMin != tt.wantSunDecMinutes || sunDecSec != tt.wantSunDecSeconds {
				t.Errorf("PrecisePositionOfSun() got = [RA] %v hours %v minutes %v seconds [Declination] %v degrees %v minutes %v seconds, want [RA] %v hours %v minutes %v seconds [Declination] %v degrees %v minutes %v seconds", sunRAHour, sunRAMin, sunRASec, sunDecDeg, sunDecMin, sunDecSec, tt.wantSunRAHour, tt.wantSunRAMinutes, tt.wantSunRASeconds, tt.wantSunDecDeg, tt.wantSunDecMinutes, tt.wantSunDecSeconds)
			} else {
				fmt.Printf("Precise position of sun: [Local Civil Time] %v:%v:%v [Local Date] %v/%v/%v [DST?] %v [Zone Correction] %v = [RA] %v hours %v minutes %v seconds [Declination] %v degrees %v minutes %v seconds\n", tt.args.lctHours, tt.args.lctMinutes, tt.args.lctSeconds, tt.args.localMonth, tt.args.localDay, tt.args.localYear, tt.args.isDaylightSaving, tt.args.zoneCorrection, sunRAHour, sunRAMin, sunRASec, sunDecDeg, sunDecMin, sunDecSec)
			}
		})
	}
}
