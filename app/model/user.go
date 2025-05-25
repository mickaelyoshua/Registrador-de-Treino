package model

import (
	"context"
	"time"

	"github.com/mickaelyoshua/Registrador-de-Treino/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func FindUserByToken(retrievedToken map[string]any) (User, error) {
	client, err := db.GetClient()
	if err != nil {
		return User{}, err
	}

	objectID, err := primitive.ObjectIDFromHex(retrievedToken["id"].(string))
	if err != nil {
		return User{}, err
	}
	filter := bson.M{"_id": objectID}
	coll := client.Database("workout_register").Collection("user")
	var user User
	err = coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserByFilter(filter bson.M) (User, error) {
	client, err := db.GetClient()
	if err != nil {
		return User{}, err
	}

	coll := client.Database("workout_register").Collection("user")
	var user User
	err = coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u User) Save() error {
	client, err := db.GetClient()
	if err != nil {
		return err
	}

	coll := client.Database("workout_register").Collection("user")
	_, err = coll.InsertOne(context.TODO(), u)
	return err
}