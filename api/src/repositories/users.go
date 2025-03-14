package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

func (repository users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repository users) List(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := repository.db.Query(
		"SELECT id, name, nick, email, created_at FROM users WHERE name LIKE ? OR nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) GetByID(ID uint64) (models.User, error) {
	rows, err := repository.db.Query("SELECT id, name, nick, email, created_at FROM users WHERE id = ?",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository users) GetByEmail(email string) (models.User, error) {
	row, err := repository.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err := row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) Follow(userID, followerID uint64) error {
	statement, err := repository.db.Prepare("INSERT INTO followers (user_id, follower_id) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (repository users) Unfollow(userID, followerID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM followers WHERE user_id =? AND follower_id =?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (repository users) GetFollowers(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		SELECT 
			u.id, u.name, u.nick, u.email, u.created_at
		FROM
			users u
		INNER JOIN 
			followers f
		ON 
			u.id = f.follower_id
		WHERE 
			f.user_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User

	for rows.Next() {
		var follower models.User

		if err = rows.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedAt,
		); err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

func (repository users) GetFollowing(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		SELECT 
			u.id, u.name, u.nick, u.email, u.created_at
		FROM
			users u
		INNER JOIN 
			followers f
		ON 
			u.id = f.user_id
		WHERE 
			f.follower_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var following []models.User

	for rows.Next() {
		var followedUser models.User

		if err = rows.Scan(
			&followedUser.ID,
			&followedUser.Name,
			&followedUser.Nick,
			&followedUser.Email,
			&followedUser.CreatedAt,
		); err != nil {
			return nil, err
		}

		following = append(following, followedUser)
	}

	return following, nil
}

func (repository users) GetPassword(userID uint64) (string, error) {
	rows, err := repository.db.Query("SELECT password FROM users WHERE id = ?", userID)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err := rows.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repository users) UpdatePassword(userID uint64, newPassword string) error {
	statement, err := repository.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(newPassword, userID); err != nil {
		return err
	}

	return nil
}
