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

type itemService struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) ItemService {
	return &itemService{
		collection: collection,
	}
}

func (item *itemService) AddItem(i models.Items, result *models.Items) error {
	return nil
}

func (item *itemService) GetItem() {}

func (item *itemService) UpdateItem() {}

func (item *itemService) DeleteItem() {}
