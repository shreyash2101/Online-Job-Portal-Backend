package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title"`
	Type string `json:"type"`
	Location string `json:"location"`
	Description string `json:"description"`
	Salary string `json:"salary"`
	Company *Company `json:"company"`
}

type Company struct {
	Name string `json:"name"`
	Description string `json:"description"`
	ContactEmail string `json:"contactEmail"`
	ContactPhone string `json:"contactPhone"`
}