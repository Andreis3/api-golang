package main

import (
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/andreis3/api-golang/configs"
	"github.com/andreis3/api-golang/internal/entity"
	"github.com/andreis3/api-golang/internal/infra/database"
	handlers "github.com/andreis3/api-golang/internal/infra/webserver"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)

	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8080", nil)
}
