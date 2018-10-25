package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arrUser []Users
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select id,name,age from users")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.ID, &users.Name, &users.Age); err != nil {
			log.Fatal(err.Error())

		} else {
			arrUser = append(arrUser, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrUser

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func insertUsersMultipart(w http.ResponseWriter, r *http.Request) {

	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	// Form-data
	name := r.FormValue("name")
	age := r.FormValue("age")

	// x-www-form-urlencoded
	// name := r.Form.Get("name")
	// age := r.Form.Get("age")

	_, err = db.Exec("INSERT INTO users (name, age) values (?,?)",
		name,
		age,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
