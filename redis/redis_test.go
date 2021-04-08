package main

import (
	"testing"
)

func GetValueShouldReturnError(t *testing.T) {
	testMap := make(map[string]string)
	testString := ""
	result, err := GetValue(&testMap, &testString)
	if err == nil {
		t.Errorf("Expected error. Got: %v", result)
	}
}