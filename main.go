package main

import "databaseTask/storage/database"

func main() {
	println(database.Client.Error)
	return
}
