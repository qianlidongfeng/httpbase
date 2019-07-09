package main

import "database/sql"

type Database struct{
	db *sql.DB
}

func NewDatabase(d *sql.DB) *Database{
	return &Database{
		db:d,
	}
}

func (this *Database) GetUserPasswordByName(name string) string{
	var password string
	this.db.QueryRow("select password from users where `name`=?",name).Scan(&password)
	return password
}