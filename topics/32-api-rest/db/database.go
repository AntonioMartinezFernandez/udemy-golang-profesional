package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(host:port)/database?charset=utf8
const url = "root:password@tcp(localhost:3306)/go_apirest_db"

// Save connection
var db *sql.DB

// Establish connection
func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("MySQL connected")
	db = conection
}

// Close connection
func Close() {
	db.Close()
	fmt.Println("MySQL disconnected")
}

// Verify connection
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// Verify if table exists
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()
}

// Create table
func CreateTable(schema, name string) {

	if !ExistsTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Delete table
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	db.Exec(sql)
}

// Polymorphism to Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// Polymorphism to Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
