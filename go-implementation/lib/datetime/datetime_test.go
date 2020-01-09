package datetime

import "testing"
import "fmt"
import "../util"
import "../macros"

func TestDateOfEaster(t *testing.T) {
	var month, day, year int = GetDateOfEaster(2009)

	if month == 4 && day == 12 && year == 2009 {
		fmt.Printf("Date of Easter for %d is %d/%d/%d\n", year, month, day, year)
	} else {
		t.Errorf("Expected 4/12/2009, actual is %d/%d/%d", month, day, year)
	}
}

func TestCivilDateToDayNumbers(t *testing.T) {
	testCivilDateToDayNumber(t, 1, 1, 2000, 1)
	testCivilDateToDayNumber(t, 3, 1, 2000, 61)
	testCivilDateToDayNumber(t, 6, 1, 2003, 152)
	testCivilDateToDayNumber(t, 11, 27, 2009, 331)
}

func testCivilDateToDayNumber(t *testing.T, month int, day int, year int, expectedValue int) {
	dayNumber := CivilDateToDayNumber(month, day, year)

	if dayNumber == expectedValue {
		fmt.Printf("Day number: [Date] %d/%d/%d = [Day Number] %d\n", month, day, year, dayNumber)
	} else {
		t.Errorf("Expected %d, actual is %d\n", expectedValue, dayNumber)
	}
}

func TestCivilTimeToDecimalHours(t *testing.T) {
	civilHours := 18.0
	civilMinutes := 31.0
	civilSeconds := 27.0

	expectedResult := 18.52416667

	decimalHours := CivilTimeToDecimalHours(civilHours, civilMinutes, civilSeconds)
	decimalHours = util.RoundFloat64(decimalHours, 8)

	if decimalHours == expectedResult {
		fmt.Printf("Civil time to decimal hours: [Time] %d:%d:%d = [Decimal Hours] %.8f\n", int(civilHours), int(civilMinutes), int(civilSeconds), decimalHours)

	} else {
		t.Errorf("Expected %.8f, actual is %.8f\n", expectedResult, decimalHours)
	}
}

func TestDecimalHoursToCivilTime(t *testing.T) {
	decimalHours := 18.52416667

	expectedCivilHours := 18
	expectedCivilMinutes := 31
	expectedCivilSeconds := 27

	civilHours, civilMinutes, civilSeconds := DecimalHoursToCivilTime(decimalHours)

	if civilHours == expectedCivilHours && civilMinutes == expectedCivilMinutes && civilSeconds == expectedCivilSeconds {
		fmt.Printf("Decimal hours to civil time: [Decimal Hours] %.8f = [Civil Time] %d:%d:%d\n", decimalHours, civilHours, civilMinutes, civilSeconds)
	} else {
		t.Errorf("Expected %d:%d:%d, actual is %d:%d:%d\n", expectedCivilHours, expectedCivilMinutes, expectedCivilSeconds, civilHours, civilMinutes, civilSeconds)
	}
}

func TestLocalCivilTimeToUniversalTime(t *testing.T) {
	lctHours := 3.0
	lctMinutes := 37.0
	lctSeconds := 0.0
	isDaylightSavings := true
	zoneCorrection := 4
	localDay := 1.0
	localMonth := 7
	localYear := 2013

	expectedUTHours := 22
	expectedUTMinutes := 37
	expectedUTSeconds := 0
	expectedGWDay := 30
	expectedGWMonth := 6
	expectedGWYear := 2013

	utHours, utMinutes, utSeconds, gwDay, gwMonth, gwYear := LocalCivilTimeToUniversalTime(lctHours, lctMinutes, lctSeconds, isDaylightSavings, zoneCorrection, localDay, localMonth, localYear)

	if utHours == expectedUTHours && utMinutes == expectedUTMinutes && utSeconds == expectedUTSeconds && gwDay == expectedGWDay && gwMonth == expectedGWMonth && gwYear == expectedGWYear {
		fmt.Printf("Civil time to universal time: [LCT] %d:%d:%d [DST?] %t [ZC] %d [Local Date] %d/%d/%d = [UT] %d:%d:%d [GWD] %d/%d/%d\n", int(lctHours), int(lctMinutes), int(lctSeconds), isDaylightSavings, zoneCorrection, int(localMonth), int(localDay), int(localYear), utHours, utMinutes, utSeconds, gwMonth, gwDay, gwYear)
	} else {
		t.Errorf("Expected [UT] %d:%d:%d [GWD] %d/%d/%d, actual is [UT] %d:%d:%d [GWD] %d/%d/%d\n", expectedUTHours, expectedUTMinutes, expectedUTSeconds, expectedGWMonth, expectedGWDay, expectedGWYear, utHours, utMinutes, utSeconds, gwMonth, gwDay, gwYear)
	}
}
func TestUniversalTimeToLocalCivilTime(t *testing.T) {
	utHours := 22.0
	utMinutes := 37.0
	utSeconds := 0.0
	isDaylightSavings := true
	zoneCorrection := 4
	gwDay := 30
	gwMonth := 6
	gwYear := 2013

	expectedLCTHours := 3
	expectedLCTMinutes := 37
	expectedLCTSeconds := 0
	expectedLocalDay := 1
	expectedLocalMonth := 7
	expectedLocalYear := 2013

	lctHours, lctMinutes, lctSeconds, localDay, localMonth, localYear := UniversalTimeToLocalCivilTime(utHours, utMinutes, utSeconds, isDaylightSavings, zoneCorrection, gwDay, gwMonth, gwYear)

	if lctHours == expectedLCTHours && lctMinutes == expectedLCTMinutes && lctSeconds == expectedLCTSeconds && localDay == expectedLocalDay && localMonth == expectedLocalMonth && localYear == expectedLocalYear {
		fmt.Printf("Universal time to civil time: [UT] %d:%d:%d [DST?] %t [ZC] %d [GWD] %d/%d/%d = [LCT] %d:%d:%d [Local Date] %d/%d/%d\n", int(utHours), int(utMinutes), int(utSeconds), isDaylightSavings, zoneCorrection, gwMonth, gwDay, gwYear, int(lctHours), int(lctMinutes), int(lctSeconds), int(localMonth), int(localDay), int(localYear))
	} else {
		t.Errorf("Expected [LCT] %d:%d:%d [Local Date] %d/%d/%d, actual is [LCT] %d:%d:%d [Local Date] %d/%d/%d\n", expectedLCTHours, expectedLCTMinutes, expectedLCTSeconds, expectedLocalMonth, expectedLocalDay, expectedLocalYear, lctHours, lctMinutes, lctSeconds, localMonth, localDay, localYear)

	}
}

func TestUniversalTimeToGreenwichSiderealTime(t *testing.T) {
	utHours := 14.0
	utMinutes := 36.0
	utSeconds := 51.67
	gwDay := 22.0
	gwMonth := 4
	gwYear := 1980

	expectedGSTHours := 4
	expectedGSTMinutes := 40
	expectedGSTSeconds := 5.23

	gstHours, gstMinutes, gstSeconds := UniversalTimeToGreenwichSiderealTime(utHours, utMinutes, utSeconds, gwDay, gwMonth, gwYear)

	if gstHours == expectedGSTHours && gstMinutes == expectedGSTMinutes && gstSeconds == expectedGSTSeconds {
		fmt.Printf("Universal time to Greenwich sidereal time: [UT] %d:%d:%.2f [Greenwich Date] %d/%d/%d = [Greenwich Sidereal Time] %d:%d:%.2f\n", int(utHours), int(utMinutes), utSeconds, gwMonth, int(gwDay), gwYear, gstHours, gstMinutes, gstSeconds)
	} else {
		t.Errorf("Expected GST %d:%d:%.2f, actual %d:%d:%.2f\n", expectedGSTHours, expectedGSTMinutes, expectedGSTSeconds, gstHours, gstMinutes, gstSeconds)
	}

}

func TestGreenwichSiderealTimeToUniversalTime(t *testing.T) {
	gstHours := 4.0
	gstMinutes := 40.0
	gstSeconds := 5.23
	gwDay := 22.0
	gwMonth := 4
	gwYear := 1980

	expectedUTHours := 14
	expectedUTMinutes := 36
	expectedUTSeconds := 51.67
	expectedWarningFlag := "OK"

	utHours, utMinutes, utSeconds, warningFlag := GreenwichSiderealTimeToUniversalTime(gstHours, gstMinutes, gstSeconds, gwDay, gwMonth, gwYear)

	if utHours == expectedUTHours && utMinutes == expectedUTMinutes && utSeconds == expectedUTSeconds && warningFlag == expectedWarningFlag {
		fmt.Printf("Greenwich sidereal time to universal time: [Greenwich Sidereal Time] %d:%d:%.2f [Greenwich Date] %d/%d/%d = [UT] %d:%d:%.2f [Warning Flag] %s\n", int(gstHours), int(gstMinutes), gstSeconds, gwMonth, int(gwDay), gwYear, utHours, utMinutes, utSeconds, warningFlag)
	} else {
		t.Errorf("Expected [UT] %d:%d:%.2f [Warning Flag] %s, actual [UT] %d:%d:%.2f [Warning Flag] %s\n", expectedUTHours, expectedUTMinutes, expectedUTSeconds, expectedWarningFlag, utHours, utMinutes, utSeconds, warningFlag)
	}
}

func TestGreenwichSiderealTimeToLocalSiderealTime(t *testing.T) {
	gstHour := 4.0
	gstMinutes := 40.0
	gstSeconds := 5.23
	geographicalLongitude := -64.0

	expectedLSTHours := 0
	expectedLSTMinutes := 24
	expectedLSTSeconds := 5.23

	lstHours, lstMinutes, lstSeconds := GreenwichSiderealTimeToLocalSiderealTime(gstHour, gstMinutes, gstSeconds, geographicalLongitude)

	if lstHours == expectedLSTHours && lstMinutes == expectedLSTMinutes && lstSeconds == expectedLSTSeconds {
		fmt.Printf("Greenwich sidereal time to local sidereal time: [GST] %.0f:%.0f:%.2f [Geographical Longitude] %.1f = [LST] %d:%d:%.2f\n", gstHour, gstMinutes, gstSeconds, geographicalLongitude, lstHours, lstMinutes, lstSeconds)
	} else {
		t.Errorf("Expected [LST] %d:%d:%.2f, actual %d:%d:%.2f\n", expectedLSTHours, expectedLSTMinutes, expectedLSTSeconds, lstHours, lstMinutes, lstSeconds)
	}
}

func TestLocalSiderealTimeToGreenwichSiderealTime(t *testing.T) {
	lstHour := 0.0
	lstMinutes := 24.0
	lstSeconds := 5.23
	geographicalLongitude := -64.0

	expectedGSTHours := 4
	expectedGSTMinutes := 40
	expectedGSTSeconds := 5.23

	gstHours, gstMinutes, gstSeconds := LocalSiderealTimeToGreenwichSiderealTime(lstHour, lstMinutes, lstSeconds, geographicalLongitude)

	if gstHours == expectedGSTHours && gstMinutes == expectedGSTMinutes && gstSeconds == expectedGSTSeconds {
		fmt.Printf("Local sidereal time to greenwich sidereal time: [LST] %.0f:%.0f:%.2f [Geographical Longitude] %.1f = [GST] %d:%d:%.2f\n", lstHour, lstMinutes, lstSeconds, geographicalLongitude, gstHours, gstMinutes, gstSeconds)
	} else {
		t.Errorf("Expected %d:%d:%.2f, actual %d:%d:%.2f\n", expectedGSTHours, expectedGSTMinutes, expectedGSTSeconds, gstHours, gstMinutes, gstSeconds)
	}
}

func TestJulianDateToDayOfWeek(t *testing.T) {
	julianDate := 2455001.5

	expectedDayOfWeek := "Friday"

	dayOfWeek := macros.FDOW(julianDate)

	if dayOfWeek == expectedDayOfWeek {
		fmt.Printf("Julian date to day of week: [Julian Date] %.1f = [Day of Week] %s\n", julianDate, dayOfWeek)
	} else {
		t.Errorf("Expected %s, actual %s\n", expectedDayOfWeek, dayOfWeek)
	}
}
