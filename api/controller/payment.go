package controllers

import (
	"e-com/internal"
	"e-com/usecase"
	"encoding/json"
	"net/http"
)

func CreateCheckoutSessionController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

	if r.Method != "POST" {
		w.WriteHeader(404)
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

	var request struct {
		ProductID string `json:"product_id"`
		Currency  string `json:"currency"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if request.ProductID == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	if request.Currency == "" {
		request.Currency = "usd"
	}

	checkoutURL, err := usecase.PaymentServiceByProductID(request.ProductID, request.Currency, userIDStr)
	if err != nil {
		http.Error(w, "Failed to create checkout session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"checkout_url": checkoutURL,
		"message":      "Checkout session created successfully",
	})
}

func PaymentSuccessController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

	sessionID := r.URL.Query().Get("session_id")
	if sessionID == "" {
		http.Error(w, "Missing session ID", http.StatusBadRequest)
		return
	}

	userID, productID, err := usecase.GetStripeSessionDetails(sessionID)
	if err != nil {
		http.Error(w, "Failed to verify payment: "+err.Error(), http.StatusBadRequest)
		return
	}

	order, err := usecase.CreateOrder(userID, productID, sessionID)
	if err != nil {
		http.Error(w, "Failed to create order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Payment successful! Order created.",
		"status":   "success",
		"order_id": order.OrderID,
		"order":    order,
	})
}

func PaymentCancelController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Payment cancelled",
		"status":  "cancelled",
	})
}
