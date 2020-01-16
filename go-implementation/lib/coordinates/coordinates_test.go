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

func TestEquatorialCoordinatesToHorizonCoordinates(t *testing.T) {
	type args struct {
		hourAngleHours       float64
		hourAngleMinutes     float64
		hourAngleSeconds     float64
		declinationDegrees   float64
		declinationMinutes   float64
		declinationSeconds   float64
		geographicalLatitude float64
	}
	tests := []struct {
		name                string
		args                args
		wantAzimuthDegrees  float64
		wantAzimuthMinutes  float64
		wantAzimuthSeconds  float64
		wantAltitudeDegrees float64
		wantAltitudeMinutes float64
		wantAltitudeSeconds float64
	}{
		{name: "EquatorialCoordinatesToHorizonCoordinates", args: args{hourAngleHours: 5, hourAngleMinutes: 51, hourAngleSeconds: 44, declinationDegrees: 23, declinationMinutes: 13, declinationSeconds: 10, geographicalLatitude: 52}, wantAzimuthDegrees: 283, wantAzimuthMinutes: 16, wantAzimuthSeconds: 15.7, wantAltitudeDegrees: 19, wantAltitudeMinutes: 20, wantAltitudeSeconds: 3.64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			azimuthDegrees, azimuthMinutes, azimuthSeconds, altitudeDegrees, altitudeMinutes, altitudeSeconds := EquatorialCoordinatesToHorizonCoordinates(tt.args.hourAngleHours, tt.args.hourAngleMinutes, tt.args.hourAngleSeconds, tt.args.declinationDegrees, tt.args.declinationMinutes, tt.args.declinationSeconds, tt.args.geographicalLatitude)

			if azimuthDegrees != tt.wantAzimuthDegrees || azimuthMinutes != tt.wantAzimuthMinutes || azimuthSeconds != tt.wantAzimuthSeconds || altitudeDegrees != tt.wantAltitudeDegrees || altitudeMinutes != tt.wantAltitudeMinutes || altitudeSeconds != tt.wantAltitudeSeconds {
				t.Errorf("EquatorialCoordinatesToHorizonCoordinates() got = [Az] %v degrees %v minutes %v seconds [Alt] %v degrees %v minutes %v seconds, want [Az] %v degrees %v minutes %v seconds [Alt] %v degrees %v minutes %v seconds\n", azimuthDegrees, azimuthMinutes, azimuthSeconds, altitudeDegrees, altitudeMinutes, altitudeSeconds, tt.wantAzimuthDegrees, tt.wantAzimuthMinutes, tt.wantAzimuthSeconds, tt.wantAltitudeDegrees, tt.wantAltitudeMinutes, tt.wantAltitudeSeconds)
			} else {
				fmt.Printf("Equatorial Coordinates to Horizon Coordinates: [Hour Angle] %v:%v:%v [Declination] %v degrees %v minutes %v seconds [Geog Lat] %v = [Azimuth] %v degrees %v minutes %v seconds [Altitude] %v degrees %v minutes %v seconds\n", tt.args.hourAngleHours, tt.args.hourAngleMinutes, tt.args.hourAngleSeconds, tt.args.declinationDegrees, tt.args.declinationMinutes, tt.args.declinationSeconds, tt.args.geographicalLatitude, azimuthDegrees, azimuthMinutes, azimuthSeconds, altitudeDegrees, altitudeMinutes, altitudeSeconds)
			}
		})
	}
}

func TestHorizonCoordinatesToEquatorialCoordinates(t *testing.T) {
	type args struct {
		azimuthDegrees       float64
		azimuthMinutes       float64
		azimuthSeconds       float64
		altitudeDegrees      float64
		altitudeMinutes      float64
		altitudeSeconds      float64
		geographicalLatitude float64
	}
	tests := []struct {
		name                   string
		args                   args
		wantHourAngleHours     float64
		wantHourAngleMinutes   float64
		wantHourAngleSeconds   float64
		wantDeclinationDegrees float64
		wantDeclinationMinutes float64
		wantDeclinationSeconds float64
	}{
		{name: "HorizonCoordinatesToEquatorialCoordinates", args: args{azimuthDegrees: 283, azimuthMinutes: 16, azimuthSeconds: 15.7, altitudeDegrees: 19, altitudeMinutes: 20, altitudeSeconds: 3.64, geographicalLatitude: 52}, wantHourAngleHours: 5, wantHourAngleMinutes: 51, wantHourAngleSeconds: 44, wantDeclinationDegrees: 23, wantDeclinationMinutes: 13, wantDeclinationSeconds: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hourAngleHours, hourAngleMinutes, hourAngleSeconds, declinationDegrees, declinationMinutes, declinationSeconds := HorizonCoordinatesToEquatorialCoordinates(tt.args.azimuthDegrees, tt.args.azimuthMinutes, tt.args.azimuthSeconds, tt.args.altitudeDegrees, tt.args.altitudeMinutes, tt.args.altitudeSeconds, tt.args.geographicalLatitude)

			if hourAngleHours != tt.wantHourAngleHours || hourAngleMinutes != tt.wantHourAngleMinutes || hourAngleSeconds != tt.wantHourAngleSeconds || declinationDegrees != tt.wantDeclinationDegrees || declinationMinutes != tt.wantDeclinationMinutes || declinationSeconds != tt.wantDeclinationSeconds {
				t.Errorf("HorizonCoordinatesToEquatorialCoordinates() got = [Hour Angle] %v:%v:%v [Declination] %v degrees %v minutes %v seconds, want [Hour Angle] %v:%v:%v [Declination] %v degrees %v minutes %v seconds", hourAngleHours, hourAngleMinutes, hourAngleSeconds, declinationDegrees, declinationMinutes, declinationSeconds, tt.wantHourAngleHours, tt.wantHourAngleMinutes, tt.wantHourAngleSeconds, tt.wantDeclinationDegrees, tt.wantDeclinationMinutes, tt.wantDeclinationSeconds)
			} else {
				fmt.Printf("Horizon Coordinates to Equatorial Coordinates: [Azimuth] %v degrees %v minutes %v seconds [Altitude] %v degrees %v minutes %v seconds [Geog Latitude] %v = [Hour Angle] %v:%v:%v [Declination] %v degrees %v minutes %v seconds\n", tt.args.azimuthDegrees, tt.args.azimuthMinutes, tt.args.azimuthSeconds, tt.args.altitudeDegrees, tt.args.altitudeMinutes, tt.args.altitudeSeconds, tt.args.geographicalLatitude, hourAngleHours, hourAngleMinutes, hourAngleSeconds, declinationDegrees, declinationMinutes, declinationSeconds)
			}
		})
	}
}

func TestMeanObliquityOfTheEcliptic(t *testing.T) {
	type args struct {
		greenwichDay   float64
		greenwichMonth int
		greenwichYear  int
	}
	tests := []struct {
		name              string
		args              args
		wantMeanObliquity float64
	}{
		{name: "MeanObliquityOfTheEcliptic", args: args{greenwichDay: 6, greenwichMonth: 7, greenwichYear: 2009}, wantMeanObliquity: 23.43805531},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			meanObliquity := util.RoundFloat64(MeanObliquityOfTheEcliptic(tt.args.greenwichDay, tt.args.greenwichMonth, tt.args.greenwichYear), 8)

			if meanObliquity != tt.wantMeanObliquity {
				t.Errorf("MeanObliquityOfTheEcliptic() = %v, want %v", meanObliquity, tt.wantMeanObliquity)
			} else {
				fmt.Printf("Mean obliquity of the ecliptic: [Greenwich Date] %v/%v/%v = [Mean Obliquity] %v\n", tt.args.greenwichMonth, tt.args.greenwichDay, tt.args.greenwichYear, meanObliquity)
			}
		})
	}
}

func TestEclipticCoordinateToEquatorialCoordinate(t *testing.T) {
	type args struct {
		eclipticLongitudeDegrees float64
		eclipticLongitudeMinutes float64
		eclipticLongitudeSeconds float64
		eclipticLatitudeDegrees  float64
		eclipticLatitudeMinutes  float64
		eclipticLatitudeSeconds  float64
		greenwichDay             float64
		greenwichMonth           int
		greenwichYear            int
	}
	tests := []struct {
		name           string
		args           args
		wantRAHours    float64
		wantRAMinutes  float64
		wantRASeconds  float64
		wantDecDegrees float64
		wantDecMinutes float64
		wantDecSeconds float64
	}{
		{name: "EclipticCoordinateToEquatorialCoordinate", args: args{eclipticLongitudeDegrees: 139, eclipticLongitudeMinutes: 41, eclipticLongitudeSeconds: 10, eclipticLatitudeDegrees: 4, eclipticLatitudeMinutes: 52, eclipticLatitudeSeconds: 31, greenwichDay: 6, greenwichMonth: 7, greenwichYear: 2009}, wantRAHours: 9, wantRAMinutes: 34, wantRASeconds: 53.4, wantDecDegrees: 19, wantDecMinutes: 32, wantDecSeconds: 8.52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raHours, raMinutes, raSeconds, decDegrees, decMinutes, decSeconds := EclipticCoordinateToEquatorialCoordinate(tt.args.eclipticLongitudeDegrees, tt.args.eclipticLongitudeMinutes, tt.args.eclipticLongitudeSeconds, tt.args.eclipticLatitudeDegrees, tt.args.eclipticLatitudeMinutes, tt.args.eclipticLatitudeSeconds, tt.args.greenwichDay, tt.args.greenwichMonth, tt.args.greenwichYear)

			if raHours != tt.wantRAHours || raMinutes != tt.wantRAMinutes || raSeconds != tt.wantRASeconds || decDegrees != tt.wantDecDegrees || decMinutes != tt.wantDecMinutes || decSeconds != tt.wantDecSeconds {
				t.Errorf("EclipticCoordinateToEquatorialCoordinate() got = [RA] %v hours %v minutes %v seconds [Dec] %v degrees %v minutes %v seconds, want [RA] %v hours %v minutes %v seconds [Dec] %v degrees %v minutes %v seconds", raHours, raMinutes, raSeconds, decDegrees, decMinutes, decSeconds, tt.wantRAHours, tt.wantRAMinutes, tt.wantRASeconds, tt.wantDecDegrees, tt.wantDecMinutes, tt.wantDecSeconds)
			} else {
				fmt.Printf("Ecliptic coordinate to equatorial coordinate: [Ecliptic] [Longitude] %v degrees %v minutes %v seconds [Latitude] %v degrees %v minutes %v seconds [Greenwich Date] %v/%v/%v = [Right Ascension] %v hours %v minutes %v seconds [Declination] %v degrees %v minutes %v seconds\n", tt.args.eclipticLongitudeDegrees, tt.args.eclipticLongitudeMinutes, tt.args.eclipticLongitudeSeconds, tt.args.eclipticLatitudeDegrees, tt.args.eclipticLatitudeMinutes, tt.args.eclipticLatitudeSeconds, tt.args.greenwichMonth, tt.args.greenwichDay, tt.args.greenwichYear, raHours, raMinutes, raSeconds, decDegrees, decMinutes, decSeconds)
			}
		})
	}
}

func TestEquatorialCoordinateToEclipticCoordinate(t *testing.T) {
	type args struct {
		raHours    float64
		raMinutes  float64
		raSeconds  float64
		decDegrees float64
		decMinutes float64
		decSeconds float64
		gwDay      float64
		gwMonth    int
		gwYear     int
	}
	tests := []struct {
		name                         string
		args                         args
		wantEclipticLongitudeDegrees float64
		wantEclipticLongitudeMinutes float64
		wantEclipticLongitudeSeconds float64
		wantEclipticLatitudeDegrees  float64
		wantEclipticLatitudeMinutes  float64
		wantEclipticLatitudeSeconds  float64
	}{
		{name: "EquatorialCoordinateToEclipticCoordinate", args: args{raHours: 9, raMinutes: 34, raSeconds: 53.4, decDegrees: 19, decMinutes: 32, decSeconds: 8.52, gwDay: 6, gwMonth: 7, gwYear: 2009}, wantEclipticLongitudeDegrees: 139, wantEclipticLongitudeMinutes: 41, wantEclipticLongitudeSeconds: 9.97, wantEclipticLatitudeDegrees: 4, wantEclipticLatitudeMinutes: 52, wantEclipticLatitudeSeconds: 30.99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eclipticLongitudeDegrees, eclipticLongitudeMinutes, eclipticLongitudeSeconds, eclipticLatitudeDegrees, eclipticLatitudeMinutes, eclipticLatitudeSeconds := EquatorialCoordinateToEclipticCoordinate(tt.args.raHours, tt.args.raMinutes, tt.args.raSeconds, tt.args.decDegrees, tt.args.decMinutes, tt.args.decSeconds, tt.args.gwDay, tt.args.gwMonth, tt.args.gwYear)

			if eclipticLongitudeDegrees != tt.wantEclipticLongitudeDegrees || eclipticLongitudeMinutes != tt.wantEclipticLongitudeMinutes || eclipticLongitudeSeconds != tt.wantEclipticLongitudeSeconds || eclipticLatitudeDegrees != tt.wantEclipticLatitudeDegrees || eclipticLatitudeMinutes != tt.wantEclipticLatitudeMinutes || eclipticLatitudeSeconds != tt.wantEclipticLatitudeSeconds {
				t.Errorf("EquatorialCoordinateToEclipticCoordinate() got = [Ecliptic] [Longitude] %v degrees %v minutes %v seconds [Latitude] %v degrees %v minutes %v seconds, want [Ecliptic] [Longitude] %v degrees %v minutes %v seconds [Latitude] %v degrees %v minutes %v seconds\\", eclipticLongitudeDegrees, eclipticLongitudeMinutes, eclipticLongitudeSeconds, eclipticLatitudeDegrees, eclipticLatitudeMinutes, eclipticLatitudeSeconds, tt.wantEclipticLongitudeDegrees, tt.wantEclipticLongitudeMinutes, tt.wantEclipticLongitudeSeconds, tt.wantEclipticLatitudeDegrees, tt.wantEclipticLatitudeMinutes, tt.wantEclipticLatitudeSeconds)
			} else {
				fmt.Printf("Equatorial coordinate to ecliptic coordinate: [Right Ascension] %v hours %v minutes %v seconds [Declination] %v degrees %v minutes %v seconds [Greenwich Date] %v/%v/%v = [Ecliptic] [Longitude] %v degrees %v minutes %v seconds [Latitude] %v degrees %v minutes %v seconds\n", tt.args.raHours, tt.args.raMinutes, tt.args.raSeconds, tt.args.decDegrees, tt.args.decMinutes, tt.args.decSeconds, tt.args.gwMonth, tt.args.gwDay, tt.args.gwYear, eclipticLongitudeDegrees, eclipticLongitudeMinutes, eclipticLongitudeSeconds, eclipticLatitudeDegrees, eclipticLatitudeMinutes, eclipticLatitudeSeconds)
			}
		})
	}
}

func TestEquatorialCoordinateToGalacticCoordinate(t *testing.T) {
	type args struct {
		raHours    float64
		raMinutes  float64
		raSeconds  float64
		decDegrees float64
		decMinutes float64
		decSeconds float64
	}
	tests := []struct {
		name                         string
		args                         args
		wantGalacticLongitudeDegrees float64
		wantGalacticLongitudeMinutes float64
		wantGalacticLongitudeSeconds float64
		wantGalacticLatitudeDegrees  float64
		wantGalacticLatitudeMinutes  float64
		wantGalacticLatitudeSeconds  float64
	}{
		{name: "EquatorialCoordinateToGalacticCoordinate", args: args{raHours: 10, raMinutes: 21, raSeconds: 0, decDegrees: 10, decMinutes: 3, decSeconds: 11}, wantGalacticLongitudeDegrees: 232, wantGalacticLongitudeMinutes: 14, wantGalacticLongitudeSeconds: 52.38, wantGalacticLatitudeDegrees: 51, wantGalacticLatitudeMinutes: 7, wantGalacticLatitudeSeconds: 20.16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			galLongDeg, galLongMin, galLongSec, galLatDeg, galLatMin, galLatSec := EquatorialCoordinateToGalacticCoordinate(tt.args.raHours, tt.args.raMinutes, tt.args.raSeconds, tt.args.decDegrees, tt.args.decMinutes, tt.args.decSeconds)

			if galLongDeg != tt.wantGalacticLongitudeDegrees || galLongMin != tt.wantGalacticLongitudeMinutes || galLongSec != tt.wantGalacticLongitudeSeconds || galLatDeg != tt.wantGalacticLatitudeDegrees || galLatMin != tt.wantGalacticLatitudeMinutes || galLatSec != tt.wantGalacticLatitudeSeconds {
				t.Errorf("EquatorialCoordinateToGalacticCoordinate() got = [Galactic] [Longitude] %v degrees %v minutes %v seconds [Latitude] %v degrees %v minutes %v seconds, want [Galactic] [Longitude] %v degrees %v minutes %v seconds [Latitude] %v degrees %v minutes %v seconds", galLongDeg, galLongMin, galLongSec, galLatDeg, galLatMin, galLatSec, tt.wantGalacticLongitudeDegrees, tt.wantGalacticLongitudeMinutes, tt.wantGalacticLongitudeSeconds, tt.wantGalacticLatitudeDegrees, tt.wantGalacticLatitudeMinutes, tt.wantGalacticLatitudeSeconds)
			} else {
				fmt.Printf("Equatorial coordinate to galactic coordinate: [RA] %v hours %v minutes %v seconds [Dec] %v degrees %v minutes %v seconds = [Galactic] [Longitude] %v degrees %v minutes %v seconds [Latitude] %v degrees %v minutes %v seconds\n", tt.args.raHours, tt.args.raMinutes, tt.args.raSeconds, tt.args.decDegrees, tt.args.decMinutes, tt.args.decSeconds, galLongDeg, galLongMin, galLongSec, galLatDeg, galLatMin, galLatSec)
			}
		})
	}
}

func TestGalacticCoordinateToEquatorialCoordinate(t *testing.T) {
	type args struct {
		galLongDeg float64
		galLongMin float64
		galLongSec float64
		galLatDeg  float64
		galLatMin  float64
		galLatSec  float64
	}
	tests := []struct {
		name           string
		args           args
		wantRAHours    float64
		wantRAMinutes  float64
		wantRASeconds  float64
		wantDecDegrees float64
		wantDecMinutes float64
		wantDecSeconds float64
	}{
		{name: "GalacticCoordinateToEquatorialCoordinate", args: args{galLongDeg: 232, galLongMin: 14, galLongSec: 52.38, galLatDeg: 51, galLatMin: 7, galLatSec: 20.16}, wantRAHours: 10, wantRAMinutes: 21, wantRASeconds: 0, wantDecDegrees: 10, wantDecMinutes: 3, wantDecSeconds: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raHours, raMinutes, raSeconds, decDegrees, decMinutes, decSeconds := GalacticCoordinateToEquatorialCoordinate(tt.args.galLongDeg, tt.args.galLongMin, tt.args.galLongSec, tt.args.galLatDeg, tt.args.galLatMin, tt.args.galLatSec)

			if raHours != tt.wantRAHours || raMinutes != tt.wantRAMinutes || raSeconds != tt.wantRASeconds || decDegrees != tt.wantDecDegrees || decMinutes != tt.wantDecMinutes || decSeconds != tt.wantDecSeconds {
				t.Errorf("GalacticCoordinateToEquatorialCoordinate() got = [RA] %v hours %v minutes %v seconds [Dec] %v degrees %v minutes %v seconds, want [RA] %v hours %v minutes %v seconds [Dec] %v degrees %v minutes %v seconds", raHours, raMinutes, raSeconds, decDegrees, decMinutes, decSeconds, tt.wantRAHours, tt.wantRAMinutes, tt.wantRASeconds, tt.wantDecDegrees, tt.wantDecMinutes, tt.wantDecSeconds)
			} else {
				fmt.Printf("Galactic coordinate to equatorial coordinate: [Galactic] [Longitude] %v degrees %v minutes %v seconds [Latitude] %v degrees %v minutes %v seconds = [RA] %v hours %v minutes %v seconds [Dec] %v degrees %v minutes %v seconds\n", tt.args.galLongDeg, tt.args.galLongMin, tt.args.galLongSec, tt.args.galLatDeg, tt.args.galLatMin, tt.args.galLatSec, raHours, raMinutes, raSeconds, decDegrees, decMinutes, decSeconds)
			}
		})
	}
}

func TestAngleBetweenTwoObjects(t *testing.T) {
	type args struct {
		raLong1HourDeg float64
		raLong1Min     float64
		raLong1Sec     float64
		decLat1Deg     float64
		decLat1Min     float64
		decLat1Sec     float64
		raLong2HourDeg float64
		raLong2Min     float64
		raLong2Sec     float64
		decLat2Deg     float64
		decLat2Min     float64
		decLat2Sec     float64
		hourOrDegree   string
	}
	tests := []struct {
		name             string
		args             args
		wantAngleDegrees float64
		wantAngleMinutes float64
		wantAngleSeconds float64
	}{
		{name: "AngleBetweenTwoObjects", args: args{raLong1HourDeg: 5, raLong1Min: 13, raLong1Sec: 31.7, decLat1Deg: -8, decLat1Min: 13, decLat1Sec: 30, raLong2HourDeg: 6, raLong2Min: 44, raLong2Sec: 13.4, decLat2Deg: -16, decLat2Min: 41, decLat2Sec: 11, hourOrDegree: "H"}, wantAngleDegrees: 23, wantAngleMinutes: 40, wantAngleSeconds: 25.86},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			angleDegrees, angleMinutes, angleSeconds := AngleBetweenTwoObjects(tt.args.raLong1HourDeg, tt.args.raLong1Min, tt.args.raLong1Sec, tt.args.decLat1Deg, tt.args.decLat1Min, tt.args.decLat1Sec, tt.args.raLong2HourDeg, tt.args.raLong2Min, tt.args.raLong2Sec, tt.args.decLat2Deg, tt.args.decLat2Min, tt.args.decLat2Sec, tt.args.hourOrDegree)

			if angleDegrees != tt.wantAngleDegrees || angleMinutes != tt.wantAngleMinutes || angleSeconds != tt.wantAngleSeconds {
				t.Errorf("AngleBetweenTwoObjects() got = [Angle] %v degrees %v minutes %v seconds, want [Angle] %v degrees %v minutes %v seconds", angleDegrees, angleMinutes, angleSeconds, tt.wantAngleDegrees, tt.wantAngleMinutes, tt.wantAngleSeconds)
			} else {
				fmt.Printf("Angle between two objects: [RA] [Long 1] %v hours/degrees %v minutes %v seconds [Dec] [Lat 1] %v degrees %v minutes %v seconds [RA] [Long 2] %v hours/degrees %v minutes %v seconds [Dec] [Lat 2] %v degrees %v minutes %v seconds [Hour or Degree] %v = [Angle] %v degrees %v minutes %v seconds\n", tt.args.raLong1HourDeg, tt.args.raLong1Min, tt.args.raLong1Sec, tt.args.decLat1Deg, tt.args.decLat1Min, tt.args.decLat1Sec, tt.args.raLong2HourDeg, tt.args.raLong2Min, tt.args.raLong2Sec, tt.args.decLat2Deg, tt.args.decLat2Min, tt.args.decLat2Sec, tt.args.hourOrDegree, angleDegrees, angleMinutes, angleSeconds)
			}
		})
	}
}

func TestRisingAndSetting(t *testing.T) {
	type args struct {
		raHours      float64
		raMinutes    float64
		raSeconds    float64
		decDeg       float64
		decMin       float64
		decSec       float64
		gwDateDay    float64
		gwDateMonth  int
		gwDateYear   int
		geogLongDeg  float64
		geogLatDeg   float64
		vertShiftDeg float64
	}
	tests := []struct {
		name              string
		args              args
		wantRiseSetStatus string
		wantUTRiseHour    float64
		wantUTRiseMinute  float64
		wantUTSetHour     float64
		wantUTSetMinute   float64
		wantAzRise        float64
		wantAzSet         float64
	}{
		{name: "Rising and Setting", args: args{raHours: 23, raMinutes: 39, raSeconds: 20, decDeg: 21, decMin: 42, decSec: 0, gwDateDay: 24, gwDateMonth: 8, gwDateYear: 2010, geogLongDeg: 64, geogLatDeg: 30, vertShiftDeg: 0.5667}, wantRiseSetStatus: "OK", wantUTRiseHour: 14, wantUTRiseMinute: 16, wantUTSetHour: 4, wantUTSetMinute: 10, wantAzRise: 64.36, wantAzSet: 295.64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			riseSetStatus, utRiseHour, utRiseMin, utSetHour, utSetMin, azRise, azSet := RisingAndSetting(tt.args.raHours, tt.args.raMinutes, tt.args.raSeconds, tt.args.decDeg, tt.args.decMin, tt.args.decSec, tt.args.gwDateDay, tt.args.gwDateMonth, tt.args.gwDateYear, tt.args.geogLongDeg, tt.args.geogLatDeg, tt.args.vertShiftDeg)

			if riseSetStatus != tt.wantRiseSetStatus || utRiseHour != tt.wantUTRiseHour || utRiseMin != tt.wantUTRiseMinute || utSetHour != tt.wantUTSetHour || utSetMin != tt.wantUTSetMinute || azRise != tt.wantAzRise || azSet != tt.wantAzSet {
				t.Errorf("RisingAndSetting() got = [Status] %v [UT Rise] %v:%v [UT Set] %v:%v [Azimuth] [Rise] %v [Set] %v, want [Status] %v [UT Rise] %v:%v [UT Set] %v:%v [Azimuth] [Rise] %v [Set] %v", riseSetStatus, utRiseHour, utRiseMin, utSetHour, utSetMin, azRise, azSet, tt.wantRiseSetStatus, tt.wantUTRiseHour, tt.wantUTRiseMinute, tt.wantUTSetHour, tt.wantUTSetMinute, tt.wantAzRise, tt.wantAzSet)
			} else {
				fmt.Printf("Rising and setting: [RA] %v hours %v minutes %v seconds [Dec] %v degrees %v minutes %v seconds [GW Date] %v/%v/%v [Geog Long/Lat] %v/%v [Vert Shift] %v degrees = [Status] %v [UT] [Rise] %v:%v [Set] %v:%v [Azimuth Rise/Set] %v/%v\n", tt.args.raHours, tt.args.raMinutes, tt.args.raSeconds, tt.args.decDeg, tt.args.decMin, tt.args.decSec, tt.args.gwDateMonth, tt.args.gwDateDay, tt.args.gwDateYear, tt.args.geogLongDeg, tt.args.geogLatDeg, tt.args.vertShiftDeg, riseSetStatus, utRiseHour, utRiseMin, utSetHour, utSetMin, azRise, azSet)
			}
		})
	}
}

func TestCorrectForPrecession(t *testing.T) {
	type args struct {
		raHour      float64
		raMinutes   float64
		raSeconds   float64
		decDeg      float64
		decMinutes  float64
		decSeconds  float64
		epoch1Day   float64
		epoch1Month int
		epoch1Year  int
		epoch2Day   float64
		epoch2Month int
		epoch2Year  int
	}
	tests := []struct {
		name                    string
		args                    args
		wantCorrectedRAHour     float64
		wantCorrectedRAMinutes  float64
		wantCorrectedRASeconds  float64
		wantCorrectedDecDeg     float64
		wantCorrectedDecMinutes float64
		wantCorrectedDecSeconds float64
	}{
		{name: "Correct for precession", args: args{raHour: 9, raMinutes: 10, raSeconds: 43, decDeg: 14, decMinutes: 23, decSeconds: 25, epoch1Day: 0.923, epoch1Month: 1, epoch1Year: 1950, epoch2Day: 1, epoch2Month: 6, epoch2Year: 1979}, wantCorrectedRAHour: 9, wantCorrectedRAMinutes: 12, wantCorrectedRASeconds: 20.18, wantCorrectedDecDeg: 14, wantCorrectedDecMinutes: 16, wantCorrectedDecSeconds: 9.12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			correctedRAHour, correctedRAMinutes, correctedRASeconds, correctedDecDeg, correctedDecMinutes, correctedDecSeconds := CorrectForPrecession(tt.args.raHour, tt.args.raMinutes, tt.args.raSeconds, tt.args.decDeg, tt.args.decMinutes, tt.args.decSeconds, tt.args.epoch1Day, tt.args.epoch1Month, tt.args.epoch1Year, tt.args.epoch2Day, tt.args.epoch2Month, tt.args.epoch2Year)

			if correctedRAHour != tt.wantCorrectedRAHour || correctedRAMinutes != tt.wantCorrectedRAMinutes || correctedRASeconds != tt.wantCorrectedRASeconds || correctedDecDeg != tt.wantCorrectedDecDeg || correctedDecMinutes != tt.wantCorrectedDecMinutes || correctedDecSeconds != tt.wantCorrectedDecSeconds {
				t.Errorf("CorrectForPrecession() got = [RA] %v hours %v minutes %v seconds [Declination] %v degrees %v minutes %v seconds, want [RA] %v hours %v minutes %v seconds [Declination] %v degrees %v minutes %v seconds", correctedRAHour, correctedRAMinutes, correctedRASeconds, correctedDecDeg, correctedDecMinutes, correctedDecSeconds, tt.wantCorrectedRAHour, tt.wantCorrectedRAMinutes, tt.wantCorrectedRASeconds, tt.wantCorrectedDecDeg, tt.wantCorrectedDecMinutes, tt.wantCorrectedDecSeconds)
			} else {
				fmt.Printf("Correct for precession: [RA] %v hours %v minutes %v seconds [Declination] %v degrees %v minutes %v seconds [Epoch 1] %v/%v/%v [Epoch 2] %v/%v/%v = [Corrected RA] %v hours %v minutes %v seconds [Corrected Declination] %v degrees %v minutes %v seconds\n", tt.args.raHour, tt.args.raMinutes, tt.args.raSeconds, tt.args.decDeg, tt.args.decMinutes, tt.args.decSeconds, tt.args.epoch1Month, tt.args.epoch1Day, tt.args.epoch1Year, tt.args.epoch2Month, tt.args.epoch2Day, tt.args.epoch2Year, correctedRAHour, correctedRAMinutes, correctedRASeconds, correctedDecDeg, correctedDecMinutes, correctedDecSeconds)
			}
		})
	}
}

func TestNutationInEclipticLongitudeAndObliquity(t *testing.T) {
	type args struct {
		greenwichDay   float64
		greenwichMonth int
		greenwichYear  int
	}
	tests := []struct {
		name                           string
		args                           args
		wantNutationInLongitudeDegrees float64
		wantNutationInObliquityDegrees float64
	}{
		{name: "NutationInEclipticLongitudeAndObliquity", args: args{greenwichDay: 1, greenwichMonth: 9, greenwichYear: 1988}, wantNutationInLongitudeDegrees: 0.001525808, wantNutationInObliquityDegrees: 0.0025671},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nutInLongDeg, nutInOblDeg := NutationInEclipticLongitudeAndObliquity(tt.args.greenwichDay, tt.args.greenwichMonth, tt.args.greenwichYear)

			nutInLongDeg = util.RoundFloat64(nutInLongDeg, 9)
			nutInOblDeg = util.RoundFloat64(nutInOblDeg, 7)

			if nutInLongDeg != tt.wantNutationInLongitudeDegrees || nutInOblDeg != tt.wantNutationInObliquityDegrees {
				t.Errorf("NutationInEclipticLongitudeAndObliquity() got = [Nutation in Longitude] %v degrees [Nutation in Obliquity] %v degrees, want [Nutation in Longitude] %v degrees [Nutation in Obliquity] %v degrees", nutInLongDeg, nutInOblDeg, tt.wantNutationInLongitudeDegrees, tt.wantNutationInObliquityDegrees)
			} else {
				fmt.Printf("Nutation in ecliptic longitude and obliquity: [Greenwich Date] %v/%v/%v = [Nutation in Longitude] %v degrees [Nutation in Obliquity] %v degrees\n", tt.args.greenwichMonth, tt.args.greenwichDay, tt.args.greenwichYear, nutInLongDeg, nutInOblDeg)
			}
		})
	}
}
