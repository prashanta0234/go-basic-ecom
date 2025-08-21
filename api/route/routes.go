package route

import (
	controllers "e-com/api/controller"
	"e-com/api/middleware"
	"e-com/internal/cache"
	"fmt"
	"net/http"
	"os"

	_ "e-com/docs" // Import generated docs

	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRoutes() *http.ServeMux {
	r := http.NewServeMux()

	// Health check
	r.HandleFunc("GET /health", HandleRoot)

	// Swagger documentation
	// Swagger UI
	r.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:5050/swagger/doc.json"),
	))

	// Serve Swagger JSON documentation
	r.HandleFunc("GET /swagger/doc.json", HandleSwagger)

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

func HandleSwagger(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Swagger JSON requested")

	// Get current working directory and print it for debugging
	wd, _ := os.Getwd()
	fmt.Printf("Current working directory: %s\n", wd)

	// Try multiple possible paths
	paths := []string{
		"docs/swagger.json",
		"../docs/swagger.json",
		"./docs/swagger.json",
	}

	var content []byte
	var err error

	for _, path := range paths {
		content, err = os.ReadFile(path)
		if err == nil {
			fmt.Printf("Found file at: %s\n", path)
			break
		}
		fmt.Printf("Tried path: %s - Error: %v\n", path, err)
	}

	if err != nil {
		fmt.Printf("All paths failed. Final error: %v\n", err)
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
}
