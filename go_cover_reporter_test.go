package main

import "testing"

func TestDummyFunction(t *testing.T) {
	const one, two, three = 1, 2, 3
	if x := dummyFunction(one, two); x != three {
		t.Errorf("dummyFunction(%d, %d) = %d, want %d", one, two, x, three)
	}
}
