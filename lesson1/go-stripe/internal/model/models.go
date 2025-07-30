package models

import (
	"context"
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
	InventoryLevel int       `json: "inventory_level"`
	Price          int       `json: "price"`
	CreatedAt      time.Time `json: "-"`
	UpdateAt       time.Time `json: "-"`
}

type Order struct {
	ID            int       `json: "id"`
	WidgetId      int       `json: "widget_id"`
	TransactionId int       `json: "transaction_id"`
	StatusId      int       `json: "status_id"`
	Quantity      int       `json: "quantity"`
	Amount        int       `json: "amount"`
	CreatedAt     time.Time `json: "-"`
	UpdateAt      time.Time `json: "-"`
}

type Status struct {
	ID        int       `json: "id"`
	Name      string    `json: "name"`
	CreatedAt time.Time `json: "-"`
	UpdateAt  time.Time `json: "-"`
}

type TransactionStatus struct {
	ID        int       `json: "id"`
	Name      string    `json: "name"`
	CreatedAt time.Time `json: "-"`
	UpdateAt  time.Time `json: "-"`
}

type Transaction struct {
	ID        int       `json: "id"`
	Amoount   int       `json: "amount"`
	Currency  string    `json: "amount"`
	CreatedAt time.Time `json: "-"`
	UpdateAt  time.Time `json: "-"`
}

func (m *DBModel) GetWidget(id int) (Widget, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var widget Widget

	row := m.DB.QueryRowContext(ctx, "select id, name from widgets where id = ?", id)
	err := row.Scan(&widget.ID, &widget.Name)
	if err != nil {
		return widget, err
	}
	return widget, nil
}
