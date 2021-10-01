package main

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

func createErrNoRows(sqlStr string) error {
	return sql.ErrNoRows
}

func doSomething(sqlStr string) error {
	err := createErrNoRows(sqlStr)

	if err == sql.ErrNoRows {
		return errors.Wrap(err, "Error With SQL:"+sqlStr)
	}

	return nil
}

func main() {
	err := doSomething("select id, name from users where id = 1")
	if err != nil {
		log.Printf("%+v", err)
	}
}
