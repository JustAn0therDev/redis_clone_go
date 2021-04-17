package redis

import "errors"

func DeleteValue(keysAndValuesMap *map[string]string, key *string) (string, error) {
	for currentKey := range *keysAndValuesMap {
		if *key == currentKey {
			delete(*keysAndValuesMap, *key)
			return "ok", nil
		}
	}

	return "", errors.New("nok")
}
