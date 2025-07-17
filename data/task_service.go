package data

import (
	"context"
	"errors"

	"task_management/db"
	"task_management/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task

	cur, err := db.TaskCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err

	}
	if err = cur.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil

}

func GetTaskByID(id string) (*models.Task, error) {
	objectID,err:= primitive.ObjectIDFromHex(id)

	if err!=nil{
		return nil,errors.New("invalid task Id")
	}
	var task models.Task
	err=db.TaskCollection.FindOne(context.TODO(),bson.M{"_id":objectID}).Decode(&task)
	if err != nil{
		return nil,err
	}
	return &task ,nil

}

func AddTask(task models.Task) error {
	_, err := db.TaskCollection.InsertOne(context.TODO(), task)
	return err
}


// func DeleteTaskByID(id int) bool {
// 	for i, task := range Tasks {
// 		if task.ID == id {
// 			Tasks = append(Tasks[:i], Tasks[i+1:]...)
// 			return true
// 		}

// 	}
// 	return false
// }

// func UpdateTaskByID(id int, updatedTask models.Task) bool {
// 	for i, task := range Tasks {
// 		if task.ID == id {
// 			updatedTask.ID = id
// 			Tasks[i] = updatedTask
// 			return true

// 		}
// 	}
// 	return false
// }
