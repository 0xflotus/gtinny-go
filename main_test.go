package main

import (
	"testing"
)

func TestValidGTIN(t *testing.T) {
	var gtin string = "97350053850012"
	if isValid(&gtin) != true {
		t.Fatalf(`%s should be a valid GTIN`, gtin)
	}
}

func TestInvalidGTIN(t *testing.T) {
	var gtin string = "1234"
	if isValid(&gtin) != false {
		t.Fatalf(`%s should be an invalid GTIN`, gtin)
	}
}
