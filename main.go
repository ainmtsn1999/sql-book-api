package main

import (
	"github.com/ainmtsn1999/sql-book-api/app"
	"github.com/ainmtsn1999/sql-book-api/config"
	"github.com/joho/godotenv"
)

// create db_go_sql;
// create schema db_go_sql;
// CREATE TABLE db_go_sql.books (
// 	id serial primary key,
// 	title varchar(200) not null,
// 	author varchar(50) not null,
// 	"desc" text not null
// );

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = config.InitPostgres()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
