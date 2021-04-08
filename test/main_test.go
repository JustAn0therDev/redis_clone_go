package main_test

import (
	"testing"
)

func GetValueShouldReturnError(t *testing.T) {
	testMap := make(map[string]string)
	testString := ""
	result, err := main.GetValue(&testMap, &testString)
	if err == nil {
		t.Errorf("Expected error. Got: %v", result)
	}
}