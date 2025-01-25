package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var IDStorage map[string]uint = make(map[string]uint)

func getNextID(tableName string) (uint, bool) {
	/* Вычисляет следующий ID для новой записи на основе текущих записей таблицы.
	Возвращает следующий ID и true, либо 0 и false в случае ошибки.*/

	IDStorage[tableName]++

	idValue, ok := IDStorage[tableName]
	if !ok {
		return 0, false
	}

	return idValue, true

}

func createTable(tableName string, fieldNames []string) bool {
	/*Создаёт новую таблицу (файл .csv) с заголовками (колонками).
	  Возвращает true, если операция успешна, иначе false.*/
	IDStorage[tableName] = 0

	tableName = checkExtension(tableName)

	dbFile, createErr := os.Create(tableName)
	defer dbFile.Close()
	if createErr != nil {
		return false
	}

	newWriter := csv.NewWriter(dbFile)
	newWriter.Comma = '|'

	fieldNames = append([]string{"ID"}, fieldNames...)

	writeErr := newWriter.Write(fieldNames)
	if writeErr != nil {
		return false
	}
	defer newWriter.Flush()

	return true
}

func insertRecord(tableName string, fieldValues []string) bool {
	// Добавляет новую запись в таблицу. Генерирует уникальный ID автоматически.
	// Возвращает true, если операция успешна, иначе false.

	newID, ok := getNextID(tableName)
	if !ok {
		return false
	}

	fieldValues = append([]string{fmt.Sprintf("%d", newID)}, fieldValues...)

	tableName = checkExtension(tableName)

	dbFile, openErr := os.OpenFile(tableName, os.O_APPEND, 0744)
	defer dbFile.Close()
	if openErr != nil {
		fmt.Println(openErr)
		return false
	}

	writer := csv.NewWriter(dbFile)
	writer.Comma = '|'

	writeErr := writer.Write(fieldValues)
	if writeErr != nil {
		fmt.Println(writeErr)
		return false
	}

	defer writer.Flush()

	return true
}

func updateRecord(tableName string, id string, fieldValues []string) bool {
	// Обновляет существующую запись в таблице по её ID.
	// Возвращает true, если операция успешна, иначе false.

	allRecords, readErr := readAllRecords(tableName)
	if readErr != nil {
		return false
	}

	for i, record := range *allRecords {
		if record[0] == id {
			(*allRecords)[i] = append([]string{id}, fieldValues...)
			writeErr := writeAllRecords(tableName, allRecords)
			if writeErr != nil {
				return false
			}
			return true
		}
	}

	return false
}

func selectAll(tableName string) (string, bool) {
	// Возвращает все записи таблицы в виде строки.
	// Возвращает строку и true, либо пустую строку и false в случае ошибки.

	allRecords, readErr := readAllRecords(tableName)
	if readErr != nil {
		return "", false
	}

	allRecordsBuilder := strings.Builder{}
	for _, record := range *allRecords {
		allRecordsBuilder.WriteString(strings.Join(record, " "))
		allRecordsBuilder.WriteRune('\n')
	}
	return allRecordsBuilder.String(), true
}

func selectRecord(tableName string, id string) (string, bool) {
	// Возвращает запись таблицы по её ID.
	// Возвращает строку с записью и true, либо пустую строку и false в случае ошибки.
	tableName = checkExtension(tableName)

	dbFile, openErr := os.Open(tableName)
	defer dbFile.Close()
	if openErr != nil {
		return "", false
	}

	reader := csv.NewReader(dbFile)
	reader.Comma = '|'

	for {
		currentRecord, readErr := reader.Read()
		if readErr != nil {
			if errors.Is(readErr, io.EOF) {
				break
			}
			return "", false
		}
		if currentRecord[0] == id {
			return strings.Join(currentRecord, " "), true
		}
	}
	return "", false
}

func deleteRecord(tableName string, id string) bool {
	// Удаляет запись из таблицы по её ID.
	// Возвращает true, если операция успешна, иначе false.
	allRecords, readErr := readAllRecords(tableName)

	if readErr != nil {
		return false
	}

	for i, record := range *allRecords {
		if record[0] == id {
			*allRecords = append((*allRecords)[:i], (*allRecords)[i+1:]...)
			writeErr := writeAllRecords(tableName, allRecords)
			if writeErr != nil {
				return false
			}
			return true
		}
	}
	return false
}
