package model

import (
	"time"
	"context"

	"github.com/mickaelyoshua/Registrador-de-Treino/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workout struct {
	Id        primitive.ObjectID `bson:"_id"` // Use ObjectID for MongoDB's _id field
	Title     string             `bson:"title"`
	Description string             `bson:"description"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	UserId    primitive.ObjectID `bson:"user_id"` // Reference to the user who created the workout
}

func NewWorkout(title, description string, userId primitive.ObjectID, created, updated time.Time) Workout {
	return Workout{
		Id:          primitive.NewObjectID(), // Generate a new ObjectID
		Title:       title,
		Description: description,
		CreatedAt:   created,
		UpdatedAt:   updated,
		UserId:      userId,
	}
}

func FindAllWorkoutsByUserId(userId primitive.ObjectID) ([]Workout, error) {
	client, err := db.GetClient()
	if err != nil {
		return nil, err
	}
	coll := client.Database("workout_register").Collection("workout")
	filter := bson.M{"user_id": userId}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var workouts []Workout
	for cursor.Next(context.TODO()) {
		var workout Workout
		if err := cursor.Decode(&workout); err != nil {
			return nil, err
		}
		workouts = append(workouts, workout)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// If no workouts found, return an empty slice
	if len(workouts) == 0 {
		return []Workout{}, nil
	}
	return workouts, nil
}

func GetWorkoutById(id string) (Workout, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Workout{}, err
	}

	client, err := db.GetClient()
	if err != nil {
		return Workout{}, err
	}

	coll := client.Database("workout_register").Collection("workout")
	filter := bson.M{"_id": objectId}
	var workout Workout
	err = coll.FindOne(context.TODO(), filter).Decode(&workout)
	if err != nil {
		return Workout{}, err
	}
	return workout, nil
}

func (w Workout) Save() error {
	client, err := db.GetClient()
	if err != nil {
		return err
	}
	coll := client.Database("workout_register").Collection("workout")
	_, err = coll.InsertOne(context.TODO(), w)
	return err
}

func (w Workout) Delete() error {
	client, err := db.GetClient()
	if err != nil {
		return err
	}
	coll := client.Database("workout_register").Collection("workout")
	_, err = coll.DeleteOne(context.TODO(), bson.M{"_id": w.Id})
	return err
}