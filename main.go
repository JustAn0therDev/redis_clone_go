package main

import "fmt"

// initially, the REPL will have only three commands: GET, SET and DELETE

// work with bytes maybe?
var keysAndValuesMap = make(map[string]string)
var command string
var key string
var value string

func main() {

	// TODO: use function pointers so I do not have to have a switch case statement in the middle of the function
	for {
		key = ""
		value = ""
		fmt.Scanln(&command)

		switch command {
		case "GET":
			fmt.Scanln(&key)
			GetValue(&key)
		case "SET":
			fmt.Scanln(&key)
			fmt.Scanln(&value)
			SetValue(key, value)
		}
	}
}

func GetValue(key *string) string {
	for currentKey, currentValue := range keysAndValuesMap {
		if *key == currentKey {
			return currentValue
		}
	}

	return ""
}

// TODO: Test pass by reference
func SetValue(key string, value string) {
	keysAndValuesMap[key] = value
}