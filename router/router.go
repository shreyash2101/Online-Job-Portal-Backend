package router

import (
	"github.com/gorilla/mux"
	controller "github.com/shreyash2101/online-job-portal/controllers"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/jobs", controller.GetJobs).Methods("GET")
	router.HandleFunc("/jobs", controller.AddJob).Methods("POST")
	router.HandleFunc("/jobs/{id}", controller.GetJobByID).Methods("GET")
	router.HandleFunc("/jobs/{id}", controller.UpdateJob).Methods("PUT")
	router.HandleFunc("/jobs/{id}", controller.DeleteJob).Methods("DELETE")

	return router
}