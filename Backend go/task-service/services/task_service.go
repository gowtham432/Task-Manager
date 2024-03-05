package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	taskmodels "task-manager-microservices/task-service/models"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

func AddTask(w http.ResponseWriter, r *http.Request) {
	addTaskForTheUser(w, r)
}

func addTaskForTheUser(w http.ResponseWriter, r *http.Request) {
	var task taskmodels.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, "Bad Input from Users", http.StatusBadGateway)
		return
	}
	var er error

	_, er = db.Exec("insert into task (task_title, asignee_username, created_username, task_description, task_status) values (?,?,?,?,?)", task.Title, task.Assignee_UserName, task.Created_UserName, task.Description, task.Status)

	if er != nil {
		http.Error(w, "Bad Input from the user", http.StatusBadRequest)
		return
	}

	if err == nil && er == nil {
		json.NewEncoder(w).Encode("Task created successfully")
	}
}

func GetTaskForUser(w http.ResponseWriter, r *http.Request) {
	getTaskForUser(w, r)
}

func getTaskForUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	rows, err := db.Query("select * from task where asignee_username=?", username)

	if err != nil {
		http.Error(w, "No task present for User", http.StatusBadRequest)
		return
	}

	var tasks []taskmodels.Task

	for rows.Next() {
		var task taskmodels.Task
		err := rows.Scan(&task.Title, &task.Assignee_UserName, &task.Created_UserName, &task.Description, &task.Status)
		if err != nil {
			http.Error(w, "Error reading rows", http.StatusBadGateway)
			return
		}

		if username == task.Assignee_UserName {
			tasks = append(tasks, task)
		}
	}

	if len(tasks) == 0 {
		json.NewEncoder(w).Encode("No Task assigned to this user")
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func ChangeTaskStatus(w http.ResponseWriter, r *http.Request) {
	changeTaskStatusForUser(w, r)
}

func changeTaskStatusForUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	title := vars["title"]
	status := vars["status"]

	res, err := db.Exec("update task set task_status=? where asignee_username=? and task_title=?", status, username, title)

	if err != nil {
		http.Error(w, "Error in updating status", http.StatusBadRequest)
		return
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		http.Error(w, "Error in getting rows Affected after update", http.StatusBadRequest)
		return
	}

	if rowsAffected == 0 {
		json.NewEncoder(w).Encode("No row got affected")
		return
	} else {
		json.NewEncoder(w).Encode("Status updated successfully")
		return
	}
}

func GetTaskCreatedUser(w http.ResponseWriter, r *http.Request) {
	getTaskCreatedUser(w, r)
}

func getTaskCreatedUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	rows, err := db.Query("select * from task where created_username=?", username)

	if err != nil {
		http.Error(w, "No task present for User", http.StatusBadRequest)
		return
	}

	var tasks []taskmodels.Task

	for rows.Next() {
		var task taskmodels.Task
		err := rows.Scan(&task.Title, &task.Assignee_UserName, &task.Created_UserName, &task.Description, &task.Status)
		if err != nil {
			http.Error(w, "Error reading rows", http.StatusBadGateway)
			return
		}

		if username == task.Created_UserName {
			tasks = append(tasks, task)
		}
	}

	if len(tasks) == 0 {
		json.NewEncoder(w).Encode("No Task created by the user")
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func ChangeTaskStatusCreated(w http.ResponseWriter, r *http.Request) {
	changeTaskStatusForCreatedUser(w, r)
}

func changeTaskStatusForCreatedUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	title := vars["title"]
	status := vars["status"]

	res, err := db.Exec("update task set task_status=? where created_username=? and task_title=?", status, username, title)

	if err != nil {
		http.Error(w, "Error in updating status", http.StatusBadRequest)
		return
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		http.Error(w, "Error in getting rows Affected after update", http.StatusBadRequest)
		return
	}

	if rowsAffected == 0 {
		json.NewEncoder(w).Encode("No row got affected")
		return
	} else {
		json.NewEncoder(w).Encode("Status updated successfully")
		return
	}
}
