package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId")

	pId, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please provide a valid id", http.StatusBadRequest)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == pId {
			util.SendData(w, product, http.StatusOK)
			return
		}
	}

	util.SendData(w, "Data not found", http.StatusNotFound)

}
