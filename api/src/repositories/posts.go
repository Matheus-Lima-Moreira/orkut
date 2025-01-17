package repositories

import (
	"api/src/models"
	"database/sql"
)

type posts struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *posts {
	return &posts{db}
}

func (repository posts) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		`
		INSERT INTO 
			posts 
			(title, content, author_id)
		VALUES
			(?, ?, ?)
		`,
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}
