package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/tomasen/realip"
	"golang.org/x/time/rate"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (app *application) rateLimiter(next http.Handler) http.Handler {
	type client struct {
		limiter  *rate.Limiter // a bucket of tokens for each client
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client) // client IP -> client strict
	)

	// A goroutine to periodically clean up the clients map
	go func() {
		for {
			time.Sleep(time.Minute) // runs each minute

			mu.Lock() // this goroutine also accesses the map, gotta lock it

			// if the client has been inactive for more than 3 minutes, remove it
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}

			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if app.config.limiter.enabled {
			// getting the client IP, we're using realip package, cause sometimes the IP could be that of loadbalancer or nginx, but we want the client's real IP
			// realip helps us get the actual client's ip if present, else defaults
			ip := realip.FromRequest(r)

			mu.Lock() // locking the resource(clients map, this would be redis in prod)

			// if client's IP doesnt already exists in the map, then we add one
			if _, found := clients[ip]; !found {
				clients[ip] = &client{limiter: rate.NewLimiter(rate.Limit(app.config.limiter.rps), app.config.limiter.burst)}
			}

			clients[ip].lastSeen = time.Now()

			// If client ran out of tokens, give back err resp
			if !clients[ip].limiter.Allow() {
				mu.Unlock()
				app.rateLimitExceededResponse(w, r)
				return
			}

			mu.Unlock()
		}
		next.ServeHTTP(w, r)
	})
}
