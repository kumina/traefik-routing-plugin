package traefik_routing_plugin

import (
	"context"
	"fmt"
	"net/http"
)

type Router struct {
	// Required by Traefik
	next http.Handler
	name string

	// Our custom configuration
	routes map[string]string
}

// Function needed for Traefik to recognize this module as a plugin
// Uses a generic http.Handler type from golang that we can use to work with the request
// by overriding different functions of the interface
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Routes) == 0 {
		return nil, fmt.Errorf("routes cannot be empty")
	}

	return &Router{
		routes:  config.Routes,
		next:     next,
		name:     name,
	}, nil
}

func (a *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for key, value := range a.routes {
		req.Header.Set(key, value)
	}

	a.next.ServeHTTP(rw, req)
}
