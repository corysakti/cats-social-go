package repository

import (
	"context"
	"database/sql"
	"github.com/corysakti/cats-social-go/model/entity"
	"github.com/corysakti/cats-social-go/model/entity/relation"
)

type UserCatRepository interface {
	Save(ctx context.Context, tx *sql.Tx, userCat entity.UserCatMatch) entity.UserCatMatch
	Update(ctx context.Context, tx *sql.Tx, userCat entity.UserCatMatch) entity.UserCatMatch
	Delete(ctx context.Context, tx *sql.Tx, userCat entity.UserCatMatch)
	FindById(ctx context.Context, tx *sql.Tx, userCatId int) (relation.UserCatMatchRelation, error)
	FindByCatId(ctx context.Context, tx *sql.Tx, catId int) (relation.UserCatMatchRelation, error)
	FindAll(ctx context.Context, tx *sql.Tx) []relation.UserCatMatchRelation
}
