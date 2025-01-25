package main

import (
	"bufio"
	"fmt"
	"github.com/ArtSEfimov/csv_db/db"
	"os"
)

func main() {

	var userQuery string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		userQuery = scanner.Text()
		ok, validateErr := db.ValidateQuery(userQuery)
		if !ok {
			fmt.Println(validateErr)
			return
		}

		result, ok := db.HandleQuery(db.ParseQuery(userQuery))
		if !ok {
			fmt.Println(fmt.Sprintf("Error: %s", result))
			return
		}

		fmt.Println(result)

	}
}
