package redis

func Flush(keysAndValuesMap *map[string]string) string {
	originalLength := len(*keysAndValuesMap)
	*keysAndValuesMap = make(map[string]string)

	if originalLength == len(*keysAndValuesMap) {
		panic("the keys and values map could not be cleared")
	}

	return "ok"
}
