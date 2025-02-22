package models

import (
	"context"
	"errors"

	"github.com/mickaelyoshua7674/Registrador-de-Treino/db"
	"github.com/mickaelyoshua7674/Registrador-de-Treino/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type User struct {
	Id       int `bson:"_id,omitempty"`
	Email    string `bson:"email,required"`
	Password string `bson:"password,required"`
}

func (u *User) Save() (primitive.ObjectID, error) {
	if u.EmailExists() {
		return primitive.NilObjectID, errors.New("email already exists")
	}

	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return primitive.NilObjectID, err
	}
	u.Password = hashedPass

	client, err := db.GetClient()
	if err != nil {
		return primitive.NilObjectID, err
	}

	coll := client.Database("workout_register").Collection("user")
	result, err := coll.InsertOne(context.TODO(), u)
	if err != mongo.ErrNilCursor {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (u User) EmailExists() bool {
	client, _ := db.GetClient()
	coll := client.Database("workout_register").Collection("user")

	var result bson.M
	filter := bson.D{{Key: "email", Value: u.Email}}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	return err != mongo.ErrNoDocuments
}

func (u User) ValidateCredentials() error {

}