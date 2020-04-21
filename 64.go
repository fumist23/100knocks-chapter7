package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Artist struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Tags []struct {
		Count int    `json:"count" bson:"count"`
		Value string `json:"value" bson:"value"`
	} `json:"tags" bson:"tags"`
	SortName string `json:"sort_name" bson:"sort_name"`
	Gid      string `json:"gid" bson:"gid"`
	Area     string `json:"area" bson:"area"`
	Aliases  []struct {
		Name     string `json:"name" bson:"name"`
		SortName string `json:"sort_name" bson:"sort_name"`
	} `json:"aliases" bson:"aliases"`
	Begin struct {
		Year  int `json:"year" bson:"year"`
		Month int `json:"month" bson:"month"`
		Date  int `json:"date" bson:"date"`
	} `json:"begin" bson:"begin"`
	End struct {
		Year  int `json:"year" bson:"year"`
		Month int `json:"month" bson:"month"`
		Date  int `json:"date" bson:"date"`
	} `json:"end" bson:"end"`
	Rating struct {
		Count int `json:"count" bson:"count"`
		Value int `json:"value" bson:"value"`
	} `json:"rating" bson:"rating"`
}

func main() {
	filename := "artist.json"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection := client.Database("MBDB").Collection("artist")

	for i := 0; i < 100; i++ {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		var artist Artist
		json.Unmarshal([]byte(line), &artist)
		_, err = collection.InsertOne(context.Background(), artist)
		if err != nil {
			panic(err)
		}
	}
	var result interface{}
	filter := bson.M{"name": "supercell"}
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
