package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"go-server/models"
	"strconv"
	"encoding/json"
)

func GetUser (w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)
	id, err := strconv.Atoi(query["userId"])
	if err != nil {
		panic(err)
	}
	result := models.FindUser(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	jErr := json.NewEncoder(w).Encode(result)

	if jErr != nil {
		panic(jErr)
	}
}

func GetAllUsers (w http.ResponseWriter, r *http.Request) {
	result := models.FindAllUsers()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	jErr := json.NewEncoder(w).Encode(result)

	if jErr != nil {
		panic(jErr)
	}
}

type userName struct{
	Name string
}

func PostUser (w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	/*
	ANOTHER OPTION OF PARSING BODIES
	var user userName
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &user);
	*/
	var user userName
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	print(user.Name)
	postedUser := models.CreateUser(user.Name)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	jErr := json.NewEncoder(w).Encode(postedUser)

	if jErr != nil {
		panic(jErr)
	}
}