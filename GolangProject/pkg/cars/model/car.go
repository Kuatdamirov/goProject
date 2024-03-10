package model

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Car struct {
	Id             string `json:"id"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Year uint   `json:"Year"`
}

type carModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (c carModel) Insert(car *Car) error {
	// Insert a new menu item into the database.
	query := `
		INSERT INTO car (title, description, year) 
		VALUES ($1, $2, $3) 
		RETURNING id, created_at, updated_at
		`
	args := []interface{}{car.Title, car.Description, car.Year}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.DB.QueryRowContext(ctx, query, args...).Scan(&car.Id, &car.CreatedAt, &car.UpdatedAt)
}

func (c carModel) Get(id int) (*Car, error) {
	// Retrieve a specific menu item based on its ID.
	query := `
		SELECT id, created_at, updated_at, title, description, year
		FROM car
		WHERE id = $1
		`
	var car Car
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := c.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&car.Id, &car.CreatedAt, &car.UpdatedAt, &car.Title, &car.Description, &car.Year)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (c carModel) Update(car *Car) error {
	// Update a specific menu item in the database.
	query := `
		UPDATE car
		SET title = $1, description = $2, year = $3
		WHERE id = $4
		RETURNING updated_at
		`
	args := []interface{}{car.Title, car.Description, car.Year, car.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.DB.QueryRowContext(ctx, query, args...).Scan(&car.UpdatedAt)
}

func (c carModel) Delete(id int) error {
	// Delete a specific menu item from the database.
	query := `
		DELETE FROM car
		WHERE id = $1
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := c.DB.ExecContext(ctx, query, id)
	return err
}