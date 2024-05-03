package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func ExecSql() {
	db := NewDB()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES ('eko', 'Eko')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success")
}

func QuerySQL() {
	db := NewDB()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name from customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

func QuerySQLComplex() {
	db := NewDB()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_At from customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		fmt.Println("Emai:", email)
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Birth Date:", birthDate)
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
	}

}
