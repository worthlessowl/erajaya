package db

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
	ID          int64     `field:"id"`
	Name        string    `field:"name"`
	Price       int64     `field:"price"`
	Description string    `field:"description"`
	Quantity    int       `field:"quantity"`
	CreateAt    time.Time `field:"create_at"`
}

type AddProduct struct {
	Name        string
	Price       int64
	Description string
	Quantity    int
}
