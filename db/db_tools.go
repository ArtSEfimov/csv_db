package db

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func addExtension(filename string) string {
	if strings.HasSuffix(filename, ".csv") {
		return filename
	}
	return fmt.Sprint(filename, ".csv")
}

func writeAllRecords(tableName string, data [][]string) error {
	tableName = addExtension(tableName)

	dbFile, openErr := os.Create(tableName)
	defer dbFile.Close()
	if openErr != nil {
		return fmt.Errorf("open error: %w", openErr)
	}

	writer := csv.NewWriter(dbFile)
	writer.Comma = '|'

	writeErr := writer.WriteAll(data)
	if writeErr != nil {
		return fmt.Errorf("write error: %w", writeErr)
	}

	return nil
}

func readAllRecords(tableName string) ([][]string, error) {
	tableName = addExtension(tableName)

	dbFile, openErr := os.Open(tableName)
	defer dbFile.Close()
	if openErr != nil {
		return nil, fmt.Errorf("open error: %w", openErr)
	}

	reader := csv.NewReader(dbFile)
	reader.Comma = '|'

	allRecords, readErr := reader.ReadAll()

	if readErr != nil {
		return nil, fmt.Errorf("reading error: %w", readErr)
	}
	return allRecords, nil
}
