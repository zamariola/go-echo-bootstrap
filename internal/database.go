package internal

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

var (
	//DB contains database reference to be used across the program
	db   *sql.DB
	conn string
	err  error
)

//DB returns a database reference
func DB() (*sql.DB, error) {

	host := viper.Get("db_host")
	port := viper.Get("db_port")
	user := viper.Get("db_user")
	pass := viper.Get("db_pass")
	name := viper.Get("db_name")

	if hasNil(host, port, user, pass, name) {
		err = fmt.Errorf("unable to create database connection, missing information")
		return nil, err
	}

	conn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	return db, db.Ping()
}

func hasNil(args ...interface{}) bool {
	for _, arg := range args {
		if arg == nil {
			return true
		}
	}
	return false
}
