package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

// Get All Products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.ProductList, 200)
}
