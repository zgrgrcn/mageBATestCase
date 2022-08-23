package db

import (
	"testing"
)

func TestClientNeedToBeEmptyAtTheBeginning(t *testing.T) {
	if client != nil {
		t.Errorf("Connection needs to be init first")
	}
}

func TestGettingMongoConnection(t *testing.T) {
	client, ctx := GetConnection()
	if client == nil {
		t.Errorf("client needs a value after init")
	}
	if ctx == nil {
		t.Errorf("ctx needs a value after init")
	}
}
