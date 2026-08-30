package posts_service

import (
	"context"

	"github.com/AlexLopatin-hub/thewall/internal/core/domain"
)

type PostsService struct {
	postsRepository PostsRepository
}

type PostsRepository interface {
	CreatePost(
		ctx context.Context,
		post domain.Post,
	) (domain.Post, error)
}

func NewPostsService(
	postsRepository PostsRepository,
) *PostsService {
	return &PostsService{
		postsRepository: postsRepository,
	}
}
