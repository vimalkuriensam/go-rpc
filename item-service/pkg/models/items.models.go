package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Items struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Value     int                `json:"value" bson:"value"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updated_at"`
	StringID  string             `json:"stringId"`
}

type ItemRPC struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	Value     int       `json:"value" bson:"value"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
}

type DeleteItem struct {
	DeleteCount int   `json:"deleteCount"`
	Item        Items `json:"item"`
}

type UpdateItemInput struct {
	ID         string
	UpdateItem Items
}

type UpdateItemResponse struct {
	ID          string `json:"id"`
	Count       int    `json:"count"`
	PriorItem   Items  `json:"priorItem"`
	UpdatedItem Items  `json:"updatedItem"`
}
