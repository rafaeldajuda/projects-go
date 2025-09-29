package db

import (
	"context"
	"encoding/json"
	"fmt"
	"main/entity"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Insert(task entity.Task, mongodb *mongo.Client) (err error) {
	ctx := context.Background()
	collection := mongodb.Database("dev").Collection("tasks")
	_, err = collection.InsertOne(ctx, task)
	if err != nil {
		return fmt.Errorf("erro ao inserir task: %s", err.Error())
	}

	return nil
}

func UpdateOne(id string, task entity.Task, mongodb *mongo.Client) (err error) {
	ctx := context.Background()
	collection := mongodb.Database("dev").Collection("tasks")
	filter := bson.M{"id": id}
	result, err := collection.Find(ctx, filter)
	if err != nil {
		return fmt.Errorf("erro ao buscar task: %s", err.Error())
	}

	if result.Err() != nil {
		return fmt.Errorf("erro ao buscar task com id %s: %s", id, result.Err().Error())
	}

	newData := bson.M{
		"$set": task,
	}
	resultUpdate, err := collection.UpdateOne(ctx, filter, newData)
	if err != nil {
		return fmt.Errorf("erro ao atualizar task: %s", err.Error())
	}

	if resultUpdate.ModifiedCount != 1 {
		return fmt.Errorf("erro ao atualizar task")
	}

	return nil
}

func FindAll(mongodb *mongo.Client) (tasks []entity.Task, err error) {
	ctx := context.Background()
	collection := mongodb.Database("dev").Collection("tasks")
	result, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar tasks: %s", err.Error())
	}

	for result.Next(ctx) {
		task := entity.Task{}
		err = json.Unmarshal([]byte(result.Current.String()), &task)
		if err != nil {
			return nil, fmt.Errorf("erro ao iterar tasks: %s", err.Error())
		}

		tasks = append(tasks, task)
	}

	return
}

func FindOne(id string, mongodb *mongo.Client) (task entity.Task, err error) {
	ctx := context.Background()
	collection := mongodb.Database("dev").Collection("tasks")
	result := collection.FindOne(ctx, bson.D{{Key: "id", Value: id}})

	if result.Err() != nil {
		return task, fmt.Errorf("erro ao buscar task: %s", result.Err().Error())
	}

	err = result.Decode(&task)
	if err != nil {
		return task, fmt.Errorf("erro converter resposta: %s", err.Error())
	}

	return
}

func DeleteOne(id string, mongodb *mongo.Client) (err error) {
	ctx := context.Background()
	collection := mongodb.Database("dev").Collection("tasks")
	// result := collection.FindOne(ctx, bson.D{{Key: "id", Value: id}})

	// if result.Err() != nil {
	// 	return task, fmt.Errorf("erro ao buscar task: %s", result.Err().Error())
	// }

	result, err := collection.DeleteOne(ctx, bson.D{{Key: "id", Value: id}})
	if err != nil {
		return fmt.Errorf("erro ao deletar task: %s", err.Error())
	}

	if result.DeletedCount != 1 {
		return fmt.Errorf("erro ao deletar task: id não existe")
	}

	return
}
