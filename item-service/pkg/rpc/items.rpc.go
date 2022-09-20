package rpc

import (
	"fmt"
	"time"

	"github.com/vimalkuriensam/item-service/pkg/config"
	"github.com/vimalkuriensam/item-service/pkg/models"
	"github.com/vimalkuriensam/item-service/pkg/services"
)

type ItemService interface {
	AddItem(models.Items, *config.JSONResponse) error
	GetItem(string, *config.JSONResponse) error
	UpdateItem(models.UpdateItemInput, *config.JSONResponse) error
	DeleteItem(string, *config.JSONResponse) error
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
	id, err := c.services.InsertItemCollection(item)
	if err != nil {
		return err
	}
	result.Message = fmt.Sprintf("Item with id %v created", id.InsertedID)
	result.Data = item
	return nil
}

func (c *ItemCollection) GetItem(id string, result *config.JSONResponse) error {
	var item models.ItemModel
	err := c.services.GetItemCollection(id).Decode(&item)
	if err != nil {
		return err
	}
	result.Message = fmt.Sprintf("item with id %v fetched.", id)
	result.Data = item
	return nil
}

func (c *ItemCollection) UpdateItem(updates models.UpdateItemInput, result *config.JSONResponse) error {
	var (
		priorItem   models.Items
		updatedItem models.Items
		itemResp    *config.JSONResponse
	)
	err := c.GetItem(updates.ID, itemResp)
	if err != nil {
		return err
	}
	priorItem = itemResp.Data.(models.Items)
	updates.UpdateItem.UpdatedAt = time.Now()
	response, err := c.services.UpdateItemCollection(updates.ID, updates.UpdateItem)
	if err != nil {
		return err
	}
	err = c.GetItem(updates.ID, itemResp)
	if err != nil {
		return err
	}
	updatedItem = itemResp.Data.(models.Items)
	result.Message = fmt.Sprintf("Item with id %v updated successfully", updates.ID)
	result.Data = models.UpdateItemResponse{
		ID:          updates.ID,
		Count:       int(response.ModifiedCount),
		PriorItem:   priorItem,
		UpdatedItem: updatedItem,
	}
	return nil
}

func (c *ItemCollection) DeleteItem(id string, result *config.JSONResponse) error {
	var (
		itemResp *config.JSONResponse
		item     models.Items
	)
	err := c.GetItem(id, itemResp)
	if err != nil {
		return err
	}
	item = itemResp.Data.(models.Items)
	deleteResp, err := c.services.DeleteItemCollection(id)
	if err != nil {
		return err
	}
	result.Message = fmt.Sprintf("Item with id %v deleted successfully", id)
	result.Data = models.DeleteItem{
		DeleteCount: int(deleteResp.DeletedCount),
		Item:        item,
	}
	return nil
}
