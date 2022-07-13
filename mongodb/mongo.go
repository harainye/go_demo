package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//定义插入数据的结构体
type sunshareboy struct {
	Name string
	Age  int
	City string
}

type mongoCourse struct {
	Title       string
	Description string
	By          string
	Url         string
	Tags        []string
	Likes       int
}

func ConnMongo() {
	clientOptions := options.Client().ApplyURI("mongodb://192.168.9.75:27017")
	var ctx = context.TODO()
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	defer client.Disconnect(ctx)

	/*//连接到test库的sunshare集合,集合不存在会自动创建
	collection := client.Database("test").Collection("sunshare")
	wanger := sunshareboy{"harainye", 30, "福州"}
	insertOne, err := collection.InsertOne(ctx, wanger)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)*/

	collection := client.Database("test").Collection("test")

	var result mongoCourse
	filter := bson.D{{"title", "MongoDB 教程"}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
}
