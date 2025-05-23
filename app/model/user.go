package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"` // Use ObjectID for MongoDB's _id field
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func NewUser(username, email, password string, created, updated time.Time) User {
	return User{
		Id:        primitive.NewObjectID(), // Generate a new ObjectID
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: created,
		UpdatedAt: updated,
	}
}

func FindUser(client *mongo.Client, filter any) (User, error) {
	coll := client.Database("workout_register").Collection("user")
	var user User
	err := coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u User) Save(client *mongo.Client) error {
	coll := client.Database("workout_register").Collection("user")
	_, err := coll.InsertOne(context.TODO(), u)
	return err
}