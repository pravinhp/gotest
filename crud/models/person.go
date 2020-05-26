package models

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/pravinhp/gotest/crud/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBPerson struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstname,omitempty"`
	Lastname  string             `bson:"lastname,omitempty"`
}

type Person struct {
	FirstName string
	Lastname  string
}

type ObjectID [12]byte

func (id ObjectID) Hex() string {
	return hex.EncodeToString(id[:])
}

func New(firstname, lastname string) *Person {
	person := &Person{
		FirstName: firstname,
		Lastname:  lastname,
	}

	return person
}

func GetPerson() *[]DBPerson {

	var person []DBPerson
	collection := db.Connect()
	data, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}

	if err = data.All(context.TODO(), &person); err != nil {
		panic(err)
	}

	return &person
}

func (p *Person) CreatePerson() interface{} {

	collection := db.Connect()

	res, err := collection.InsertOne(context.TODO(), p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in create person %T", res.InsertedID)
	return res.InsertedID
}

func GetPersonById(id string) DBPerson {

	collection := db.Connect()
	var person DBPerson

	objID, _ := primitive.ObjectIDFromHex(id)
	fmt.Println("in person ---", objID)
	err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&person)

	fmt.Println(person)
	if err != nil {
		panic(err)
	}

	return person
}
