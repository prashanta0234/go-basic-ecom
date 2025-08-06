package controllers

import (
	"e-com/internal"
	"e-com/usecase"
	"encoding/json"
	"net/http"
)

func GetUserOrdersController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.Context().Value("userID")
	if userID == nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	orders, err := usecase.GetOrdersByUserID(userIDStr)
	if err != nil {
		http.Error(w, "Failed to get orders: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Orders retrieved successfully",
		"orders":  orders,
	})
}

func GetOrderByIDController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	orderID := r.URL.Path[len("/order/"):]
	if orderID == "" {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID")
	if userID == nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	order, err := usecase.GetOrderByID(orderID)
	if err != nil {
		http.Error(w, "Failed to get order: "+err.Error(), http.StatusNotFound)
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	if order.UserID.Hex() != userIDStr {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Order retrieved successfully",
		"order":   order,
	})
}
