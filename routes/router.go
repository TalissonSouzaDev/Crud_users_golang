package routes

//go get -u github.com/gorilla/mux

import (
	"CrudGo/handles"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func IniRoute() {
	routes()
}

func routes() {

	router := mux.NewRouter()

	router.HandleFunc("/userlist", handles.ListUsers).Methods("GET")
	router.HandleFunc("/showuser/{id}", handles.ShowUser).Methods("GET")
	router.HandleFunc("/createuser", handles.CreateUser).Methods("POST")
	router.HandleFunc("/updateuser/{id}", handles.UpdateUser).Methods("PUT")
	router.HandleFunc("/deleteuser/{id}", handles.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}
