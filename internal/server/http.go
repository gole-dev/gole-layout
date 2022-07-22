package server

import (
	"github.com/gole-dev/gole-layout/internal/routers"
	"github.com/gole-dev/gole/pkg/app"
	"github.com/gole-dev/gole/pkg/transport/http"
)

// NewHTTPServer creates an HTTP server
func NewHTTPServer(c *app.ServerConfig) *http.Server {
	router := routers.NewRouter()

	srv := http.NewServer(
		http.WithAddress(c.Addr),
		http.WithReadTimeout(c.ReadTimeout),
		http.WithWriteTimeout(c.WriteTimeout),
	)

	srv.Handler = router

	return srv
}
