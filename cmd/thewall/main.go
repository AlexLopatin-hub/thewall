package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_conn "github.com/AlexLopatin-hub/thewall/internal/core/repository/conn"
	core_http_server "github.com/AlexLopatin-hub/thewall/internal/core/transport/http/server"
	posts_repository_postgres "github.com/AlexLopatin-hub/thewall/internal/features/posts/repository/postgres"
	posts_service "github.com/AlexLopatin-hub/thewall/internal/features/posts/service"
	posts_transport_http "github.com/AlexLopatin-hub/thewall/internal/features/posts/transport/http"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	conn, err := core_conn.NewConn(
		ctx,
		core_conn.NewConfigMust(),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	postsRepository := posts_repository_postgres.NewPostsRepository(*conn)
	postsService := posts_service.NewPostsService(postsRepository)
	postsTransport := posts_transport_http.NewPostsHTTPHandler(postsService)

	// testRoute := core_http_server.NewRoute(
	// 	"GET",
	// 	"/test",
	// 	func(w http.ResponseWriter, r *http.Request) {
	// 		w.WriteHeader(http.StatusOK)
	// 		w.Write([]byte("Hello, thewall"))
	// 	},
	// )

	server := core_http_server.NewHTTPServer()
	server.RegisterRoutes(postsTransport.Routes()...)

	fmt.Println("starting HTTP server...")
	if err := server.Run(ctx); err != nil {
		fmt.Println("error starting HTTP server:", err)
	}
}
