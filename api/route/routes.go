package route

import (
	controllers "e-com/api/controller"
	"e-com/api/middleware"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	r := http.NewServeMux()

	// Health check
	r.HandleFunc("/health", HandleRoot)

	// Auth routes
	r.HandleFunc("/registration", controllers.RegisterUserController)
	r.HandleFunc("/login", controllers.LoginController)

	// Product routes (protected)
	r.HandleFunc("/product", middleware.AuthMiddleware(controllers.Products))
	r.HandleFunc("/product/", middleware.AuthMiddleware(controllers.Products))

	// Payment routes
	r.HandleFunc("/payment/checkout", controllers.CreateCheckoutSessionController)
	r.HandleFunc("/payment/success", controllers.PaymentSuccessController)
	r.HandleFunc("/payment/cancel", controllers.PaymentCancelController)

	return r
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Health is good bro"}`))
}
