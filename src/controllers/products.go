package controllers

import (
	"e-com/src/dto"
	"e-com/src/helper"
	"e-com/src/services"
	"encoding/json"
	"net/http"
	"strings"
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

	if r.Method == "PUT" {
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		
		if len(pathParts) < 2 || pathParts[1] == "" {
			http.Error(w, "Product ID is required", http.StatusBadRequest)
			return
		}

		productID := pathParts[1]
		userID := r.Context().Value("userID").(string)

		var input dto.ProductsSchema
		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
			return
		}

		updatedProduct, err := services.UpdateProduct(productID, input, userID)

		if err != nil {
			http.Error(w, "Update failed: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Product updated successfully!",
			"data":    updatedProduct,
		})
	}

	if r.Method == "DELETE" {
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		
		if len(pathParts) < 2 || pathParts[1] == "" {
			http.Error(w, "Product ID is required", http.StatusBadRequest)
			return
		}

		productID := pathParts[1]
		userID := r.Context().Value("userID").(string)

		err := services.DeleteProduct(productID, userID)

		if err != nil {
			http.Error(w, "Delete failed: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Product deleted successfully!",
		})
	}

	if r.Method == "GET" {
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		
		if len(pathParts) > 1 && pathParts[1] != "" {
			productID := pathParts[1]
			product, err := services.GetProductByID(productID)
			
			if err != nil {
				http.Error(w, "Product not found: "+err.Error(), http.StatusNotFound)
				return
			}

			w.WriteHeader(200)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Product fetched successfully!",
				"data":    product,
			})
			return
		}

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
