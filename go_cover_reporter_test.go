package main

import "testing"

func TestDummyFunction(t *testing.T) {
	const one, two, three = 1, 2, 3
	if x := dummyFunction(one, two); x != three {
		t.Errorf("dummyFunction(%d, %d) = %d, want %d", one, two, x, three)
	}
}

func TestToTextFunction0(t *testing.T) {
	const (
		frameNumber    = 40
		percentage     = 0.0
		expectedString = "cover:0.00%"
	)
	if generatedString := toText(frameNumber, percentage); generatedString != expectedString {
		t.Errorf("toText(%d, %.2f) = %s, want %s", frameNumber, percentage, generatedString, expectedString)
	}
}

func TestToTextFunction333(t *testing.T) {
	const (
		frameNumber    = 40
		percentage     = 33.3
		expectedString = "cover:33.30%"
	)
	if generatedString := toText(frameNumber, percentage); generatedString != expectedString {
		t.Errorf("toText(%d, %.2f) = %s, want %s", frameNumber, percentage, generatedString, expectedString)
	}
}
