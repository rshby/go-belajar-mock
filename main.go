package main

import "go-belajar-mock/internal/database"

func main() {
	dbMysql, dbMysqlCloser := database.InitializeMySQLDatabase()
	defer dbMysqlCloser()

	// repository
}
