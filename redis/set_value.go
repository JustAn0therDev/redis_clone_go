package redis

func SetValue(keysAndValuesMap *map[string]string, key *string, value *string) string {
	lastLength := len(*keysAndValuesMap)
	(*keysAndValuesMap)[*key] = *value

	// assert the length is different after insertion
	if len(*keysAndValuesMap) == lastLength {
		panic("key value pair could not be added to map")
	}

	return "ok"
}
