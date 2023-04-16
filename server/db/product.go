package db

import "log"

func (database *Database) CreateProductTable() error {
	_, err := database.db.Exec(initProductTable)
	if err != nil {
		log.Println("Failed to init Product Table. Error: ", err.Error())
		return err
	}
	return nil
}

func (database *Database) InsertProduct(param AddProduct) error {
	_, err := database.db.Exec(insertQuery, param.Name, param.Price, param.Description, param.Quantity)
	if err != nil {
		log.Println("Failed to insert Product. Error: ", err.Error())
		return err
	}
	return nil
}

func (database *Database) ListProduct(sortBy Filter) ([]Product, error) {
	finalQuery := getQuery

	if sortBy == Newest {
		finalQuery += "ORDER BY created_at DESC"
	}

	if sortBy == Cheapest {
		finalQuery += "ORDER BY price ASC"
	}

	if sortBy == MostExpensive {
		finalQuery += "ORDER BY price DESC"
	}

	if sortBy == NameAsc {
		finalQuery += "ORDER BY name ASC"
	}

	if sortBy == NameDesc {
		finalQuery += "ORDER BY name DESC"
	}

	finalQuery += ";"

	rows, err := database.db.Query(finalQuery)
	if err != nil {
		log.Println("Failed to get Product. Error: ", err.Error())
		return nil, err
	}
	defer rows.Close()

	productList := []Product{}

	for rows.Next() {
		var product Product

		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Quantity, &product.CreateAt)
		if err != nil {
			log.Println("Failed to Scan product. Error: ", err.Error())
			return nil, err
		}

		productList = append(productList, product)

	}
	return productList, nil
}
