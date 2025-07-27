package models

import (
	"database/sql"
	"time"
)

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
	ID             int       `json: "id"`
	Name           string    `json: "name"`
	Description    string    `json: "description"`
	InventoryLevel string    `json: "inventory_level"`
	Price          int       `json: "price"`
	CreatedAt      time.Time `json: "-"`
	UpdateAt       time.Time `json: "-"`
}
