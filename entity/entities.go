package entity

import "time"

type User struct {
	ID        int64
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID         int64
	Name       string
	Price      float64
	CategoryID int64
}

type Category struct {
	ID   int64
	Name string
}
