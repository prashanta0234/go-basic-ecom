package controllers

import (
	"e-com/internal"
	"e-com/internal/reponse"
	"e-com/usecase"
	"encoding/json"
	"errors"
	"net/http"
)

func CreateCheckoutSessionController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

	if r.Method != "POST" {
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

	var request struct {
		ProductID string `json:"product_id"`
		Currency  string `json:"currency"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		reponse.Error(w, 400, "Invalid request body", err)
		return
	}

	if request.ProductID == "" {
		reponse.Error(w, 400, "Product ID is required", errors.New("product ID is required"))
		return
	}

	if request.Currency == "" {
		request.Currency = "usd"
	}

	checkoutURL, err := usecase.PaymentServiceByProductID(request.ProductID, request.Currency, userIDStr)
	if err != nil {
		reponse.Error(w, 500, "Failed to create checkout session", err)
		return
	}

	reponse.Success(w, 200, "Checkout session created successfully", map[string]interface{}{
		"checkout_url": checkoutURL,
	})
}

func PaymentSuccessController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

	sessionID := r.URL.Query().Get("session_id")
	if sessionID == "" {
		reponse.Error(w, 400, "Missing session ID", errors.New("missing session ID"))
		return
	}

	userID, productID, err := usecase.GetStripeSessionDetails(sessionID)
	if err != nil {
		reponse.Error(w, 400, "Failed to verify payment: "+err.Error(), err)
		return
	}

	order, err := usecase.CreateOrder(userID, productID, sessionID)
	if err != nil {
		reponse.Error(w, 500, "Failed to create order: "+err.Error(), err)
		return
	}

	reponse.Success(w, 200, "Payment successful! Order created.", map[string]interface{}{
		"order_id": order.OrderID,
		"order":    order,
	})
}

func PaymentCancelController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

	reponse.Success(w, 200, "Payment cancelled", map[string]interface{}{
		"status": "cancelled",
	})
}
