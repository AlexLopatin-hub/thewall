package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type HTTPServer struct {
	mux *http.ServeMux
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		mux: http.NewServeMux(),
	}
}

func (h *HTTPServer) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		h.mux.Handle(pattern, route.Handler)
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: h.mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and serve HTTP: %w", err)
		}
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			3*time.Second,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			server.Close()
			return fmt.Errorf("shutdown HTTP server: %w", err)
		}

		fmt.Println("HTTP server gracefully stopped")
	}

	return nil
}
