package services

import (
	"github.com/vimalkuriensam/item-service/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemService interface {
	AddItem(models.Items, *models.Items) error
	GetItem()
	UpdateItem()
	DeleteItem()
}

type ItemCollection struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) ItemService {
	return &ItemCollection{
		collection: collection,
	}
}

func (c *ItemCollection) AddItem(item models.Items, result *models.Items) error {
	return nil
}

func (c *ItemCollection) GetItem() {}

func (c *ItemCollection) UpdateItem() {}

func (c *ItemCollection) DeleteItem() {}
