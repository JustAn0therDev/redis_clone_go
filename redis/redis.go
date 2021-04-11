package redis

import (
	"errors"
	"fmt"
)

func InitRedis() {
	var keysAndValuesMap = make(map[string]string)
	var command string
	var key string
	var value string
	var stopExecution bool
	var err error
	var result string

	for !stopExecution {
		key = ""
		value = ""
		fmt.Scanln(&command)

		switch command {
		case "GET":
			fmt.Scanln(&key)
			value, err = GetValue(&keysAndValuesMap, &key)
			HandlePrintlnOfCommandResult(&value, err)
		case "SET":
			fmt.Scanln(&key)
			fmt.Scanln(&value)
			result = SetValue(&keysAndValuesMap, &key, &value)
			HandlePrintlnOfCommandResult(&result, nil)
		case "DELETE":
			fmt.Scanln(&key)
			result, err = DeleteValue(&keysAndValuesMap, &key)
			HandlePrintlnOfCommandResult(&result, err)
		case "GETALL":
			result = GetAll(&keysAndValuesMap)
			HandlePrintlnOfCommandResult(&result, nil)
		case "EXISTS":
			fmt.Scanln(&key)
			result = Exists(&keysAndValuesMap, &key)
			HandlePrintlnOfCommandResult(&result, nil)
		case "FLUSH":
			result = Flush(&keysAndValuesMap)
			HandlePrintlnOfCommandResult(&result, nil)
		case "QUIT":
			stopExecution = true
		default:
			fmt.Println("invalid command")
		}
	}
}

func GetValue(keysAndValuesMap *map[string]string, key *string) (string, error) {
	for currentKey, currentValue := range *keysAndValuesMap {
		if *key == currentKey {
			return currentValue, nil
		}
	}

	return "", errors.New("nok")
}

func SetValue(keysAndValuesMap *map[string]string, key *string, value *string) string {
	lastLength := len(*keysAndValuesMap)
	(*keysAndValuesMap)[*key] = *value

	// assert the length is different after insertion
	if len(*keysAndValuesMap) == lastLength {
		panic("key value pair could not added to map")
	}

	return "ok"
}

func DeleteValue(keysAndValuesMap *map[string]string, key *string) (string, error) {
	for currentKey := range *keysAndValuesMap {
		if *key == currentKey {
			delete(*keysAndValuesMap, *key)
			return "ok", nil
		}
	}

	return "", errors.New("nok")
}

func GetAll(keysAndValuesMap *map[string]string) string {
	for currentKey, currentValue := range *keysAndValuesMap {
		fmt.Printf("[%v]: %v\n", currentKey, currentValue)
	}

	return "ok"
}

func Exists(keysAndValuesMap *map[string]string, key *string) string {
	for currentKey := range *keysAndValuesMap {
		if *key == currentKey {
			return "ok"
		}
	}

	return "nok"
}

func Flush(keysAndValuesMap *map[string]string) string {
	originalLength := len(*keysAndValuesMap)
	*keysAndValuesMap = make(map[string]string)

	if originalLength == len(*keysAndValuesMap) {
		panic("the keys and values map could not be cleared")
	}

	return "ok"
}