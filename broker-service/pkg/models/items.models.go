package models

import (
	"time"
)

var ItemAcceptableFields = map[string][]string{
	"create": {"name", "value"},
	"update": {"name", "value"},
}

type Items struct {
	StringID  string    `json:"id"`
	Name      string    `json:"name"`
	Value     int       `json:"value"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
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

type DeleteItem struct {
	DeleteCount int   `json:"deleteCount"`
	Item        Items `json:"item"`
}
