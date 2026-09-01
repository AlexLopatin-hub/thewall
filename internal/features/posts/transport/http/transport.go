package posts_transport_http

import (
	"context"

	"github.com/AlexLopatin-hub/thewall/internal/core/domain"
	core_http_server "github.com/AlexLopatin-hub/thewall/internal/core/transport/http/server"
)

type PostsHTTPHandler struct {
	postsService PostsService
}

type PostsService interface {
	CreatePost(
		ctx context.Context,
		post domain.Post,
	) (domain.Post, error)

	GetPosts(
		ctx context.Context,
	) ([]domain.Post, error)
}

func NewPostsHTTPHandler(
	postsService PostsService,
) *PostsHTTPHandler {
	return &PostsHTTPHandler{
		postsService: postsService,
	}
}

func (h *PostsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  "POST",
			Path:    "/posts",
			Handler: h.CreatePost,
		},
		{
			Method:  "GET",
			Path:    "/posts",
			Handler: h.GetPosts,
		},
	}
}
