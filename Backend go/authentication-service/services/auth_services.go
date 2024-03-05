package services

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"task-manager-microservices/authentication-service/models"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SqlConnection() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "Sai@1243",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "TaskManager",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Print("Connected to DB")

}

func GetAllRegisteredUsersController(w http.ResponseWriter, r *http.Request) {
	var registeredUsers []models.Registration = getAllUsersFromDB(w)
	// Return the users as JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(registeredUsers)

}

func getAllUsersFromDB(w http.ResponseWriter) []models.Registration {

	var registeredUsers []models.Registration

	if db == nil {
		http.Error(w, "Database connection is nil", http.StatusInternalServerError)
	}

	rows, err := db.Query("select email,username,password from registration")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for rows.Next() {
		var user models.Registration

		err := rows.Scan(&user.Email, &user.UserName, &user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		registeredUsers = append(registeredUsers, user)

	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return registeredUsers
}

func CheckLogInDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var loggedInUser models.Login

	err := json.NewDecoder(r.Body).Decode(&loggedInUser)

	if err != nil {
		log.Fatal(err)
	}

	isUserExsist, username := isLoggedIn(w, &loggedInUser)

	if !isUserExsist {
		json.NewEncoder(w).Encode("Invalid login credentials")
	} else {
		json.NewEncoder(w).Encode(username)
	}

}

func isLoggedIn(w http.ResponseWriter, logInDetails *models.Login) (bool, string) {

	var registeredUsers []models.Registration = getAllUsersFromDB(w)

	var flag bool

	for _, user := range registeredUsers {
		if user.UserName == logInDetails.UserName && user.Password == logInDetails.Password {
			flag = true
		}
	}

	var err error

	var username string

	err = db.QueryRow("select username from registration where username= ?", logInDetails.UserName).Scan(&username)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given username")
		} else {
			log.Fatal(err)
		}
	}

	return flag, username
}

func PostRegistartionContorller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var registeredUsers []models.Registration = getAllUsersFromDB(w)

	var registeredUser models.Registration

	err := json.NewDecoder(r.Body).Decode(&registeredUser)

	if err != nil {
		log.Fatal(err)
	}

	var flag bool

	for _, user := range registeredUsers {
		if user.Email == registeredUser.Email || user.UserName == registeredUser.UserName {
			flag = true
		}
	}

	if flag {
		json.NewEncoder(w).Encode("User Already exsist")
	} else {
		insertOneItem(w, &registeredUser)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(registeredUser)
	}

}

func insertOneItem(w http.ResponseWriter, registeredUser *models.Registration) {

	_, err := db.Exec("INSERT INTO registration (email, username, password) VALUES (?, ?, ?)", registeredUser.Email, registeredUser.UserName, registeredUser.Password)

	if err != nil {
		log.Fatal(err)
	}

}