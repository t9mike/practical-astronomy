package datetime

import (
	"fmt"
	"testing"

	"../util"
)

func TestGetDateOfEaster(t *testing.T) {
	type args struct {
		inputYear int
	}
	tests := []struct {
		name      string
		args      args
		wantMonth int
		wantDay   int
		wantYear  int
	}{
		{name: "GetDateOfEaster", args: args{inputYear: 2009}, wantMonth: 4, wantDay: 12, wantYear: 2009},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			month, day, year := GetDateOfEaster(tt.args.inputYear)
			if month != tt.wantMonth || day != tt.wantDay || year != tt.wantYear {
				t.Errorf("GetDateOfEaster() got = %v/%v/%v, want %v/%v/%v", month, day, year, tt.wantMonth, tt.wantDay, tt.wantYear)
			} else {
				fmt.Printf("Date of Easter for %v is %v/%v/%v\n", tt.args.inputYear, month, day, year)
			}
		})
	}
}

func TestCivilDateToDayNumber(t *testing.T) {
	type args struct {
		month int
		day   int
		year  int
	}
	tests := []struct {
		name          string
		args          args
		wantDayNumber int
	}{
		{name: "CivilDateToDayNumber", args: args{month: 1, day: 1, year: 2000}, wantDayNumber: 1},
		{name: "CivilDateToDayNumber", args: args{month: 3, day: 1, year: 2000}, wantDayNumber: 61},
		{name: "CivilDateToDayNumber", args: args{month: 6, day: 1, year: 2003}, wantDayNumber: 152},
		{name: "CivilDateToDayNumber", args: args{month: 11, day: 27, year: 2009}, wantDayNumber: 331},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dayNumber := CivilDateToDayNumber(tt.args.month, tt.args.day, tt.args.year)

			if dayNumber != tt.wantDayNumber {
				t.Errorf("CivilDateToDayNumber() = %v, want %v", dayNumber, tt.wantDayNumber)
			} else {
				fmt.Printf("Day number: [Date] %v/%v/%v = [Day Number] %v\n", tt.args.month, tt.args.day, tt.args.year, dayNumber)
			}
		})
	}
}

func TestCivilTimeToDecimalHours(t *testing.T) {
	type args struct {
		hours   float64
		minutes float64
		seconds float64
	}
	tests := []struct {
		name             string
		args             args
		wantDecimalHours float64
	}{
		{name: "CivilTimeToDecimalHours", args: args{hours: 18.0, minutes: 31.0, seconds: 27.0}, wantDecimalHours: 18.52416667},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decimalHours := util.RoundFloat64(CivilTimeToDecimalHours(tt.args.hours, tt.args.minutes, tt.args.seconds), 8)

			if decimalHours != tt.wantDecimalHours {
				t.Errorf("CivilTimeToDecimalHours() = %v, want %v", decimalHours, tt.wantDecimalHours)
			} else {
				fmt.Printf("Civil time to decimal hours: [Time] %v:%v:%v = [Decimal Hours] %v\n", int(tt.args.hours), int(tt.args.minutes), int(tt.args.seconds), decimalHours)
			}
		})
	}
}

func TestDecimalHoursToCivilTime(t *testing.T) {
	type args struct {
		decimalHours float64
	}
	tests := []struct {
		name             string
		args             args
		wantCivilHours   int
		wantCivilMinutes int
		wantCivilSeconds int
	}{
		{name: "DecimalHoursToCivilTime", args: args{decimalHours: 18.52416667}, wantCivilHours: 18, wantCivilMinutes: 31, wantCivilSeconds: 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			civilHours, civilMinutes, civilSeconds := DecimalHoursToCivilTime(tt.args.decimalHours)
			if civilHours != tt.wantCivilHours || civilMinutes != tt.wantCivilMinutes || civilSeconds != tt.wantCivilSeconds {
				t.Errorf("DecimalHoursToCivilTime() got = %v:%v:%v, want %v:%v:%v", civilHours, civilMinutes, civilSeconds, tt.wantCivilHours, tt.wantCivilMinutes, tt.wantCivilSeconds)
			} else {
				fmt.Printf("Decimal hours to civil time: [Decimal Hours] %v = [Civil Time] %v:%v:%v\n", tt.args.decimalHours, civilHours, civilMinutes, civilSeconds)
			}
		})
	}
}

func TestLocalCivilTimeToUniversalTime(t *testing.T) {
	type args struct {
		lctHours          float64
		lctMinutes        float64
		lctSeconds        float64
		isDayLightSavings bool
		zoneCorrection    int
		localDay          float64
		localMonth        int
		localYear         int
	}
	tests := []struct {
		name          string
		args          args
		wantUTHours   int
		wantUTMinutes int
		wantUTSeconds int
		wantGWDay     int
		wantGWMonth   int
		wantGWYear    int
	}{
		{name: "LocalCivilTimeToUniversalTime", args: args{lctHours: 3.0, lctMinutes: 37.0, lctSeconds: 0.0, isDayLightSavings: true, zoneCorrection: 4, localDay: 1.0, localMonth: 7, localYear: 2013}, wantUTHours: 22, wantUTMinutes: 37, wantUTSeconds: 0, wantGWDay: 30, wantGWMonth: 6, wantGWYear: 2013},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utHours, utMinutes, utSeconds, gwDay, gwMonth, gwYear := LocalCivilTimeToUniversalTime(tt.args.lctHours, tt.args.lctMinutes, tt.args.lctSeconds, tt.args.isDayLightSavings, tt.args.zoneCorrection, tt.args.localDay, tt.args.localMonth, tt.args.localYear)

			if utHours != tt.wantUTHours || utMinutes != tt.wantUTMinutes || utSeconds != tt.wantUTSeconds || gwDay != tt.wantGWDay || gwMonth != tt.wantGWMonth || gwYear != tt.wantGWYear {
				t.Errorf("LocalCivilTimeToUniversalTime() got = [UT] %v:%v:%v [GWD] %v/%v/%v, want [UT] %v:%v:%v [GWD] %v/%v/%v", utHours, utMinutes, utSeconds, gwMonth, gwDay, gwYear, tt.wantUTHours, tt.wantUTMinutes, tt.wantUTSeconds, tt.wantGWMonth, tt.wantGWDay, tt.wantGWYear)
			} else {
				fmt.Printf("Civil time to universal time: [LCT] %v:%v:%v [DST?] %v [ZC] %v [Local Date] %v/%v/%v = [UT] %v:%v:%v [GWD] %v/%v/%v\n", int(tt.args.lctHours), int(tt.args.lctMinutes), int(tt.args.lctSeconds), tt.args.isDayLightSavings, tt.args.zoneCorrection, int(tt.args.localMonth), int(tt.args.localDay), int(tt.args.localYear), utHours, utMinutes, utSeconds, gwMonth, gwDay, gwYear)
			}
		})
	}
}

func TestUniversalTimeToLocalCivilTime(t *testing.T) {
	type args struct {
		utHours           float64
		utMinutes         float64
		utSeconds         float64
		isDaylightSavings bool
		zoneCorrection    int
		gwDay             int
		gwMonth           int
		gwYear            int
	}
	tests := []struct {
		name           string
		args           args
		wantLCTHours   int
		wantLCTMinutes int
		wantLCTSeconds int
		wantLocalDay   int
		wantLocalMonth int
		wantLocalYear  int
	}{
		{name: "UniversalTimeToLocalCivilTime", args: args{utHours: 22.0, utMinutes: 37.0, utSeconds: 0.0, isDaylightSavings: true, zoneCorrection: 4, gwDay: 30, gwMonth: 6, gwYear: 2013}, wantLCTHours: 3, wantLCTMinutes: 37, wantLCTSeconds: 0, wantLocalDay: 1, wantLocalMonth: 7, wantLocalYear: 2013},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lctHours, lctMinutes, lctSeconds, localDay, localMonth, localYear := UniversalTimeToLocalCivilTime(tt.args.utHours, tt.args.utMinutes, tt.args.utSeconds, tt.args.isDaylightSavings, tt.args.zoneCorrection, tt.args.gwDay, tt.args.gwMonth, tt.args.gwYear)

			if lctHours != tt.wantLCTHours || lctMinutes != tt.wantLCTMinutes || lctSeconds != tt.wantLCTSeconds || localDay != tt.wantLocalDay || localMonth != tt.wantLocalMonth || localYear != tt.wantLocalYear {
				t.Errorf("UniversalTimeToLocalCivilTime() got = [LCT] %v:%v:%v [Local Date] %v/%v/%v, want [LCT] %v:%v:%v [Local Date] %v/%v/%v", lctHours, lctMinutes, lctSeconds, localMonth, localDay, localYear, tt.wantLCTHours, tt.wantLCTMinutes, tt.wantLCTSeconds, tt.wantLocalMonth, tt.wantLocalDay, tt.wantLocalYear)
			} else {
				fmt.Printf("Universal time to civil time: [UT] %v:%v:%v [DST?] %v [ZC] %v [GWD] %v/%v/%v = [LCT] %v:%v:%v [Local Date] %v/%v/%v\n", int(tt.args.utHours), int(tt.args.utMinutes), int(tt.args.utSeconds), tt.args.isDaylightSavings, tt.args.zoneCorrection, tt.args.gwMonth, tt.args.gwDay, tt.args.gwYear, int(lctHours), int(lctMinutes), int(lctSeconds), int(localMonth), int(localDay), int(localYear))
			}
		})
	}
}

func TestUniversalTimeToGreenwichSiderealTime(t *testing.T) {
	type args struct {
		utHours   float64
		utMinutes float64
		utSeconds float64
		gwDay     float64
		gwMonth   int
		gwYear    int
	}
	tests := []struct {
		name           string
		args           args
		wantGSTHours   int
		wantGSTMinutes int
		wantGSTSeconds float64
	}{
		{name: "UniversalTimeToGreenwichSiderealTime", args: args{utHours: 14.0, utMinutes: 36.0, utSeconds: 51.67, gwDay: 22.0, gwMonth: 4, gwYear: 1980}, wantGSTHours: 4, wantGSTMinutes: 40, wantGSTSeconds: 5.23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gstHours, gstMinutes, gstSeconds := UniversalTimeToGreenwichSiderealTime(tt.args.utHours, tt.args.utMinutes, tt.args.utSeconds, tt.args.gwDay, tt.args.gwMonth, tt.args.gwYear)

			if gstHours != tt.wantGSTHours || gstMinutes != tt.wantGSTMinutes || gstSeconds != tt.wantGSTSeconds {
				t.Errorf("UniversalTimeToGreenwichSiderealTime() got = GST %v:%v:%v, want GST %v:%v:%v", gstHours, gstMinutes, gstSeconds, tt.wantGSTHours, tt.wantGSTMinutes, tt.wantGSTSeconds)
			} else {
				fmt.Printf("Universal time to Greenwich sidereal time: [UT] %v:%v:%v [Greenwich Date] %v/%v/%v = [Greenwich Sidereal Time] %v:%v:%v\n", int(tt.args.utHours), int(tt.args.utMinutes), tt.args.utSeconds, tt.args.gwMonth, int(tt.args.gwDay), tt.args.gwYear, gstHours, gstMinutes, gstSeconds)
			}
		})
	}
}

func TestGreenwichSiderealTimeToUniversalTime(t *testing.T) {
	type args struct {
		gstHours   float64
		gstMinutes float64
		gstSeconds float64
		gwDay      float64
		gwMonth    int
		gwYear     int
	}
	tests := []struct {
		name            string
		args            args
		wantUTHours     int
		wantUTMinutes   int
		wantUTSeconds   float64
		wantWarningFlag string
	}{
		{name: "GreenwichSiderealTimeToUniversalTime", args: args{gstHours: 4.0, gstMinutes: 40.0, gstSeconds: 5.23, gwDay: 22.0, gwMonth: 4, gwYear: 1980}, wantUTHours: 14, wantUTMinutes: 36, wantUTSeconds: 51.67, wantWarningFlag: "OK"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utHours, utMinutes, utSeconds, warningFlag := GreenwichSiderealTimeToUniversalTime(tt.args.gstHours, tt.args.gstMinutes, tt.args.gstSeconds, tt.args.gwDay, tt.args.gwMonth, tt.args.gwYear)
			if utHours != tt.wantUTHours || utMinutes != tt.wantUTMinutes || utSeconds != tt.wantUTSeconds || warningFlag != tt.wantWarningFlag {
				t.Errorf("GreenwichSiderealTimeToUniversalTime() got = [UT] %v:%v:%v [Warning Flag] %v, want [UT] %v:%v:%v [Warning Flag] %v", utHours, utMinutes, utSeconds, warningFlag, tt.wantUTHours, tt.wantUTMinutes, tt.wantUTSeconds, tt.wantWarningFlag)
			} else {
				fmt.Printf("Greenwich sidereal time to universal time: [Greenwich Sidereal Time] %v:%v:%v [Greenwich Date] %v/%v/%v = [UT] %v:%v:%v [Warning Flag] %v\n", int(tt.args.gstHours), int(tt.args.gstMinutes), tt.args.gstSeconds, tt.args.gwMonth, int(tt.args.gwDay), tt.args.gwYear, utHours, utMinutes, utSeconds, warningFlag)
			}
		})
	}
}

func TestGreenwichSiderealTimeToLocalSiderealTime(t *testing.T) {
	type args struct {
		gstHour               float64
		gstMinutes            float64
		gstSeconds            float64
		geographicalLongitude float64
	}
	tests := []struct {
		name           string
		args           args
		wantLSTHours   int
		wantLSTMinutes int
		wantLSTSeconds float64
	}{
		{name: "GreenwichSiderealTimeToLocalSiderealTime", args: args{gstHour: 4.0, gstMinutes: 40.0, gstSeconds: 5.23, geographicalLongitude: -64.0}, wantLSTHours: 0, wantLSTMinutes: 24, wantLSTSeconds: 5.23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lstHours, lstMinutes, lstSeconds := GreenwichSiderealTimeToLocalSiderealTime(tt.args.gstHour, tt.args.gstMinutes, tt.args.gstSeconds, tt.args.geographicalLongitude)

			if lstHours != tt.wantLSTHours || lstMinutes != tt.wantLSTMinutes || lstSeconds != tt.wantLSTSeconds {
				t.Errorf("GreenwichSiderealTimeToLocalSiderealTime() got = [LST] %v:%v:%v, want [LST] %v:%v:%v", lstHours, lstMinutes, lstSeconds, tt.wantLSTHours, tt.wantLSTMinutes, tt.wantLSTSeconds)
			} else {
				fmt.Printf("Greenwich sidereal time to local sidereal time: [GST] %v:%v:%v [Geographical Longitude] %v = [LST] %v:%v:%v\n", tt.args.gstHour, tt.args.gstMinutes, tt.args.gstSeconds, tt.args.geographicalLongitude, lstHours, lstMinutes, lstSeconds)
			}
		})
	}
}

func TestLocalSiderealTimeToGreenwichSiderealTime(t *testing.T) {
	type args struct {
		lstHours              float64
		lstMinutes            float64
		lstSeconds            float64
		geographicalLongitude float64
	}
	tests := []struct {
		name           string
		args           args
		wantGSTHours   int
		wantGSTMinutes int
		wantGSTSeconds float64
	}{
		{name: "LocalSiderealTimeToGreenwichSiderealTime", args: args{lstHours: 0.0, lstMinutes: 24.0, lstSeconds: 5.23, geographicalLongitude: -64.0}, wantGSTHours: 4, wantGSTMinutes: 40, wantGSTSeconds: 5.23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gstHours, gstMinutes, gstSeconds := LocalSiderealTimeToGreenwichSiderealTime(tt.args.lstHours, tt.args.lstMinutes, tt.args.lstSeconds, tt.args.geographicalLongitude)

			if gstHours != tt.wantGSTHours || gstMinutes != tt.wantGSTMinutes || gstSeconds != tt.wantGSTSeconds {
				t.Errorf("LocalSiderealTimeToGreenwichSiderealTime() got = %v:%v:%v, want %v:%v:%v", gstHours, gstMinutes, gstSeconds, tt.wantGSTHours, tt.wantGSTMinutes, tt.wantGSTSeconds)
			} else {
				fmt.Printf("Local sidereal time to greenwich sidereal time: [LST] %v:%v:%v [Geographical Longitude] %v = [GST] %v:%v:%v\n", tt.args.lstHours, tt.args.lstMinutes, tt.args.lstSeconds, tt.args.geographicalLongitude, gstHours, gstMinutes, gstSeconds)
			}
		})
	}
}
