package redis

import "fmt"

func GetAll(keysAndValuesMap *map[string]string) string {
	for currentKey, currentValue := range *keysAndValuesMap {
		fmt.Printf("[%v]: %v\n", currentKey, currentValue)
	}

	return "ok"
}
