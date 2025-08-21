package controllers

import (
	"e-com/internal"
	"e-com/internal/reponse"
	"e-com/usecase"
	"errors"
	"net/http"
)

// GetUserOrdersController godoc
// @Summary Get user orders
// @Description Retrieve all orders for the authenticated user
// @Tags Orders
// @Accept json
// @Produce json
// @Success 200 {object} reponse.SuccessResponse "Orders retrieved successfully"
// @Failure 401 {object} reponse.ErrorResponse "Unauthorized - Authentication required"
// @Failure 500 {object} reponse.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /orders [get]
func GetUserOrdersController(w http.ResponseWriter, r *http.Request) {
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

	orders, err := usecase.GetOrdersByUserID(userIDStr)
	if err != nil {
		reponse.Error(w, 500, "Failed to get orders: "+err.Error(), err)
		return
	}

	reponse.Success(w, 200, "Orders retrieved successfully", orders)
}

// GetOrderByIDController godoc
// @Summary Get order by ID
// @Description Retrieve a specific order by its ID for the authenticated user
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} reponse.SuccessResponse "Order retrieved successfully"
// @Failure 400 {object} reponse.ErrorResponse "Bad request - Order ID required"
// @Failure 401 {object} reponse.ErrorResponse "Unauthorized - Authentication required"
// @Failure 403 {object} reponse.ErrorResponse "Access denied"
// @Failure 404 {object} reponse.ErrorResponse "Order not found"
// @Security BearerAuth
// @Router /order/{id} [get]
func GetOrderByIDController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

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
