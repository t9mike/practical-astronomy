package macros

import (
	"fmt"
	"testing"
)

func TestFDOW(t *testing.T) {
	type args struct {
		julianDate float64
	}
	tests := []struct {
		name          string
		args          args
		wantDayOfWeek string
	}{
		{name: "FDOW", args: args{julianDate: 2455001.5}, wantDayOfWeek: "Friday"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dayOfWeek := FDOW(tt.args.julianDate)

			if dayOfWeek != tt.wantDayOfWeek {
				t.Errorf("FDOW() = %v, want %v", dayOfWeek, tt.wantDayOfWeek)
			} else {
				fmt.Printf("Julian date to day of week: [Julian Date] %v = [Day of Week] %v\n", tt.args.julianDate, dayOfWeek)
			}
		})
	}
}