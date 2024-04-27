package helper

import (
	"context"
	"fmt"
	"log"

	db "github.com/shreyash2101/online-job-portal/database"
	model "github.com/shreyash2101/online-job-portal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert a Job
func AddJob(job *model.Job) {
	insertResult,err := db.Collection.InsertOne(context.Background(),job)

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted a job with id:", insertResult.InsertedID )
}

// Delete a Job
func DeleteJob(jobId string) {
	id,_ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id":id}
	_, err := db.Collection.DeleteOne(context.Background(),filter)

	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println("Deleted a job with id:", id)
}

// Update a Job
func UpdateJob(jobId string, job *model.Job){
	id,_ := primitive.ObjectIDFromHex(jobId)

	filter := bson.M{"_id":id}
	update := bson.M{"$set": job }
	// bson.M{"title":job.Title, "type": job.Type, "description":job.Description, "location":job.Location, "salary":job.Salary, "company":job.Company}

	_, err := db.Collection.UpdateOne(context.Background(),filter,update)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Update a job with id:", id)
}

// Fetch Job by ID
func FetchJobByID(jobId string) bson.M {
	id,_ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id":id}
	cur, err := db.Collection.Find(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.Background())

	var job bson.M

	for cur.Next(context.Background()){
		err := cur.Decode(&job)
		if err != nil{
			log.Fatal(err)
		}
	}

	return job
}

// Fetch some Jobs
func FetchSomeJobs(limit int) []bson.M {
	cur, err := db.Collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.Background())


	var jobs []bson.M

	for cur.Next(context.Background()){
		var job bson.M
		err := cur.Decode(&job)
		if err != nil{
			log.Fatal(err)
		}

		jobs = append(jobs, job)
		limit--
		if(limit==0){
			break
		}
	}

	return jobs
}

// Fetch all Jobs
func FetchAllJobs() []bson.M {
	cur, err := db.Collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.Background())


	var jobs []bson.M

	for cur.Next(context.Background()){
		var job bson.M
		err := cur.Decode(&job)
		if err != nil{
			log.Fatal(err)
		}

		jobs = append(jobs, job)
	}

	return jobs
}