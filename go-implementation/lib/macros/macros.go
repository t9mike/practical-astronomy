package macros

import "math"
import "../util"

// HmsDh converts a Civil Time (hours,minutes,seconds) to Decimal Hours
//
// Original macro name: HMSDH
func HmsDh(hours float64, minutes float64, seconds float64) float64 {
	fHours := hours
	fMinutes := minutes
	fSeconds := seconds

	a := math.Abs(fSeconds) / 60
	b := (math.Abs(fMinutes) + a) / 60
	c := math.Abs(fHours) + b

	if fHours < 0 || fMinutes < 0 || fSeconds < 0 {
		return -c
	}

	return c
}

// DhHour returns the hour part of a Decimal Hours.
//
// Original macro name: DHHour
func DhHour(decimalHours float64) int {
	a := math.Abs(decimalHours)
	b := a * 3600
	c := util.RoundFloat64(b-60*math.Floor(b/60), 2)
	e := b
	if c == 60 {
		e = b + 60.0
	}

	if decimalHours < 0.0 {
		return int(-math.Floor(e / 3600))
	}

	return int(math.Floor(e / 3600))
}

// DhMin returns the minutes part of a Decimal Hours.
//
// Original macro name: DHMin
func DhMin(decimalHours float64) int {
	a := math.Abs(decimalHours)
	b := a * 3600
	c := util.RoundFloat64(b-60*math.Floor(b/60), 2)
	e := b
	if c == 60 {
		e = b + 60.0
	}

	return int(math.Mod(math.Floor(e/60), 60))
}

// DhSec returns the seconds part of a Decimal Hours.
//
// Original macro name: DHSec
func DhSec(decimalHours float64) float64 {
	a := math.Abs(decimalHours)
	b := a * 3600
	c := util.RoundFloat64(b-60*math.Floor(b/60), 2)
	d := c
	if c == 60.0 {
		d = 0
	}

	return d
}
