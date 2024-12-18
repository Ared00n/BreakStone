package main

import (
	"github.com/gorilla/mux"
)

func setuplinks() *mux.Router {
	links := mux.NewRouter()
	links.HandleFunc("/", homeHandler)
	links.HandleFunc("/login", loginHandler)
	links.HandleFunc("/register", registerHandler)
	links.HandleFunc("/logout", logoutHandler)
	links.HandleFunc("/personal_cabinet", personalCabinetHandler)
	links.HandleFunc("/uslugi", uslugi)
	links.HandleFunc("/candidates", candidatesHandler)
	links.HandleFunc("/register_candidate", registerCandidateHandler)
	links.HandleFunc("/whatis", whatis)
	links.HandleFunc("/add-comment", addCommentHandler)

	links.HandleFunc("/works", worksHandler).Methods("GET")
	links.HandleFunc("/works/add", addWorkHandler).Methods("POST")
	links.HandleFunc("/works/delete", deleteWorkHandler).Methods("DELETE")
	links.HandleFunc("/works/update", updateWorkHandler).Methods("POST")
	links.HandleFunc("/works/{id}", getWorkHandler).Methods("GET")

	return links
}
