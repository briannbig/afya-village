package database

import (
	"database/sql"

	"github.com/briannbig/afya-village/internal/domain/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

// FindById implements repository.UserRepository.
func (r *UserRepository) FindById(id string) (*model.User, error) {
	panic("unimplemented")
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user *model.User) error {
	query := `INSERT INTO users (id, username, email, telephone, password, created_at, updated_at) 
			   VALUES (:id, :username, :email, :telephone, :password, :created_at, :updated_at)`

	_, err := r.db.NamedExec(query, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE id = $1`

	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByIdentifier(identifier string) (*model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE username = ? OR email = ? OR telephone = ?`
	err := r.db.Get(&user, query, identifier, identifier, identifier)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
