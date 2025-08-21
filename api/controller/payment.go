package controllers

import (
	"e-com/internal"
	"e-com/internal/reponse"
	"e-com/usecase"
	"encoding/json"
	"errors"
	"net/http"
)

// CreateCheckoutSessionController godoc
// @Summary Create Stripe checkout session
// @Description Create a new Stripe checkout session for payment processing
// @Tags Payment
// @Accept json
// @Produce json
// @Param checkout body map[string]interface{} true "Checkout session details"
// @Success 200 {object} reponse.SuccessResponse "Checkout session created successfully"
// @Failure 400 {object} reponse.ErrorResponse "Bad request - Invalid input data"
// @Failure 401 {object} reponse.ErrorResponse "Unauthorized - Authentication required"
// @Failure 500 {object} reponse.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /payment/checkout [post]
func CreateCheckoutSessionController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

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

// PaymentSuccessController godoc
// @Summary Handle successful payment
// @Description Handle successful payment callback from Stripe
// @Tags Payment
// @Accept json
// @Produce json
// @Success 200 {object} reponse.SuccessResponse "Payment successful"
// @Failure 400 {object} reponse.ErrorResponse "Bad request - Missing session ID"
// @Failure 500 {object} reponse.ErrorResponse "Internal server error"
// @Router /payment/success [post]
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

// PaymentCancelController godoc
// @Summary Handle cancelled payment
// @Description Handle cancelled payment callback from Stripe
// @Tags Payment
// @Accept json
// @Produce json
// @Success 200 {object} reponse.SuccessResponse "Payment cancelled"
// @Router /payment/cancel [post]
func PaymentCancelController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

	reponse.Success(w, 200, "Payment cancelled", map[string]interface{}{
		"status": "cancelled",
	})
}
