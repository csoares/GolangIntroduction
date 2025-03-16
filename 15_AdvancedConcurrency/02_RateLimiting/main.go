package main

import (
	"fmt"
	"sync"
	"time"
)

// RateLimiter implements a token bucket algorithm
type RateLimiter struct {
	rate       time.Duration // time between tokens
	capacity   int           // bucket capacity
	tokens     int           // current tokens
	lastRefill time.Time     // last token refill time
	mu         sync.Mutex    // for thread safety
}

func NewRateLimiter(rate time.Duration, capacity int) *RateLimiter {
	return &RateLimiter{
		rate:       rate,
		capacity:   capacity,
		tokens:     capacity,
		lastRefill: time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	timePassed := now.Sub(rl.lastRefill)
	tokensToAdd := int(timePassed / rl.rate)

	if tokensToAdd > 0 {
		rl.tokens = min(rl.capacity, rl.tokens+tokensToAdd)
		rl.lastRefill = now
	}

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// simulateAPIRequest simulates an API call
func simulateAPIRequest(id int) {
	fmt.Printf("API request %d processed at %v\n", id, time.Now().Format(time.StampMilli))
	time.Sleep(100 * time.Millisecond) // simulate work
}

func main() {
	// Create a rate limiter: 2 requests per second
	limiter := NewRateLimiter(500*time.Millisecond, 1)

	// Simulate 10 concurrent API requests
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Wait until we're allowed to proceed
			for !limiter.Allow() {
				time.Sleep(100 * time.Millisecond)
			}

			simulateAPIRequest(id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All requests completed")
}
