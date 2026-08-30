package posts_transport_http

import (
	"time"

	"github.com/AlexLopatin-hub/thewall/internal/core/domain"
)

type PostDTOResponse struct {
	ID        int       `json:"id"`
	Version   int       `json:"version"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func postDTOFromDomain(post domain.Post) PostDTOResponse {
	return PostDTOResponse{
		ID:        post.ID,
		Version:   post.Version,
		Text:      post.Text,
		CreatedAt: post.CreatedAt,
	}
}
