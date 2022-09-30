package models

import (
	"time"
)

var ItemAcceptableFields = map[string][]string{
	"create": {"item", "value"},
	"update": {"item", "value"},
}

type Items struct {
	StringID  string    `json:"id"`
	Name      string    `json:"name"`
	Value     int       `json:"value"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
