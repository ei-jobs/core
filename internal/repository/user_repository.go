package repository

import (
	"database/sql"

	"github.com/aidosgal/gust/internal/dto"
	"github.com/aidosgal/gust/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByPhone(phone string, app_id int) (model.User, error) {
	var user model.User
	query := `
		SELECT id, name, phone, password
		FROM users
		WHERE phone = $1 AND app_id = $2
	`

	err := r.db.QueryRow(query, phone, app_id).Scan(&user.Id, &user.Name, &user.Phone, &user.Password)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetAppByUser(app_id int) (model.App, error) {
	var app model.App
	query := `
		SELECT id, name, secret
		FROM apps
		WHERE id = $1
	`

	r.db.QueryRow(query, app_id).Scan(&app.Id, &app.Name, &app.Secret)

	return app, nil
}

func (r *UserRepository) GetUser(user_id int) (model.User, error) {
	var user model.User
	query := `
		SELECT id, name, phone, password, description
		FROM users
		WHERE user_id = $1
	`

	r.db.QueryRow(query, user_id).Scan(&user.Id, &user.Name, &user.Phone, &user.Password, &user.Description)

	return user, nil
}

func (r *UserRepository) CreateUser(req dto.RegisterRequest) (int, error) {
	query := `
		INSERT INTO users (name, phone, password, app_id)
		VALUES($1, $2, $3, $4)
	`

	var user_id int

	r.db.QueryRow(query, req.Name, req.Phone, req.Password, req.AppId).Scan(&user_id)

	return user_id, nil
}
