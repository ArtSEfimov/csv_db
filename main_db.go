package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var userQuery string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		userQuery = scanner.Text()
		ok, validateErr := validateQuery(userQuery)
		if !ok {
			fmt.Println(validateErr)
			return
		}

		ok, result := handleQuery(parseQuery(userQuery))
		if !ok {
			fmt.Println(fmt.Sprintf("Error: %s", result))
			return
		}

		fmt.Println(result)

	}
}
