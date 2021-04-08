package redis

import (
	"testing"
)

func TestGetValueShouldNotReturnError(t *testing.T) {
	testString := "key"
	testMap := map[string]string{"key": "value"}
	result, err := GetValue(&testMap, &testString)
	if err != nil {
		t.Errorf("Expected return value %v. Got error: %v", result, err)
	}
}

func TestGetValueShouldReturnError(t *testing.T) {
	testString := ""
	testMap := make(map[string]string)
	result, err := GetValue(&testMap, &testString)
	if err == nil {
		t.Errorf("Expected error. Got: %v", result)
	}
}