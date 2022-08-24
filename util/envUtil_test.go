package util

import (
	"os"
	"testing"
)

func TestGetEnvVariable(t *testing.T) {
	os.Setenv("FOO", "BOO")
	val := GetEnvVariable("FOO")
	if val != "BOO" {
		t.Errorf("Expected FOO but got %s", val)
	}
}

func TestATLAS_URI(t *testing.T) {
	val := GetEnvVariable("DATABASE_URL")
	if val == "" {
		t.Error("Expected DATABASE_URL")
	}
}
func TestSERVER_PORT(t *testing.T) {
	val := GetEnvVariable("PORT")
	if val != ":5050" {
		t.Error("Expected PORT but got ", val)
	}
}
