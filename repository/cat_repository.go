package repository

import (
	"context"
	"database/sql"
	"github.com/corysakti/cats-social-go/model/entity"
)

type CatRepository interface {
	Save(ctx context.Context, tx *sql.Tx, cat entity.Cat) entity.Cat
	Update(ctx context.Context, tx *sql.Tx, cat entity.Cat) entity.Cat
	Delete(ctx context.Context, tx *sql.Tx, cat entity.Cat)
	FindById(ctx context.Context, tx *sql.Tx, catId int) (entity.Cat, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Cat
}
