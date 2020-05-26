package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pravinhp/gotest/crud/models"
)

const (
	url = "mongodb://0.0.0.0:27017"
)

func main() {

	fmt.Println("Staring application....")
	router := mux.NewRouter()
	router.HandleFunc("/person", GetPerson).Methods("GET")
	router.HandleFunc("/person", CreatePerson).Methods("POST")
	router.HandleFunc("/person/{id}", GetPersonById).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func GetPersonById(response http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	fmt.Println(id)
	person := models.GetPersonById(id)

	js, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(js)

}

func GetPerson(response http.ResponseWriter, request *http.Request) {

	person := models.GetPerson()
	response.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(person)

	if err != nil {
		panic(err)
	}
	response.Write(js)

}

type Person models.Person

func CreatePerson(response http.ResponseWriter, request *http.Request) {

	var p Person
	b, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	err := json.Unmarshal(b, &p)
	if err != nil {
		panic(err)
	}

	person := models.New(p.FirstName, p.Lastname)
	fmt.Println(p)
	//data := &p{FirstName: personmap["firstname"], Lastname: personmap["lastname"]}
	id := person.CreatePerson()

	js, err := json.Marshal(id)
	// fmt.Println("id in main", id)
	//data := models.GetPersonById(id)

	final, err := json.Marshal(js)

	if err != nil {
		panic(err)
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(final)

}
