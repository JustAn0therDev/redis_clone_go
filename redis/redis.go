package redis

import "fmt"

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
