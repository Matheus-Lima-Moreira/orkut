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

func (repository posts) GetByID(postID uint64) (models.Post, error) {
	rows, err := repository.db.Query(`
		SELECT 
			p.*, 
			u.nick 
		FROM
		  posts p
		INNER JOIN
			users u
		ON 
		  u.id = p.author_id
		WHERE
		  p.id = ?
	`, postID)
	if err != nil {
		return models.Post{}, err
	}
	defer rows.Close()

	var post models.Post

	if rows.Next() {
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}
