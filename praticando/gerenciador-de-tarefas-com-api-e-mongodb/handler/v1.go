package handler

import (
	"encoding/json"
	"errors"
	"main/db"
	"main/entity"
	"math/rand/v2"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func PostTask(body []byte, mongdb *mongo.Client) (id string, err error) {
	task := entity.Task{}
	err = json.Unmarshal(body, &task)
	if err != nil {
		return "", err
	}

	// validar campos
	err = validateFields(task)
	if err != nil {
		return "", err
	}

	// gerar id
	id = strconv.FormatInt(int64(rand.Int32()), 10)
	task.ID = id

	// guardar no banco
	err = db.Insert(task, mongdb)
	if err != nil {
		return "", err
	}

	return
}

func GetTasks(mongdb *mongo.Client) (tasks []entity.Task, err error) {

	// pegar registros do banco
	tasks, err = db.FindAll(mongdb)
	if err != nil {
		return
	}

	return
}

func GetTask(id string, mongdb *mongo.Client) (task entity.Task, err error) {

	// pegar registro do banco
	task, err = db.FindOne(id, mongdb)
	if err != nil {
		return
	}

	return
}

func PutTask(id string, body []byte, mongdb *mongo.Client) (err error) {
	task := entity.Task{}
	err = json.Unmarshal(body, &task)
	if err != nil {
		return err
	}
	task.ID = id

	// validar campos
	err = validateFields(task)
	if err != nil {
		return err
	}

	// atualizar no banco
	err = db.UpdateOne(id, task, mongdb)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(id string, mongdb *mongo.Client) (err error) {

	// deletar registro do banco
	err = db.DeleteOne(id, mongdb)
	if err != nil {
		return
	}

	return
}

func validateFields(task entity.Task) error {
	if task.Title == "" {
		return errors.New("campo title vazio")
	} else if task.Description == "" {
		return errors.New("campo description vazio")
	} else if task.Completed == nil {
		return errors.New("campo completed vazio")
	}

	return nil
}
