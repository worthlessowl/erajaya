package cache

import "time"

type Product struct {
	ID          int64
	Name        string
	Price       int64
	Description string
	Quantity    int
	CreateAt    time.Time
}
