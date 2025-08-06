package cache

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func CacheStatsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		if strings.Contains(r.URL.Path, "/product") {
			w.Header().Set("Cache-Control", "public, max-age=1800") // 30 minutes
		}

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("Request %s %s completed in %v", r.Method, r.URL.Path, duration)
	}
}

func CacheInvalidationEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cacheService := NewCacheService()

	cacheType := r.URL.Query().Get("type")

	var err error
	switch cacheType {
	case "products":
		err = cacheService.InvalidateProductCaches()
	case "all":
		err = cacheService.DeletePattern("*")
	default:
		http.Error(w, "Invalid cache type. Use: products, all", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Failed to invalidate cache: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(`{"message": "Cache invalidated successfully", "type": "` + cacheType + `"}`))
}
