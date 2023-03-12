package adapter

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/davidterranova/userstore/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

const (
	getUserQuery    = `SELECT * FROM users WHERE id = $1`
	insertUserQuery = `INSERT INTO users (id, created_at, first_name, last_name, email) VALUES (:id, :created_at, :first_name, :last_name, :email)`
)

type PGUserRepository struct {
	db *sqlx.DB
}

type pgUser struct {
	Id uuid.UUID `db:"id"`

	CreatedAt time.Time `db:"created_at"`

	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

func NewPGUserRepository(db *sqlx.DB) *PGUserRepository {
	return &PGUserRepository{
		db: db,
	}
}

func (r *PGUserRepository) GetUser(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var user pgUser
	err := r.db.GetContext(ctx, &user, getUserQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return toUser(&user), nil
}

func (r *PGUserRepository) CreateUser(ctx context.Context, u *model.User) (*model.User, error) {
	user := fromUser(u)
	_, err := r.db.NamedExecContext(ctx, insertUserQuery, user)

	if err != nil {
		if isPgError(err, "23505") {
			return nil, ErrUserAlreadyExist
		}
		return nil, err
	}

	return u, err
}

func fromUser(u *model.User) *pgUser {
	return &pgUser{
		Id:        u.GetId(),
		CreatedAt: u.GetCreatedAt(),
		FirstName: u.GetFirstName(),
		LastName:  u.GetLastName(),
		Email:     u.GetEmail(),
	}
}

func toUser(u *pgUser) *model.User {
	return model.NewUserFromRepository(
		u.Id,
		u.CreatedAt,
		u.FirstName,
		u.LastName,
		u.Email,
	)
}

func isPgError(err error, code string) bool {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		pgErr, _ = err.(*pgconn.PgError)
		if pgErr.Code == code {
			return true
		}
	}

	return false
}
