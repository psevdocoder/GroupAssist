package psql

import (
	"GroupAssist/internal/domain"
	"github.com/jmoiron/sqlx"
	"time"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (a *AuthRepository) ApplyRegister(user domain.ApplyRegister) (domain.ResponseUser, error) {
	var returnUser domain.ResponseUser
	applyUserQuery := `UPDATE users SET username=$1, password_hash=$2, is_active=$3, register_token=NULL
             WHERE (register_token=$4 AND is_active=false) RETURNING id, name, username, role`

	err := a.db.QueryRow(applyUserQuery, user.SignUpInput.Username, user.SignUpInput.Password,
		user.IsActive, user.RegisterToken).Scan(
		&returnUser.ID, &returnUser.Name, &returnUser.Username, &returnUser.Role)

	return returnUser, err
}

func (a *AuthRepository) GetRegisterToken(id int) (string, error) {
	var registerTokenHash string
	getRegisterTokenHashQuery := `SELECT register_token FROM users WHERE id=$1`
	err := a.db.QueryRow(getRegisterTokenHashQuery, id).Scan(&registerTokenHash)
	return registerTokenHash, err
}

func (a *AuthRepository) GetByUsername(username string) (domain.User, error) {
	var user domain.User
	getUserQuery := `SELECT id, name, username, role, password_hash FROM users WHERE username=$1`
	err := a.db.QueryRow(getUserQuery, username).Scan(
		&user.ID, &user.Name, &user.Username, &user.Role, &user.PasswordHash)
	return user, err
}

func (a *AuthRepository) GetByID(id int) (domain.User, error) {
	var user domain.User
	getUserQuery := `SELECT id, name, username, role, password_hash FROM users WHERE id=$1`
	err := a.db.QueryRow(getUserQuery, id).Scan(
		&user.ID, &user.Name, &user.Username, &user.Role, &user.PasswordHash)
	return user, err
}

func (a *AuthRepository) SetRefreshToken(userID int, refreshToken string, expiresAt time.Time, ip string) error {
	query := `INSERT INTO sessions (user_id, refresh_token, expires_at, ip_address) VALUES ($1, $2, $3, $4)`
	_, err := a.db.Exec(query, userID, refreshToken, expiresAt, ip)
	return err
}

func (a *AuthRepository) GetRefreshToken(token string) (domain.Session, error) {
	var session domain.Session
	getRefreshTokenQuery := `DELETE FROM sessions WHERE refresh_token=$1 RETURNING user_id, refresh_token, expires_at, ip_address`
	if err := a.db.QueryRow(getRefreshTokenQuery, token).Scan(
		&session.UserID, &session.RefreshToken, &session.ExpiresAt, &session.IP); err != nil {
		return domain.Session{}, err
	}

	return session, nil
}
