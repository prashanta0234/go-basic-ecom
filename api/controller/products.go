package controllers

import (
	"e-com/internal"
	"e-com/internal/reponse"
	usecase "e-com/usecase"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func Products(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

	if r.Method == "POST" {
		var input internal.ProductsSchema
		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			reponse.Error(w, 400, "Invalid input: "+err.Error(), err)
			return
		}

		userID := r.Context().Value("userID").(string)

		data, err := usecase.CreateProductsService(input, userID)

		if err != nil {
			reponse.Error(w, 500, "Something went wrong: "+err.Error(), err)
			return
		}

		reponse.Success(w, 201, "Product Created successfully!", data)
	}

	if r.Method == "PUT" {
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

		if len(pathParts) < 2 || pathParts[1] == "" {
			reponse.Error(w, 400, "Product ID is required", errors.New("product ID is required"))
			return
		}

		productID := pathParts[1]
		userID := r.Context().Value("userID").(string)

		var input internal.ProductsSchema
		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			reponse.Error(w, 400, "Invalid input: "+err.Error(), err)
			return
		}

		updatedProduct, err := usecase.UpdateProduct(productID, input, userID)

		if err != nil {
			reponse.Error(w, 400, "Update failed: "+err.Error(), err)
			return
		}

		reponse.Success(w, 200, "Product updated successfully!", updatedProduct)
	}

	if r.Method == "DELETE" {
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

		if len(pathParts) < 2 || pathParts[1] == "" {
			reponse.Error(w, 400, "Product ID is required", errors.New("product ID is required"))
			return
		}

		productID := pathParts[1]
		userID := r.Context().Value("userID").(string)

		err := usecase.DeleteProduct(productID, userID)

		if err != nil {
			reponse.Error(w, 400, "Delete failed: "+err.Error(), err)
			return
		}

		reponse.Success(w, 200, "Product deleted successfully!", nil)
	}

	if r.Method == "GET" {
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

		if len(pathParts) > 1 && pathParts[1] != "" {
			productID := pathParts[1]
			product, err := usecase.GetProductByID(productID)

			if err != nil {
				reponse.Error(w, 404, "Product not found: "+err.Error(), err)
				return
			}

			reponse.Success(w, 200, "Product fetched successfully!", product)
			return
		}

		nameFilter := r.URL.Query().Get("name")

		products, err := usecase.GetProducts(nameFilter)
		if err != nil {
			reponse.Error(w, 500, "Something went wrong: "+err.Error(), err)
			return
		}

		reponse.Success(w, 200, "Products fetched successfully!", products)
	}
}
