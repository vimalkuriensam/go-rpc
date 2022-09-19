package rpc

import (
	"github.com/vimalkuriensam/item-service/pkg/models"
	"github.com/vimalkuriensam/item-service/pkg/services"
)

type ItemService interface {
	AddItem(models.Items, *models.Items) error
	GetItem()
	UpdateItem()
	DeleteItem()
}

type ItemCollection struct {
	services services.ItemService
}

func New(services services.ItemService) ItemService {
	return &ItemCollection{
		services: services,
	}
}

func (c *ItemCollection) AddItem(item models.Items, result *models.Items) error {
	return nil
}

func (c *ItemCollection) GetItem() {}

func (c *ItemCollection) UpdateItem() {}

func (c *ItemCollection) DeleteItem() {}
