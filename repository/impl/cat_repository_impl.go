package impl

import (
	"context"
	"database/sql"
	"errors"
	"github.com/corysakti/cats-social-go/helper"
	"github.com/corysakti/cats-social-go/model/entity"
	"github.com/corysakti/cats-social-go/repository"
)

type CatRepositoryImpl struct {
}

func NewCatRepositoryImpl() repository.CatRepository {
	return &CatRepositoryImpl{}
}

func (repository CatRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, cat entity.Cat) entity.Cat {
	SQL := "insert into cat(name, race, sex, age_in_month, description, image_urls) values ($1, $2, $3, $4, $5, $6);"
	var id int
	err := tx.QueryRowContext(ctx, SQL, cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, cat.Description, cat.ImageUrls).Scan(&id)
	helper.PanicIfError(err)

	cat.Id = int32(id)
	return cat
}

func (repository CatRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, cat entity.Cat) entity.Cat {
	SQL := "update cat set name = $1, race = $2, sex = $3, age_in_month = $4, description = $5, image_urls = $6 where id = $7"
	_, err := tx.ExecContext(ctx, SQL, cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, cat.Description, cat.ImageUrls, cat.Id)
	helper.PanicIfError(err)

	return cat
}

func (repository CatRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, cat entity.Cat) {
	SQL := "delete from cat where id = $1"
	_, err := tx.ExecContext(ctx, SQL, cat.Id)
	helper.PanicIfError(err)
}

func (repository CatRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, catId int) (entity.Cat, error) {
	SQL := "select id, name, race, sex, age_in_month, description, image_urls from cat where id = $1"
	rows, err := tx.QueryContext(ctx, SQL, catId)
	helper.PanicIfError(err)
	defer rows.Close()

	cat := entity.Cat{}
	if rows.Next() {
		err := rows.Scan(&cat.Id, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.ImageUrls)
		helper.PanicIfError(err)
		return cat, nil
	} else {
		return cat, errors.New("category is not found")
	}
}

func (repository CatRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Cat {
	SQL := "select id, name, race, sex, age_in_month, description, image_urls from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var cats []entity.Cat
	for rows.Next() {
		cat := entity.Cat{}
		err := rows.Scan(&cat.Id, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.ImageUrls)
		helper.PanicIfError(err)
		cats = append(cats, cat)
	}
	return cats
}
