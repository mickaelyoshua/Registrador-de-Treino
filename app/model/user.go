package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type User struct {
	Id uint `bson:",omitempty"`
	Username string
	Email string
	Password string
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

func (u User) Save(client *mongo.Client) error {
	coll := client.Database("workout_register").Collection("user")
	writeResult, err := coll.InsertOne(context.TODO(), u)

	writeJsonData, _ := json.MarshalIndent(writeResult, "", "	")
	fmt.Printf("%s\n", writeJsonData)
	fmt.Println("Write result:", writeResult)

	return err
}