package controllers

import (
	"e-com/internal"
	"e-com/internal/reponse"
	"e-com/usecase"
	"errors"
	"net/http"
)

func GetUserOrdersController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

	if r.Method != "GET" {
		reponse.Error(w, 404, "Method not allowed", errors.New("method not allowed"))
		return
	}

	userID := r.Context().Value("userID")
	if userID == nil {
		reponse.Error(w, 401, "User not authenticated", errors.New("user not authenticated"))
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		reponse.Error(w, 500, "Invalid user ID", errors.New("invalid user ID"))
		return
	}

	orders, err := usecase.GetOrdersByUserID(userIDStr)
	if err != nil {
		reponse.Error(w, 500, "Failed to get orders: "+err.Error(), err)
		return
	}

	reponse.Success(w, 200, "Orders retrieved successfully", orders)
}

func GetOrderByIDController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

	if r.Method != "GET" {
		reponse.Error(w, 404, "Method not allowed", errors.New("method not allowed"))
		return
	}

	orderID := r.URL.Path[len("/order/"):]
	if orderID == "" {
		reponse.Error(w, 400, "Order ID is required", errors.New("order ID is required"))
		return
	}

	userID := r.Context().Value("userID")
	if userID == nil {
		reponse.Error(w, 401, "User not authenticated", errors.New("user not authenticated"))
		return
	}

	order, err := usecase.GetOrderByID(orderID)
	if err != nil {
		reponse.Error(w, 404, "Failed to get order: "+err.Error(), err)
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		reponse.Error(w, 500, "Invalid user ID", errors.New("invalid user ID"))
		return
	}

	if order.UserID.Hex() != userIDStr {
		reponse.Error(w, 403, "Access denied", errors.New("access denied"))
		return
	}

	reponse.Success(w, 200, "Order retrieved successfully", order)
}
