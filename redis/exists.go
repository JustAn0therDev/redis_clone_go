package redis

func Exists(keysAndValuesMap *map[string]string, key *string) string {
	for currentKey := range *keysAndValuesMap {
		if *key == currentKey {
			return "ok"
		}
	}

	return "nok"
}
