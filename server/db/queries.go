package db

const (
	insertQuery = "INSERT INTO product(name, price, description, quantity) VALUES ($1, $2, $3, $4);"

	getQuery = "SELECT * FROM product "

	initProductTable = ` CREATE TABLE IF NOT EXISTS product (
		id bigint GENERATED ALWAYS AS IDENTITY,
		name varchar NOT NULL,
		price bigint NOT NULL,
		description varchar(1000) NOT NULL,
		quantity int NOT NULL,
		created_at timestamp default now()
	);
	`
)
