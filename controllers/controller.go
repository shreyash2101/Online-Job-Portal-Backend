package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	helper "github.com/shreyash2101/online-job-portal/helpers"
	model "github.com/shreyash2101/online-job-portal/models"
)

func GetJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	limit := r.URL.Query().Get("_limit")

	if limit != ""{
		limitCount,err := strconv.Atoi(limit)
		if err!=nil{
			log.Fatal(err)
		}
		fetchedJobs := helper.FetchSomeJobs(limitCount)
		json.NewEncoder(w).Encode(fetchedJobs)
	}else {
		fetchedJobs := helper.FetchAllJobs()
		json.NewEncoder(w).Encode(fetchedJobs)
	}
}

func GetJobByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	
	params := mux.Vars(r)

	fetchedJob := helper.FetchJobByID(params["id"])
	json.NewEncoder(w).Encode(fetchedJob)
}

func AddJob(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var job *model.Job
	_ = json.NewDecoder(r.Body).Decode(&job)

	helper.AddJob(job)
	json.NewEncoder(w).Encode(job)
}

func UpdateJob(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	var job *model.Job
	_ = json.NewDecoder(r.Body).Decode(&job)

	helper.UpdateJob(params["id"], job)

	json.NewEncoder(w).Encode(job)

}

func DeleteJob(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	helper.DeleteJob(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}