package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "luke"
	password = "dsf"
	dbname   = "test_go"
)

type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var users []User

	rows, err := db.Query("SELECT id, first_name, last_name, email, age FROM users LIMIT $1", 3)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, age int
		var firstName, lastName, email string

		err = rows.Scan(&id, &firstName, &lastName, &email, &age)
		if err != nil {
			// handle this error
			panic(err)
		}
		users = append(users, User{ID: id, FirstName: firstName, Email: email, LastName: lastName, Age: age})

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Endpoint Hit: All Users")
	json.NewEncoder(w).Encode(users)
}

func dbConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// rows, err := db.Query("SELECT id, first_name FROM users LIMIT $1", 3)
	// if err != nil {
	// 	// handle this error better than this
	// 	panic(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var id int
	// 	var firstName string
	// 	err = rows.Scan(&id, &firstName)
	// 	if err != nil {
	// 		// handle this error
	// 		panic(err)
	// 	}
	// 	fmt.Println(id, firstName)
	// }
	// // get any error encountered during iteration
	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }

	// sqlStatement := `SELECT * FROM users WHERE id=$1;`
	// var user User
	// row := db.QueryRow(sqlStatement, 11)
	// err = row.Scan(&user.ID, &user.Age, &user.FirstName,
	// 	&user.LastName, &user.Email)
	// switch err {
	// case sql.ErrNoRows:
	// 	fmt.Println("No rows were returned!")
	// 	return
	// case nil:
	// 	fmt.Println(user)
	// default:
	// 	panic(err)
	// }
	// sqlStatement := `
	// 	INSERT INTO users (age, email, first_name, last_name)
	// 	VALUES ($1, $2, $3, $4)
	// 	RETURNING id`
	// id := 0
	// err = db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("New record ID is:", id)

	// sqlStatement := `
	// UPDATE users
	// SET first_name = $2, last_name = $3
	// WHERE id = $1
	// returning id, email;`
	// var email string
	// var id int
	// err = db.QueryRow(sqlStatement, 7, "NewFirst", "NewLast").Scan(&id, &email)
	// if err != nil {
	// 	panic(err)
	// }

	// // count, err := res.RowsAffected()
	// // if err != nil {
	// // 	panic(err)
	// // }

	// fmt.Println(id, email)

	// sqlStatement := `
	// delete from users;`
	// _, err = db.Exec(sqlStatement)
	// if err != nil {
	// 	panic(err)
	// }
}
