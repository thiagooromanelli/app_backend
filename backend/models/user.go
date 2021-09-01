package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Task - Model of a basic task
type User struct {
	ID        primitive.ObjectID
	FirstName string
	LastName  string
	Age       int32
}
