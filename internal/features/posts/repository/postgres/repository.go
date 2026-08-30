package posts_repository_postgres

import (
	core_conn "github.com/AlexLopatin-hub/thewall/internal/core/repository/conn"
)

type PostsRepository struct {
	Conn core_conn.Conn
}

func NewPostsRepository(
	conn core_conn.Conn,
) *PostsRepository {
	return &PostsRepository{
		Conn: conn,
	}
}
