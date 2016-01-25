package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

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

func TestLandingPageHttpHandler(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(`Test server responded with error on the GET request for handler`)
	}

	wholepage, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(`Reading landing page generated from template was unsuccessful`)
	}

	if !strings.Contains(string(wholepage), "Most recent reported code coverage is") {
		t.Error(`Landing page generated from template doesn't contain expected text`)
	}

}

func TestGifBadgeHttpHandler(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(badge))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(`Test server responded with error on the GET request for badge handler`)
	}

	if res.Status != "200 OK" {
		t.Errorf(`Test server responded: %s instead of 200 OK`, res.Status)
	}
}

func TestReceiverHttpHandler(t *testing.T) {

	Filename = "forthistestonly.txt"

	file, err := os.Create(Filename)
	if err != nil {
		t.Error("System was unable to create file named forthistestonly.txt.")
	}
	defer os.Remove(file.Name())

	_, err = file.Stat()
	if err != nil {
		t.Error("File named forthistestonly.txt was not created.")
	}

	ts := httptest.NewServer(http.HandlerFunc(receiver))
	defer ts.Close()

	jsonStr := []byte(`{"Name":"GoCoverReporter","Body":"x 12.3% x"}`)

	req, err := http.NewRequest("POST", ts.URL, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "GoCoverReporter")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(`Test server responded with error on the POST request with JSON`)
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		t.Errorf(`Test server responded: %s instead of 200 OK`, resp.Status)
	}

	if readPercentageFromFile() != 12.30 {
		t.Errorf(`Value read from the test file was: %s instead of 12.30`, readPercentageFromFile())
	}

}
