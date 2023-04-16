package usecase

import "time"

type Filter int64

const (
	Newest Filter = iota
	Cheapest
	MostExpensive
	NameAsc
	NameDesc
)

type Product struct {
	ID          int64
	Name        string
	Price       int64
	Description string
	Quantity    int
	CreateAt    time.Time
}

type AddProduct struct {
	Name        string
	Price       int64
	Description string
	Quantity    int
}
