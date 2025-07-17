package models
import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskStatus string

const (
	StatusNotStarted TaskStatus = "not-started"
	StatusInProgress TaskStatus = "in-progress"
	StatusCompleted  TaskStatus = "completed"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	DueDate     time.Time          `bson:"dueDate" json:"dueDate"`
	Status      TaskStatus         `bson:"status" json:"status"`
	
}