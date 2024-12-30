package getenv

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {

	err := LoadEnv(".env")
	if err != nil {
		t.Error(err)
	}

	firstValue := os.Getenv("FIRST_VALUE")
	if firstValue != "10101001" {
		t.Error("cannot assert first value")
	}

	secondValue := os.Getenv("SECOND_VALUE")
	if secondValue != "0n3hndr3ds3xtYSX" {
		t.Error("cannot assert second value")
	}

}
