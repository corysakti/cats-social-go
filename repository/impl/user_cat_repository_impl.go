package impl

import (
	"context"
	"database/sql"
	"errors"
	"github.com/corysakti/cats-social-go/helper"
	"github.com/corysakti/cats-social-go/model/entity"
	"github.com/corysakti/cats-social-go/model/entity/relation"
	"github.com/corysakti/cats-social-go/repository"
)

type UserCatRepositoryImpl struct {
}

func NewUserCatRepositoryImpl() repository.UserCatRepository {
	return &UserCatRepositoryImpl{}
}

func (repository UserCatRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, userCat entity.UserCatMatch) entity.UserCatMatch {
	SQL := "insert into user_cat_match(user_id, cat_id, status) values ($1, $2, $3);"
	var id int
	err := tx.QueryRowContext(ctx, SQL, userCat.UserId, userCat.CatId, userCat.Status).Scan(&id)
	helper.PanicIfError(err)

	userCat.Id = int32(id)
	return userCat
}

func (repository UserCatRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, userCat entity.UserCatMatch) entity.UserCatMatch {
	SQL := "update user_cat_match set user_id = $1, cat_id = $2, status = $3 where id = $4"
	_, err := tx.ExecContext(ctx, SQL, userCat.UserId, userCat.CatId, userCat.Status, userCat.Id)
	helper.PanicIfError(err)

	return userCat
}

func (repository UserCatRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userCat entity.UserCatMatch) {
	SQL := "delete from user_cat_match where id = $1"
	_, err := tx.ExecContext(ctx, SQL, userCat.Id)
	helper.PanicIfError(err)
}

func (repository UserCatRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userCatId int) (relation.UserCatMatchRelation, error) {
	SQL := `select ucm.id, ucm.user_id, ucm.cat_id, c.name, c.race, c.sex, c.description, c.age_in_month, 
       		c.image_urls, c.created_at,u.email, u.name, u.password
       		from user_cat_match ucm  
       		INNER JOIN cat c ON c.id = ucm.cat_id
       		inner join "user" u on u.id = ucm.user_id
       		where ucm.id = $1`

	rows, err := tx.QueryContext(ctx, SQL, userCatId)
	helper.PanicIfError(err)
	defer rows.Close()

	userCatMatch := relation.UserCatMatchRelation{}
	if rows.Next() {
		err := rows.Scan(&userCatMatch.Id, &userCatMatch.UserId, &userCatMatch.CatId, &userCatMatch.CatName,
			&userCatMatch.Race, &userCatMatch.Sex, &userCatMatch.Description, &userCatMatch.AgeInMonth, &userCatMatch.ImageUrls,
			&userCatMatch.CreatedAt, &userCatMatch.Email, &userCatMatch.UserName, &userCatMatch.Password)
		helper.PanicIfError(err)
		return userCatMatch, nil
	} else {
		return userCatMatch, errors.New("category is not found")
	}
}

func (repository UserCatRepositoryImpl) FindByCatId(ctx context.Context, tx *sql.Tx, catId int) (relation.UserCatMatchRelation, error) {
	SQL := `select ucm.id, ucm.user_id, ucm.cat_id, c.name, c.race, c.sex, c.description, c.age_in_month, 
       		c.image_urls, c.created_at,u.email, u.name, u.password
       		from user_cat_match ucm  
       		INNER JOIN cat c ON c.id = ucm.cat_id
       		inner join "user" u on u.id = ucm.user_id
       		where c.id = $1`

	rows, err := tx.QueryContext(ctx, SQL, catId)
	helper.PanicIfError(err)
	defer rows.Close()

	userCatMatch := relation.UserCatMatchRelation{}
	if rows.Next() {
		err := rows.Scan(&userCatMatch.Id, &userCatMatch.UserId, &userCatMatch.CatId, &userCatMatch.CatName,
			&userCatMatch.Race, &userCatMatch.Sex, &userCatMatch.Description, &userCatMatch.AgeInMonth, &userCatMatch.ImageUrls,
			&userCatMatch.CreatedAt, &userCatMatch.Email, &userCatMatch.UserName, &userCatMatch.Password)
		helper.PanicIfError(err)
		return userCatMatch, nil
	} else {
		return userCatMatch, errors.New("category is not found")
	}
}

func (repository UserCatRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []relation.UserCatMatchRelation {
	SQL := `select ucm.id, ucm.user_id, ucm.cat_id, c.name, c.race, c.sex, c.description, c.age_in_month, 
       		c.image_urls, c.created_at,u.email, u.name, u.password
       		from user_cat_match ucm  
       		INNER JOIN cat c ON c.id = ucm.cat_id
       		inner join "user" u on u.id = ucm.user_id`

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var usersCatMatch []relation.UserCatMatchRelation
	for rows.Next() {
		userCatMatch := relation.UserCatMatchRelation{}
		err := rows.Scan(&userCatMatch.Id, &userCatMatch.UserId, &userCatMatch.CatId, &userCatMatch.CatName,
			&userCatMatch.Race, &userCatMatch.Sex, &userCatMatch.Description, &userCatMatch.AgeInMonth, &userCatMatch.ImageUrls,
			&userCatMatch.CreatedAt, &userCatMatch.Email, &userCatMatch.UserName, &userCatMatch.Password)
		helper.PanicIfError(err)
		usersCatMatch = append(usersCatMatch, userCatMatch)
	}
	return usersCatMatch
}
