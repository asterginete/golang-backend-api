package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	// Define the number of requests allowed per duration
	MaxRequests = 100
	// Define the duration for rate limiting
	WindowDuration = 1 * time.Minute
)

var (
	// Use a map to store request counts for each client IP
	ipRequestCounts = make(map[string]int)
	// Use a map to store the timestamp of the first request for each client IP
	ipTimestamps = make(map[string]time.Time)
	// Use a mutex to synchronize access to the maps
	mu = &sync.Mutex{}
)

// RateLimiting is a middleware function to rate limit requests based on IP.
func RateLimiting() gin.HandlerFunc {
	return func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()

		clientIP := c.ClientIP()
		now := time.Now()

		// Check if the IP is already in the map
		if _, exists := ipRequestCounts[clientIP]; !exists {
			ipRequestCounts[clientIP] = 1
			ipTimestamps[clientIP] = now
			c.Next()
			return
		}

		// Check if the time window has expired
		if now.Sub(ipTimestamps[clientIP]) > WindowDuration {
			ipRequestCounts[clientIP] = 1
			ipTimestamps[clientIP] = now
			c.Next()
			return
		}

		// Check if the request count has exceeded the max allowed
		if ipRequestCounts[clientIP] >= MaxRequests {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests, please try again later."})
			c.Abort()
			return
		}

		// Increment the request count
		ipRequestCounts[clientIP]++
		c.Next()
	}
}
