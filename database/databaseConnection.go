package database

import (
	"log"
	"os"
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client{
	err := godotenv.Load("configs/.env")
	if(err!=nil){
		log.Fatalln("error in loading .env",err)
	} 

	x := os.Getenv("MONGO_URI")
	uri := x + "hrms"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// creating new client and connecting it to server
	client, err := mongo.Connect(context.TODO(),opts)
	if(err!=nil){
		log.Fatalln("error in connecting uri to mongo",err)
	}
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }()

	// this code is to check if connection is succesfull or not
	// var result bson.M
	// if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
	// 	log.Fatalln(err)
	// }
	fmt.Println("connected to mongoDB!")
	return client
}

// this is client which contains all info about the database
var Client *mongo.Client = DBinstance()

// in mongobd collections are like tables in mysql

// this function helps to return particular collection from database
func OpenCollection(collectionName string) *mongo.Collection{
	collection := Client.Database("hrms").Collection(collectionName)
	return collection
}