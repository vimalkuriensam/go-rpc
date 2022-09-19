package rpc

import (
	"fmt"

	"github.com/vimalkuriensam/item-service/pkg/config"
	"github.com/vimalkuriensam/item-service/pkg/models"
	"github.com/vimalkuriensam/item-service/pkg/services"
)

type ItemService interface {
	AddItem(models.Items, *config.JSONResponse) error
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

func (c *ItemCollection) AddItem(item models.Items, result *config.JSONResponse) error {
	id, err := c.services.InsertItem(item)
	if err != nil {
		return err
	}
	result.Message = fmt.Sprintf("Item with id %v created", id.InsertedID)
	result.Data = item
	return nil
}

func (c *ItemCollection) GetItem() {}

func (c *ItemCollection) UpdateItem() {}

func (c *ItemCollection) DeleteItem() {}
