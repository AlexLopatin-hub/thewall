package domain

import (
	"fmt"
	"time"

	core_errors "github.com/AlexLopatin-hub/thewall/internal/core/errors"
)

type Post struct {
	ID        int
	Version   int
	Text      string
	CreatedAt time.Time
}

func NewPost(
	id int,
	version int,
	text string,
	createdAt time.Time,
) Post {
	return Post{
		ID:        id,
		Version:   version,
		Text:      text,
		CreatedAt: createdAt,
	}
}

func NewPostUninitialized(
	text string,
) Post {
	return Post{
		ID:        uninitializedID,
		Version:   uninitializedVersion,
		Text:      text,
		CreatedAt: UninitializedCreatedAt,
	}
}

func (p Post) Validate() error {
	textLength := len([]rune(p.Text))
	if textLength <= 1 || textLength > 1000 {
		return fmt.Errorf(
			"invalid `Text` length (%d): %w",
			textLength,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
