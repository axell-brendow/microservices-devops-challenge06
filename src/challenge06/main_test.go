package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	value := greeting("Golang")
	expected := "<b>Golang</b>"
	if value != expected {
		t.Errorf("greeting(\"Golang\") should be %s but got %s", expected, value)
	}
}
