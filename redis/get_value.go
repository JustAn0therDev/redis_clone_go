package redis

import "errors"

func GetValue(keysAndValuesMap *map[string]string, key *string) (string, error) {
	for currentKey, currentValue := range *keysAndValuesMap {
		if *key == currentKey {
			return currentValue, nil
		}
	}

	return "", errors.New("nok")
}
