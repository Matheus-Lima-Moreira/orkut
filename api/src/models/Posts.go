package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"author_id,omitempty"`
	AuthorNick uint64    `json:"author_nick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
    return err
  }

  if err := post.format(); err != nil {
    return err
  }

  return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
    return errors.New("title must be provided")
  }

  if post.Content == "" {
    return errors.New("content must be provided")
  }

  return nil
}

func (post *Post) format() error {
	post.Title = strings.TrimSpace(post.Title)
  post.Content = strings.TrimSpace(post.Content)

  return nil
}