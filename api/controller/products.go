package controllers

import (
	"e-com/internal"
	"e-com/internal/reponse"
	usecase "e-com/usecase"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// GetProducts godoc
// @Summary Get all products
// @Description Retrieve all products with optional filtering and pagination
// @Tags Products
// @Accept json
// @Produce json
// @Param name query string false "Filter by product name"
// @Param page query int false "Page number (default: 1)"
// @Param skip query int false "Number of items to skip (default: 0)"
// @Param limit query int false "Number of items per page (default: 10)"
// @Success 200 {object} reponse.SuccessResponse "Products fetched successfully"
// @Failure 400 {object} reponse.ErrorResponse "Bad request - Invalid parameters"
// @Failure 500 {object} reponse.ErrorResponse "Internal server error"
// @Router /products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {

	nameFilter := r.URL.Query().Get("name")
	page := r.URL.Query().Get("page")
	skip := r.URL.Query().Get("skip")
	limit := r.URL.Query().Get("limit")

	if page == "" {
		page = "1"
	}
	if skip == "" {
		skip = "0"
	}
	if limit == "" {
		limit = "10"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		reponse.Error(w, 400, "Invalid page: "+err.Error(), err)
		return
	}
	skipInt, err := strconv.Atoi(skip)
	if err != nil {
		reponse.Error(w, 400, "Invalid skip: "+err.Error(), err)
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		reponse.Error(w, 400, "Invalid limit: "+err.Error(), err)
		return
	}

	productsResponse, err := usecase.GetProducts(usecase.Filter{
		Name:  nameFilter,
		Page:  pageInt,
		Skip:  skipInt,
		Limit: limitInt,
	})
	if err != nil {
		reponse.Error(w, 500, "Something went wrong: "+err.Error(), err)
		return
	}

	reponse.Success(w, 200, "Products fetched successfully!", productsResponse)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product (requires authentication)
// @Tags Products
// @Accept json
// @Produce json
// @Param product body internal.ProductsSchema true "Product information"
// @Success 201 {object} reponse.SuccessResponse "Product created successfully"
// @Failure 400 {object} reponse.ErrorResponse "Bad request - Invalid input data"
// @Failure 401 {object} reponse.ErrorResponse "Unauthorized - Authentication required"
// @Failure 500 {object} reponse.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /products [post]
func CreateProduct(w http.ResponseWriter, r *http.Request) {
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

// UpdateProduct godoc
// @Summary Update a product
// @Description Update an existing product (requires authentication)
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body internal.ProductsSchema true "Updated product information"
// @Success 200 {object} reponse.SuccessResponse "Product updated successfully"
// @Failure 400 {object} reponse.ErrorResponse "Bad request - Invalid input data or product ID"
// @Failure 401 {object} reponse.ErrorResponse "Unauthorized - Authentication required"
// @Failure 500 {object} reponse.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /products/{id} [put]
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
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

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product (requires authentication)
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} reponse.SuccessResponse "Product deleted successfully"
// @Failure 400 {object} reponse.ErrorResponse "Bad request - Product ID required"
// @Failure 401 {object} reponse.ErrorResponse "Unauthorized - Authentication required"
// @Failure 500 {object} reponse.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /products/{id} [delete]
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
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

// GetProductByID godoc
// @Summary Get product by ID
// @Description Retrieve a specific product by its ID (requires authentication)
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} reponse.SuccessResponse "Product fetched successfully"
// @Failure 400 {object} reponse.ErrorResponse "Bad request - Product ID required"
// @Failure 401 {object} reponse.ErrorResponse "Unauthorized - Authentication required"
// @Failure 404 {object} reponse.ErrorResponse "Product not found"
// @Security BearerAuth
// @Router /products/{id} [get]
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	productID := pathParts[1]
	product, err := usecase.GetProductByID(productID)

	if err != nil {
		reponse.Error(w, 404, "Product not found: "+err.Error(), err)
		return
	}

	reponse.Success(w, 200, "Product fetched successfully!", product)
}
