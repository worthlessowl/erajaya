package main

import (
	"log"
	"net/http"

	"github.com/worthlessowl/erajaya/cache"
	"github.com/worthlessowl/erajaya/db"
	"github.com/worthlessowl/erajaya/server"
	"github.com/worthlessowl/erajaya/usecase"
)

func main() {
	redisCache := cache.InitRedis()

	postgreDB, err := db.InitDB()
	if err != nil {
		panic("failed to initialize DB")
	}

	cch := cache.New(redisCache)
	db := db.New(postgreDB)

	err = db.CreateProductTable()
	if err != nil {
		panic(err)
		panic("failed to create table on DB")
	}

	uc := usecase.New(&db, &cch)

	sv := server.New(&uc)

	// Register endpoints here
	http.HandleFunc("/", sv.HandleGetProduct)
	http.HandleFunc("/insert", sv.HandleInsertProduct)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
