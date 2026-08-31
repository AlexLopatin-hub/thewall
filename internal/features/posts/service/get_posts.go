package posts_service

import (
	"context"
	"fmt"

	"github.com/AlexLopatin-hub/thewall/internal/core/domain"
)

func (s *PostsService) GetPosts(
	ctx context.Context,
) ([]domain.Post, error) {
	posts, err := s.postsRepository.GetPosts(ctx)
	if err != nil {
		return []domain.Post{}, fmt.Errorf(
			"get posts: %w",
			err,
		)
	}

	return posts, nil
}
