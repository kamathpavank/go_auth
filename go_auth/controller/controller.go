package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"../model"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var users []model.User

// type Tag struct {
// 	password string `json:"password"`
// }

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user model.User

	db, err := sql.Open("mysql", "root:abcd@tcp(127.0.0.1:3306)/user")

	if err != nil {
		error := "Error while hashing"
		json.NewEncoder(w).Encode(error)
	}

	//decode body into user struct
	json.NewDecoder(r.Body).Decode(&user)

	//check if username already exists or not
	var tag string
	var pass string

	rs := db.QueryRow("SELECT username,password FROM users where username = ?", user.Username).Scan(&tag, &pass)

	if string(tag) == string(user.Username) {
		json.NewEncoder(w).Encode("user already exists")
		defer db.Close()
	} else {
		fmt.Println(rs)

		// if tag != nil {
		// 	json.NewEncoder(w).Encode("username already exists")
		// }

		//hash the password
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

		if err != nil {
			error := "Error while hashing"
			json.NewEncoder(w).Encode(error)
		}

		user.Password = string(hash)

		//append to users slice.
		users = append(users, user)

		if err != nil {
			panic(err.Error())
		}

		insert, err := db.Query("INSERT INTO users VALUES ( ?, ? )", user.Username, user.Password)

		if err != nil {
			error := err
			json.NewEncoder(w).Encode(error)
		}
		fmt.Println(insert)
		defer db.Close()
		json.NewEncoder(w).Encode(user)
	}

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user model.User

	//decode body into user struct
	json.NewDecoder(r.Body).Decode(&user)

	//hash the password

	db, err := sql.Open("mysql", "root:abcd@tcp(127.0.0.1:3306)/user")

	if err != nil {
		error := err
		json.NewEncoder(w).Encode(error)
	}

	if err != nil {
		panic(err.Error())
	}
	var tag string

	rs := db.QueryRow("SELECT password FROM users where username = ?", user.Username).Scan(&tag)

	msg := bcrypt.CompareHashAndPassword([]byte(tag), []byte(user.Password))

	fmt.Println(msg)

	if msg == nil {
		json.NewEncoder(w).Encode("successfull")
	} else {
		json.NewEncoder(w).Encode("unsuccessfull")
	}

	fmt.Println(rs)

	defer db.Close()

}
