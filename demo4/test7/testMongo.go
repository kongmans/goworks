package test7

//func main() {
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	client, err := mongo.Connect(ctx, &options.ClientOptions{Hosts: []string{"localhost:27017"}})
//	if err != nil {
//		log.Fatal(err)
//
//	}
//	// Check the connection
//	err = client.Ping(context.TODO(), nil)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("Connected to MongoDB!")
//
//}

//package main
//
//import (
//	"context"
//	"fmt"
//	"github.com/mongodb/mongo-go-driver/mongo"
//	"github.com/mongodb/mongo-go-driver/mongo/options"
//	"log"
//	"time"
//)
//
//type Trainer struct {
//	Name string
//	Age  int
//	City string
//}
//
//func main() {
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	client, err := mongo.Connect(ctx, &options.ClientOptions{Hosts: []string{"localhost:27017"}})
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Check the connection
//	err = client.Ping(context.TODO(), nil)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("Connected to MongoDB!")
//
//	collection := client.Database("test").Collection("trainers")
//	_ = collection
//
//	err = client.Disconnect(context.TODO())
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Connection to MongoDB closed.")
//
//}
