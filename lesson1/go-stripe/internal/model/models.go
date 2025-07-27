package models

import "database/sql"

type DBModel struct {
	DB *sql.DB
}

type Models struct {
	DB DBModel
}
