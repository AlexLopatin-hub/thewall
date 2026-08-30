package posts_repository_postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/AlexLopatin-hub/thewall/internal/core/domain"
)

func (r *PostsRepository) CreatePost(
	ctx context.Context,
	post domain.Post,
) (domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Conn.OpTimeout())
	defer cancel()

	now := time.Now()
	query := `
		INSERT INTO thewall.posts (text, created_at)
		VALUES($1, $2)
		RETURNING id, version, text, created_at
	`

	row := r.Conn.QueryRow(ctx, query, post.Text, now)

	var postModel PostModel
	err := row.Scan(
		&postModel.ID,
		&postModel.Version,
		&postModel.Text,
		&postModel.CreatedAt,
	)

	if err != nil {
		return domain.Post{}, fmt.Errorf("scan error: %w", err)
	}

	postDomain := domain.NewPost(
		postModel.ID,
		postModel.Version,
		postModel.Text,
		postModel.CreatedAt,
	)

	return postDomain, nil
}
