package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please provide a valid id", http.StatusBadRequest)
		return
	}

	product := database.Get(pId)
	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	util.SendData(w, product, http.StatusOK)

}
