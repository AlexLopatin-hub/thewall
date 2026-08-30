package posts_repository_postgres

import "time"

type PostModel struct {
	ID        int
	Version   int
	Text      string
	CreatedAt time.Time
}
