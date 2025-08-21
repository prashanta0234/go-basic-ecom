package route

import (
	controllers "e-com/api/controller"
	"e-com/api/middleware"
	"e-com/internal/cache"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	r := http.NewServeMux()

	// Health check
	r.HandleFunc("GET /health", HandleRoot)

	// Auth routes
	r.HandleFunc("POST /registration", http.HandlerFunc(controllers.RegisterUserController))
	r.HandleFunc("POST /login", http.HandlerFunc(controllers.LoginController))

	r.HandleFunc("GET /products", http.HandlerFunc(controllers.GetProducts))
	r.HandleFunc("POST /products", http.HandlerFunc(middleware.AuthMiddleware(controllers.CreateProduct)))
	r.HandleFunc("PUT /products/{id}", http.HandlerFunc(middleware.AuthMiddleware(controllers.UpdateProduct)))
	r.HandleFunc("DELETE /products/{id}", http.HandlerFunc(middleware.AuthMiddleware(controllers.DeleteProduct)))
	r.HandleFunc("GET /products/{id}", http.HandlerFunc(middleware.AuthMiddleware(controllers.GetProductByID)))

	// Payment routes
	r.HandleFunc("POST /payment/checkout", middleware.AuthMiddleware(controllers.CreateCheckoutSessionController))
	r.HandleFunc("POST /payment/success", controllers.PaymentSuccessController)
	r.HandleFunc("POST /payment/cancel", controllers.PaymentCancelController)

	// Order routes (protected)
	r.HandleFunc("GET /orders", middleware.AuthMiddleware(controllers.GetUserOrdersController))
	r.HandleFunc("GET /order/", middleware.AuthMiddleware(controllers.GetOrderByIDController))

	r.HandleFunc("/cache/invalidate", cache.CacheInvalidationEndpoint)

	return r
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Health is good bro"}`))
}
