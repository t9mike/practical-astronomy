package util

import (
	"fmt"
	"testing"
)

func TestIsLeapYear(t *testing.T) {
	type args struct {
		inputYear int
	}
	tests := []struct {
		name           string
		args           args
		wantIsLeapYear bool
	}{
		{name: "IsLeapYear", args: args{inputYear: 2017}, wantIsLeapYear: false},
		{name: "IsLeapYear", args: args{inputYear: 2020}, wantIsLeapYear: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isLeapYear := IsLeapYear(tt.args.inputYear)

			if isLeapYear != tt.wantIsLeapYear {
				t.Errorf("IsLeapYear() = %t, want %t", isLeapYear, tt.wantIsLeapYear)
			} else {
				fmt.Printf("IsLeapYear: [Year] %d = [IsLeapYear] %t\n", tt.args.inputYear, isLeapYear)
			}
		})
	}
}

func TestRoundFloat64(t *testing.T) {
	type args struct {
		inputValue float64
		places     int
	}
	tests := []struct {
		name             string
		args             args
		wantRoundedFloat float64
	}{
		{name: "RoundFloat64", args: args{inputValue: 12.227, places: 2}, wantRoundedFloat: 12.23},
		{name: "RoundFloat64", args: args{inputValue: 12.243, places: 2}, wantRoundedFloat: 12.24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			roundedFloat := RoundFloat64(tt.args.inputValue, tt.args.places)

			if roundedFloat != tt.wantRoundedFloat {
				t.Errorf("RoundFloat64() = %v, want %v", roundedFloat, tt.wantRoundedFloat)
			} else {
				fmt.Printf("Round float value: [Original] %v [Places] %d = [Rounded] %v\n", tt.args.inputValue, tt.args.places, roundedFloat)
			}
		})
	}
}
