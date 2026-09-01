package posts_repository_postgres

import (
	"context"
	"fmt"

	"github.com/AlexLopatin-hub/thewall/internal/core/domain"
)

func (r *PostsRepository) GetPosts(
	ctx context.Context,
) ([]domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Conn.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, text, created_at
		FROM thewall.posts
		ORDER BY id ASC
	`

	rows, err := r.Conn.Query(ctx, query)
	if err != nil {
		return []domain.Post{}, fmt.Errorf(
			"get posts from database: %w",
			err,
		)
	}

	posts := make([]domain.Post, 0)

	for rows.Next() {
		var postModel PostModel

		err := rows.Scan(
			&postModel.ID,
			&postModel.Version,
			&postModel.Text,
			&postModel.CreatedAt,
		)

		if err != nil {
			return []domain.Post{}, fmt.Errorf(
				"scan error: %w",
				err,
			)
		}

		postDomain := postDomainFromDTO(postModel)

		posts = append(posts, postDomain)
	}

	return posts, nil
}
