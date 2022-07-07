package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isError  bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0, true},
	{"expect 5", 50.0, 10.0, 5, false},
	// {"expect fraction", -1, -777.0, 5, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := Divide(tt.dividend, tt.divisor)
		if tt.isError {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one anyway", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

// func TestDivide(t *testing.T) {
// 	_, err := Divide(10.0, 1.0)

// 	if err != nil {
// 		t.Error("Got an error when we should not have!")
// 	}
// }

// func TestBadDivide(t *testing.T) {
// 	_, err := Divide(10.0, 0)

// 	if err == nil {
// 		t.Error("Did NOT an error when we should not have!")
// 	}
// }
