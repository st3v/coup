package fleet_test

import (
	"testing"

	"github.com/st3v/coup/fleet"
)

// TestMinNumFE contains negative test cases for fleet.MinNumFE
func TestErrorMinNumFE(t *testing.T) {
	tests := []struct {
		scooters []int
		c        int
		p        int
		want     error
	}{
		{[]int{}, 1, 1, fleet.ErrNumDistricts},
		{make([]int, 101), 1, 1, fleet.ErrNumDistricts},
		{[]int{-1}, 1, 1, fleet.ErrNumScootersInDistrict},
		{[]int{1001}, 1, 1, fleet.ErrNumScootersInDistrict},
		{[]int{1}, 0, 1, fleet.ErrNumScootersForFM},
		{[]int{1}, 1000, 1, fleet.ErrNumScootersForFM},
		{[]int{1}, 1, 0, fleet.ErrNumScootersForFE},
		{[]int{1}, 1, 1001, fleet.ErrNumScootersForFE},
	}

	for i, tc := range tests {
		if _, got := fleet.MinNumFE(tc.scooters, tc.c, tc.p); got != tc.want {
			t.Errorf("unexpected result for test %d: want '%v', got '%v'", i, tc.want, got)
		}
	}
}

// TestMinNumFE contains positive test cases for fleet.MinNumFE
func TestMinNumFE(t *testing.T) {
	tests := []struct {
		scooters []int
		c        int
		p        int
		want     int
	}{
		{[]int{10, 15}, 12, 5, 3},
		{[]int{11, 15, 13}, 9, 5, 7},
		{[]int{999}, 999, 1000, 0},
		{[]int{1000}, 999, 1000, 1},
		{make([]int, 1000), 1, 1, 0},
		{[]int{1, 2}, 1, 1, 2},
		{[]int{0}, 1, 1, 0},
		{[]int{2, 4, 3}, 3, 3, 3},
	}

	for i, tc := range tests {
		if got, _ := fleet.MinNumFE(tc.scooters, tc.c, tc.p); got != tc.want {
			t.Errorf("unexpected result for test %d: want %d, got %d", i, tc.want, got)
		}
	}
}
