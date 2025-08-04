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

	checkoutURL, err := usecase.PaymentServiceByProductID(request.ProductID, request.Currency)
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

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Payment successful!",
		"status":  "success",
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
