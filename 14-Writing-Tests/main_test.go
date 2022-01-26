package main

import "testing"

// The "RIGHT" way to write tests. Collect Test Cases and iterate then and run your checks
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expected-5", 100.0, 5.0, 20.0, false},
}

func TestDivisionUsingTestStruct(t *testing.T) {
	for _, tt := range tests {
		ans, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Did not get error when it was expected")
			}
		} else {
			if err != nil {
				t.Error("Got an Error while it was not expected with error:", err.Error())
			}
		}
		if tt.expected != ans {
			t.Errorf("Error: expected %v, but got %v", tt.expected, ans)
		}
	}
}

// Manual Test - One each for the Scenarios
func TestDivide(t *testing.T) {
	_, err := divide(10, 1)
	if err != nil {
		t.Error("Got an Error when we should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Error("Did not get an error when we should have ")
	}
}
