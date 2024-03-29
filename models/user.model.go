package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	_Id      primitive.ObjectID `json:"_id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" validate:"required"`
	Title    string             `json:"title,omitempty" validate:"required"`
}

type EditUser struct {
	_Id      primitive.ObjectID `json:"_id,omitempty"`
	Name     string             `json:"name,omitempty"`
	Location string             `json:"location,omitempty"`
	Title    string             `json:"title,omitempty"`
}
