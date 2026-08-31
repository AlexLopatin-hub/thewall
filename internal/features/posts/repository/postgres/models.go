package posts_repository_postgres

import (
	"time"

	"github.com/AlexLopatin-hub/thewall/internal/core/domain"
)

type PostModel struct {
	ID        int
	Version   int
	Text      string
	CreatedAt time.Time
}

func postDomainFromDTO(model PostModel) domain.Post {
	return domain.NewPost(
		model.ID,
		model.Version,
		model.Text,
		model.CreatedAt,
	)
}
