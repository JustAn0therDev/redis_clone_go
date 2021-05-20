package redis

import "fmt"

func HandlePrintlnOfCommandResult(result *string, err error) {
	if err != nil {
		fmt.Println(err)
	} else if *result != "" {
		fmt.Println(*result)
	}
}
