package impl

import (
	"context"
	"database/sql"
	"errors"
	"github.com/corysakti/cats-social-go/model/entity"
	"strconv"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepositoryImpl(db *sql.DB) *CommentRepositoryImpl {
	return &CommentRepositoryImpl{DB: db}
}

func (repository CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comments)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return entity.Comment{}, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository CommentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ?"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comments)
		return comment, nil
	} else {
		return comment, errors.New("Comment with Id" + strconv.Itoa(int(id)) + " not found")
	}
}

func (repository CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comments)
		comments = append(comments, comment)
	}
	return comments, nil
}
