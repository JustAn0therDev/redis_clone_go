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

func TestDeleteValueShouldReturnError(t *testing.T) {
	testString := ""
	testMap := make(map[string]string)
	result, err := DeleteValue(&testMap, &testString)
	if err == nil {
		t.Errorf("expected error, got result: %v", result)
	}
}
func TestDeleteValueShouldNotReturnError(t *testing.T) {
	testString := "key"
	testMap := map[string]string{"key": "value"}
	_, err := DeleteValue(&testMap, &testString)
	if err != nil {
		t.Errorf("expected error to be nil, got: %v", err)
	}
}

func TestExistsShouldReturnOK(t *testing.T) {
	testString := "key"
	testValue := "val"
	testMap := make(map[string]string)
	result := SetValue(&testMap, &testString, &testValue)

	if result != "ok" || len(testMap) != 1 {
		panic("expected SetValue not to fail during Exists test")
	}

	result = Exists(&testMap, &testString)

	if result == "nok" {
		t.Error("expected Exists to return correctly, got nok")
	}
}