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
	val := GetEnvVariable("ATLAS_URI")
	if val == "" {
		t.Error("Expected ATLAS_URI")
	}
}
func TestSERVER_PORT(t *testing.T) {
	val := GetEnvVariable("SERVER_PORT")
	if val != ":9090" {
		t.Error("Expected SERVER_PORT but got ", val)
	}
}
func TestBASE_PATH(t *testing.T) {
	val := GetEnvVariable("BASE_PATH")
	if val != "api/v1" {
		t.Error("Expected BASE_PATH but got ", val)
	}
}
