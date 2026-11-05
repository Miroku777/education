package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//connection
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:secretpassword@localhost:27017/?authSource=admin"))
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println(err)
			return
		}
	}()
	cursor := client.Database("users").Collection("users")

	_, err = cursor.InsertOne(context.TODO(), bson.M{"name": "Serega", "work": "CPO"})
	if err != nil {
		log.Println(err)
		return
	}

	var Serega map[string]any
	err = cursor.FindOne(context.TODO(), bson.M{"name": "Serega"}).Decode(&Serega)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No such document")
		} else {
			log.Println(err)
			return
		}
	} else {
		fmt.Println(Serega)
	}

	cur, err := cursor.Find(context.TODO(), bson.M{"name": "Serega"})
	if err != nil {
		log.Println(err)
		return
	}
	defer cur.Close(context.TODO())
	var res []map[string]any
	err = cur.All(context.TODO(), &res)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
