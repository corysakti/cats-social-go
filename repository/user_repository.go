package repository

import (
	"context"
	"database/sql"
	"github.com/corysakti/cats-social-go/model/entity"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
	Update(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
	Delete(ctx context.Context, tx *sql.Tx, user entity.User)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (entity.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.User
}
