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
