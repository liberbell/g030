package models

import "database/sql"

type DBModel struct {
	DB *sql.DB
}

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Widget struct {
	ID          int    `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
}
