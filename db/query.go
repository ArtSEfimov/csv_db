package db

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

const (
	CREATE = "CREATE_TABLE"
	SELECT = "SELECT"
	INSERT = "INSERT"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
)

func ValidateQuery(query string) (bool, string) {

	pythonScript := exec.Command("python", "./validators/main_validator.py", query)

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

func ParseQuery(query string) (string, string, []string) {
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

func HandleQuery(requestType string, table string, arguments []string) (string, bool) {
	// Выполняет запрос, переданный пользователем.
	// В зависимости от типа операции вызывает соответствующую функцию из db.go.
	// Возвращает true и результат операции, либо false и сообщение об ошибке.
	switch strings.ToUpper(requestType) {
	case CREATE:
		ok := createTable(table, arguments)
		if ok {
			return fmt.Sprintf("Table %s created successfully", table), true
		}
		return fmt.Sprint("Creat table failed"), false
	case SELECT:
		argument := arguments[0]
		if argument == "*" {
			result, ok := selectAll(table)
			if ok {
				return result, true
			}
			return fmt.Sprint("Select all records failed"), false
		}

		result, ok := selectRecord(table, argument)
		if ok {
			return result, true
		}
		return fmt.Sprintf("Select record with ID %s failed", argument), false

	case INSERT:
		ok := insertRecord(table, arguments)
		if ok {
			return fmt.Sprintf("New record inserted successfully"), true
		}
		return fmt.Sprint("Insert record failed"), false

	case UPDATE:
		id := arguments[0]
		arguments = arguments[1:]
		ok := updateRecord(table, id, arguments)
		if ok {
			return fmt.Sprintf("Update record with ID %s successfully", id), true
		}
		return fmt.Sprint("Update record failed"), false

	case DELETE:
		id := arguments[0]
		ok := deleteRecord(table, id)
		if ok {
			return fmt.Sprintf("Delete record with ID %s successfully", id), true
		}
		return fmt.Sprint("Delete record failed"), false
	}

	return "", false
}
