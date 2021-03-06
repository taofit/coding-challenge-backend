package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/taofit/coding-challenge-backend/internal"
)

func handleOfficers(router *mux.Router) {
	router.HandleFunc("/officers", internal.GetOfficers).Methods("GET")
	router.HandleFunc("/officers/{id}", internal.GetOfficer).Methods("GET")
	router.HandleFunc("/officers/{id}", internal.UpdateOfficer).Methods("PUT")
	router.HandleFunc("/officers", internal.CreateOfficer).Methods("POST")
	router.HandleFunc("/officers/{id}", internal.DeleteOfficer).Methods("DELETE")
}

func handleBikeThefts(router *mux.Router) {
	router.HandleFunc("/bike-thefts", internal.CreateCase).Methods("POST")
	router.HandleFunc("/bike-thefts", internal.GetCases).Methods("GET")
	router.HandleFunc("/bike-thefts-no-image", internal.CreateCaseNoImage).Methods("POST")
	router.HandleFunc("/bike-thefts/{id}", internal.GetCase).Methods("GET")
	router.HandleFunc("/bike-thefts/{id}", internal.UpdateCase).Methods("PUT")
	router.HandleFunc("/bike-thefts/image/{id}", internal.GetImage).Methods("GET")
}

func assignCaseToOfficer(router *mux.Router) {
	router.HandleFunc("/case-to-officer", internal.AssignCaseToEnOfficer).Methods("POST")
}

func autoAssignCaseToOfficer() {
	for {
		time.Sleep(30 * time.Minute)
		internal.AssignCases()
	}
}

func initializeRoutes(router *mux.Router) {
	handleOfficers(router)
	handleBikeThefts(router)
	assignCaseToOfficer(router)
}

func main() {
	fmt.Println("Bike Theft Report API")
	router := mux.NewRouter()
	initializeRoutes(router)
	go autoAssignCaseToOfficer()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err.Error())
	}

}
