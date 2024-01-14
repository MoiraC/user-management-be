package models

import (
	"database/sql"
	"log"

	_ "github.com/ncruces/go-sqlite3"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var DB *sql.DB

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Gender  string `json:"gender"`
}

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func GetUsers() ([]User, error) {

	rows, err := DB.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		singleUser := User{}
		err = rows.Scan(&singleUser.Id, &singleUser.Name, &singleUser.Surname, &singleUser.Email, &singleUser.Gender)

		if err != nil {
			return nil, err
		}

		users = append(users, singleUser)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func AddUser(newUser User) (bool, error) {

	_, err := DB.Exec("INSERT INTO users (name, surname, email, gender) values(?,?,?,?)", newUser.Name, newUser.Surname, newUser.Email, newUser.Gender)
	log.Println("Inserted the user into database!")

	if err != nil {
		return false, err
	}

	return true, nil
}

func RemoveUser(id string) (bool, error) {

	_, err := DB.Exec("DELETE FROM Users WHERE id=$1", id)

	if err != nil {
		return false, err
	}

	log.Println("Deleted the user from database!")

	return true, nil
}

func UpdateUser(id string, changedUser User) (bool, error) {

	_, err := DB.Exec("UPDATE users SET name=?, surname=?, email=?, gender=? WHERE id=?;", changedUser.Name, changedUser.Surname, changedUser.Email, changedUser.Gender, id)

	if err != nil {
		return false, err
	}

	log.Println("Updated the user at database!")

	return true, nil
}

func GetUser(id string) (User, error) {

	var user User
	row := DB.QueryRow("SELECT * FROM users WHERE id=?;", id)
	err := row.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &user.Gender)

	if err != nil {
		return user, err
	}

	log.Println("Found the user at database!")

	return user, nil
}
