package hellowordls

import "github.com/jmoiron/sqlx"

type HelloWorldRepositories struct {
	db *sqlx.DB
}

func NewHelloWorldRepositories(db *sqlx.DB) *HelloWorldRepositories {
	return &HelloWorldRepositories{db: db}
}
