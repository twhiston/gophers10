package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestAssetBox(t *testing.T) {
	handler := http.FileServer(getAssetBox())
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}
	expected := "Image by: Ashley McNamara"
	actual, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(actual), "Image by: Ashley McNamara") {
		t.Errorf("You should attribute the artist! '%s'\n", expected)
	}

}
