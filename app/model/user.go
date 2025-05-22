package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type User struct {
	Id        string    `bson:"_id,omitempty"` // Changed from uint to string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(username, email, password string, created, updated time.Time) User {
	return User{
		Username: username,
		Email: email,
		Password: password,
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