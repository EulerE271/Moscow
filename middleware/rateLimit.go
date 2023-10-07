package middleware

import (
	"net/http"
	"sync"
	"time"
)

var requestCounts = make(map[string]int)
var lock sync.Mutex

func RateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logic to get the User or IP Address
		userIP := r.RemoteAddr // Simple version, consider using a forwarded IP if behind a proxy

		lock.Lock()
		defer lock.Unlock()

		// Check if user has sent requests before
		if count, exists := requestCounts[userIP]; exists && count >= 10 {
			// Check if user has exceeded the limit
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		// If the user does not exist in the map, or has not exceeded the limit,
		// Increment the request count and let the request proceed
		requestCounts[userIP]++

		// Incrementing request count and setting expiry for reset
		go func(userIP string) {
			// Reset the count after 24 hours
			time.Sleep(24 * time.Hour)
			lock.Lock()
			requestCounts[userIP] = 0
			lock.Unlock()
		}(userIP)

		// Call the next handler/function
		next(w, r)
	}
}
