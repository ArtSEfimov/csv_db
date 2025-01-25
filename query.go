package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func validateQuery(query string) (bool, string) {

	pythonScript := exec.Command("python", "main_validator.py", query)

	var stdout bytes.Buffer
	pythonScript.Stdout = &stdout

	execErr := pythonScript.Run()

	if execErr != nil {
		return false, execErr.Error()
	}

	fromPython := bytes.Split(stdout.Bytes(), []byte("\n"))

	boolResult, parseErr := strconv.ParseBool(strings.TrimSpace(string(fromPython[0])))
	if parseErr != nil {
		return false, parseErr.Error()
	}
	errorString := string(fromPython[1])

	return boolResult, errorString

}

func parseQuery(query string) (string, string, []string) {
	// Разбирает запрос на части: тип операции, имя таблицы, аргументы.
	// Возвращает тип операции (CREATE_TABLE, SELECT, INSERT и т.д.), имя таблицы и список аргументов.

	splitQuery := strings.Split(query, " ")
	requestType := splitQuery[0]
	tableName := splitQuery[1]
	arguments := splitQuery[2:]
	for i := range arguments {
		arguments[i] = strings.Trim(arguments[i], ",")
	}
	return requestType, tableName, arguments

}

func handleQuery(requestType string, table string, arguments []string) (bool, string) {
	// Выполняет запрос, переданный пользователем.
	// В зависимости от типа операции вызывает соответствующую функцию из db.go.
	// Возвращает true и результат операции, либо false и сообщение об ошибке.
	switch strings.ToUpper(requestType) {
	case "CREATE_TABLE":
		ok := createTable(table, arguments)
		if ok {
			return true, fmt.Sprintf("Table %s created successfully", table)
		}
		return false, fmt.Sprint("Creat table failed")
	case "SELECT":
		argument := arguments[0]
		if argument == "*" {
			result, ok := selectAll(table)
			if ok {
				return true, result
			}
			return false, fmt.Sprint("Select all records failed")
		}

		result, ok := selectRecord(table, argument)
		if ok {
			return true, result
		}
		return false, fmt.Sprintf("Select record with ID %s failed", argument)

	case "INSERT":
		ok := insertRecord(table, arguments)
		if ok {
			return true, fmt.Sprintf("New record inserted successfully")
		}
		return false, fmt.Sprint("Insert record failed")

	case "UPDATE":
		id := arguments[0]
		arguments = arguments[1:]
		ok := updateRecord(table, id, arguments)
		if ok {
			return true, fmt.Sprintf("Update record with ID %s successfully", id)
		}
		return false, fmt.Sprint("Update record failed")

	case "DELETE":
		id := arguments[0]
		ok := deleteRecord(table, id)
		if ok {
			return true, fmt.Sprintf("Delete record with ID %s successfully", id)
		}
		return false, fmt.Sprint("Delete record failed")
	}

	return false, ""
}
