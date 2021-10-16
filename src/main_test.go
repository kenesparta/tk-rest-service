package main

import (
	"fmt"
	"os"
	"testing"
)

//
const coverageAllowed = 0.0

func TestMain(m *testing.M) {
	rc := m.Run()

	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < coverageAllowed {
			fmt.Printf("\n---\nTest Coverage Error \nAllowed: %.2f / Current: %.2f\n---\n", coverageAllowed, c)
			rc = -1
		}
	}
	os.Exit(rc)
}