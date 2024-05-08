package impl

import (
	"context"
	"database/sql"
	"errors"
	"github.com/corysakti/cats-social-go/helper"
	"github.com/corysakti/cats-social-go/model/entity"
	"github.com/corysakti/cats-social-go/repository"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (repository UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	SQL := `insert into "user"(email, name, password) values ($1,$2, $3)`
	var id int
	err := tx.QueryRowContext(ctx, SQL, user.Email, user.Name, user.Password).Scan(&id)
	helper.PanicIfError(err)

	user.Id = int32(id)
	return user
}

func (repository UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	SQL := `update "user" set email = $1, name = $2, password = $3 where id = $4`
	_, err := tx.ExecContext(ctx, SQL, user.Email, user.Name, user.Password, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user entity.User) {
	SQL := `delete from "user" where id = $1`
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (entity.User, error) {
	SQL := `select id, email, name, password from "user" where id = $1`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := entity.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Name, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (repository UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.User {
	SQL := `select id, email, name, password from "user" where id = $1`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var users []entity.User
	for rows.Next() {
		user := entity.User{}
		err := rows.Scan(&user.Id, &user.Email, &user.Name, &user.Password)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}

func (repository UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.User, error) {
	SQL := `select id, email, name, password from "user" where email = $1`
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)
	defer rows.Close()

	user := entity.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Name, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}
