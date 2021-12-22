package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"app.com/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

// GetConnection - Retrieves a client to the DocumentDB
func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")

	log.Printf(username)
	log.Printf(password)
	log.Printf(clusterEndpoint)

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)

	log.Printf(connectionURI)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}

func GetAllUsers() ([]*models.User, error) {
	var users []*models.User

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database("users")
	collection := db.Collection("users")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &users)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	return users, nil
}

// GetTaskByID Retrives a user by its id from db
func GetUserByID(id primitive.ObjectID) (*models.User, error) {
	var user *models.User

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database("users")
	collection := db.Collection("users")
	result := collection.FindOne(ctx, bson.D{})
	if result == nil {
		return nil, errors.New("could not find the requested user")
	}
	err := result.Decode(&user)

	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	log.Printf("Users: %v", user)
	return user, nil
}

//Create creating a task in a mongo
func CreateUser(task *models.User) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	task.ID = primitive.NewObjectID()

	result, err := client.Database("users").Collection("users").InsertOne(ctx, task)
	if err != nil {
		log.Printf("Could not create User: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

//Update updating an existing user in mongo
func UpdateUser(user *models.User) (*models.User, error) {
	var updatedUser *models.User
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	update := bson.M{
		"$set": user,
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}

	err := client.Database("users").Collection("users").FindOneAndUpdate(ctx, bson.M{"_id": user.ID}, update, &opt).Decode(&updatedUser)
	if err != nil {
		log.Printf("Could not save User: %v", err)
		return nil, err
	}
	return updatedUser, nil
}
