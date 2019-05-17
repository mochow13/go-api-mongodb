package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Question model of mongodb
type Question struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	Question        string             `json:"question" bson:"question"`
	Options         []string           `json:"options" bson:"options"`
	Answer          string             `json:"answer" bson:"answer"`
	Hint            string             `json:"hint" bson:"hint"`
	ImageLink       string             `json:"image_link" bson:"image_link"`
	Subject         string             `json:"subject" bson:"subject"`
	Topic           string             `json:"topic" bson:"topic"`
	Tags            []string           `json:"tags" bson:"tags"`
	DifficultyLevel string             `json:"difficulty_level" bson:"difficulty_level"`
	Source          string             `json:"source" bson:"source"`
	Class           string             `json:"class" bson:"class"`
	Category        string             `json:"category" bson:"category"`
	Type            string             `json:"type" bson:"type"`
	Sequence        int                `json:"sequence" bson:"sequence"`
}
