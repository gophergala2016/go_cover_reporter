package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

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
		t.Errorf("toText(%d, %.2f) = %s, expected %s", frameNumber, percentage, generatedString, expectedString)
	}
}

func TestToTextFunction333(t *testing.T) {
	const (
		frameNumber    = 40
		percentage     = 33.3
		expectedString = "cover:33.30%"
	)
	if generatedString := toText(frameNumber, percentage); generatedString != expectedString {
		t.Errorf("toText(%d, %.2f) = %s, expected %s", frameNumber, percentage, generatedString, expectedString)
	}
}

func TestCoverBadgeFunction(t *testing.T) {

	file, err := ioutil.TempFile(os.TempDir(), "tmp_test_dir")
	if err != nil {
		t.Error("Unable to create temp directory.")
	}
	defer os.Remove(file.Name())

	coverBadge(io.Writer(file), 12.3)

	fileinfo, err := file.Stat()
	if err != nil {
		t.Error("No gif image file was created.")
	}

	if fileinfo.Size() < 20000 {
		t.Error("File is significantly smaller than expected.")
	}

}
