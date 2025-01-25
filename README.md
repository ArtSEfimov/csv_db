# CSV Database

A simple and efficient database implemented using CSV files in Go.

---

## Features
- Read and write data to CSV files.
- Perform basic CRUD (Create, Read, Update, Delete) operations.
- Easy integration into Go projects.
- Lightweight and dependency-free.

---

## Requirements
- Go 1.22 or later.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ArtSEfimov/csv_db.git
   ```

2. Navigate to the project directory:
   ```bash
   cd csv_db
   ```

3. Install dependencies (if any):
   ```bash
   go mod tidy
   ```

---

## Usage

### 1. Initializing the Database
To use the CSV database, initialize it with the path to your CSV file:

```go
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
```

### 2. CRUD Operations

#### Create
Add a new record to the database:
```go
		ok := createTable(table, arguments)
		if ok {
			return fmt.Sprintf("Table %s created successfully", table), true
		}
		return fmt.Sprint("Creat table failed"), false
```

#### Read
Fetch all records or specific ones:
```go
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
```

#### Update
Update an existing record by its ID:
```go
		id := arguments[0]
		arguments = arguments[1:]
		ok := updateRecord(table, id, arguments)
		if ok {
			return fmt.Sprintf("Update record with ID %s successfully", id), true
		}
		return fmt.Sprint("Update record failed"), false
```

#### Delete
Delete a record by its ID:
```go
		id := arguments[0]
		ok := deleteRecord(table, id)
		if ok {
			return fmt.Sprintf("Delete record with ID %s successfully", id), true
		}
		return fmt.Sprint("Delete record failed"), false
```

## File Format
The CSV file should follow this structure:

| ID  | Name      | Age | Occupation      |
|-----|-----------|-----|-----------------|
| 1   | John Doe  | 30  | Engineer        |
| 2   | Jane Smith| 25  | Data Scientist  |

### Notes:
- The first row is considered the header.
- Each subsequent row represents a record.
- Ensure unique IDs for each record.

---

## Contributing
Feel free to submit issues or pull requests. Contributions are welcome!

---

## Contact
For questions or feedback, contact [ArtSEfimov](https://github.com/ArtSEfimov).

