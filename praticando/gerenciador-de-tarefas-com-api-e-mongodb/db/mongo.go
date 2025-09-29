package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// docker run -d -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin -p 27017:27017 --name my-mongo mongo:latest
// docker exec -it my-mongo mongosh -u admin -p admin

// create database - use dev
// create collection - db.createCollection("tasks")
// insert - db.tasks.insertOne({"id":"1","title":"Estudar Go","description":"Praticar structs, interfaces e goroutines","completed":false})
// select - db.tasks.find()

func Connect() (client *mongo.Client) {

	// MONGODB_URI=mongodb+srv://<db_username>:<db_password>@<cluster-url>?retryWrites=true&w=majority
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/",
		"admin", "admin", "0.0.0.0", "27017")

	// Uses the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// Defines the options for the MongoDB client
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// Sends a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("dev").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return

}

func ConfigDB(client *mongo.Client) error {
	ctx := context.Background()
	err := client.Database("dev").CreateCollection(ctx, "tasks")
	if err != nil {
		return fmt.Errorf("erro criar coleção task: %s", err.Error())
	}

	task := bson.D{
		{Key: "id", Value: "999"},
		{Key: "title", Value: "teste"},
		{Key: "description", Value: "testando mongo"},
		{Key: "completed", Value: false},
	}
	collection := client.Database("dev").Collection("tasks")
	if collection.FindOne(ctx, bson.D{{Key: "id", Value: "999"}}).Err() != nil {
		_, err = collection.InsertOne(ctx, task)
		if err != nil {
			return fmt.Errorf("erro ao inserir default task: %s", err.Error())
		}
	}

	return nil
}
