package psql

import (
	"GroupAssist/internal/domain"
	"github.com/jmoiron/sqlx"
	"log"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (a *AuthRepository) ApplyRegister(user domain.ApplyRegister) (domain.User, error) {
	var returnUser domain.User
	applyUserQuery := `UPDATE users SET username=$1, password_hash=$2, is_active=$3, register_token=NULL
             WHERE (register_token=$4 AND is_active=false) RETURNING id, name, username, role`

	err := a.db.QueryRow(applyUserQuery, user.SignUpInput.Username, user.SignUpInput.Password,
		user.IsActive, user.RegisterToken).Scan(
		&returnUser.ID, &returnUser.Name, &returnUser.Username, &returnUser.Role)

	log.Printf("Error: %v", err)
	log.Printf("returnUser: %+v", returnUser)

	return returnUser, err
}

func (a *AuthRepository) GetRegisterTokenHash(id int) (string, error) {
	var registerTokenHash string
	getRegisterTokenHashQuery := `SELECT register_token FROM users WHERE id=$1`
	err := a.db.QueryRow(getRegisterTokenHashQuery, id).Scan(&registerTokenHash)
	return registerTokenHash, err
}

func (a *AuthRepository) GetByCredentials(input domain.SignInInput) (domain.User, error) {
	var user domain.User
	getUserQuery := `SELECT id, name, username, role FROM users WHERE username=$1 AND password_hash=$2`
	err := a.db.QueryRow(getUserQuery, input.Username, input.Password).Scan(
		&user.ID, &user.Name, &user.Username, &user.Role)
	return user, err
}
