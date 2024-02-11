package psql

import (
	"GroupAssist/internal/domain"
	"errors"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Create(user domain.CreateUser) (domain.CreateUser, error) {
	createUserQuery := `INSERT INTO users (name, role, register_token) VALUES ($1, $2, $3) RETURNING id, name, role`
	err := u.db.QueryRow(createUserQuery, user.Name, user.Role, user.RegisterToken).Scan(
		&user.ID, &user.Name, &user.Role)
	return user, err
}

func (u *UserRepository) Update(id int, user domain.UpdateUserInput) error {
	updateUserQuery := `UPDATE users SET name=$1, role=$2 WHERE id=$3`
	_, err := u.db.Exec(updateUserQuery, user.Name, user.Role, id)
	return err
}

func (u *UserRepository) Delete(id int) error {
	deleteUserQuery := `DELETE FROM users WHERE id=$1`
	result, err := u.db.Exec(deleteUserQuery, id)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("not found")
	}
	return err
}

func (u *UserRepository) GetAll() ([]domain.User, error) {
	users := make([]domain.User, 0)
	getAllUsersQuery := `SELECT id, name, username, role FROM users`
	err := u.db.Select(&users, getAllUsersQuery)
	return users, err
}
