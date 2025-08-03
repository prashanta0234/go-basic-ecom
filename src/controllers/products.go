package controllers

import (
	"e-com/src/dto"
	"e-com/src/helper"
	"e-com/src/services"
	"encoding/json"
	"net/http"
)

func Products(w http.ResponseWriter, r *http.Request) {
	helper.HandleHeader(w)

	if r.Method == "POST" {
		var input dto.ProductsSchema
		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
			return
		}

		userID := r.Context().Value("userID").(string)

		data, err := services.CreateProductsService(input, userID)

		if err != nil {
			http.Error(w, "Something went wrong: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Product Created successfully!",
			"data":    data,
		})
	}

	if r.Method == "GET" {
		nameFilter := r.URL.Query().Get("name")

		products, err := services.GetProducts(nameFilter)
		if err != nil {
			http.Error(w, "Something went wrong: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Products fetched successfully!",
			"data":    products,
		})
	}
}
