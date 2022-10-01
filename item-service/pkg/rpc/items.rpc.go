package rpc

import (
	"bytes"
	"fmt"

	"github.com/vimalkuriensam/item-service/pkg/config"
	"github.com/vimalkuriensam/item-service/pkg/models"
	"github.com/vimalkuriensam/item-service/pkg/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type DataToSend struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (c *ItemCollection) AddItem(item models.Items, result *config.JSONResponse) error {
	var insertedItem models.Items
	var b bytes.Buffer
	id, err := c.services.InsertItemCollection(item)
	if err != nil {
		return err
	}
	stringId := id.InsertedID.(primitive.ObjectID).Hex()
	result.Message = fmt.Sprintf("Item with id %v created", stringId)
	if err = c.services.GetItemCollection(stringId).Decode(&insertedItem); err != nil {
		return err
	}
	insertedItem.StringID = stringId
	if err := services.EncodeData(insertedItem, &b); err != nil {
		return err
	}
	result.Data = b.Bytes()
	return nil
}

func (c *ItemCollection) GetItem(id string, result *config.JSONResponse) error {
	item := models.Items{}
	var b bytes.Buffer
	err := c.services.GetItemCollection(id).Decode(&item)
	if err != nil {
		return err
	}
	result.Message = fmt.Sprintf("item with id %v fetched.", id)
	item.StringID = id
	if err := services.EncodeData(item, &b); err != nil {
		return err
	}
	result.Data = b.Bytes()
	return nil
}

func (c *ItemCollection) UpdateItem(updates models.UpdateItemInput, result *config.JSONResponse) error {
	var (
		priorItem   models.Items
		updatedItem models.Items
		b           bytes.Buffer
	)
	if err := c.services.GetItemCollection(updates.ID).Decode(&priorItem); err != nil {
		return err
	}
	response, err := c.services.UpdateItemCollection(updates.ID, updates.UpdateItem)
	if err != nil {
		return err
	}
	if err := c.services.GetItemCollection(updates.ID).Decode(&updatedItem); err != nil {
		return err
	}
	result.Message = fmt.Sprintf("Item with id %v updated successfully", updates.ID)
	priorItem.StringID = updates.ID
	updatedItem.StringID = updates.ID
	result.Data = models.UpdateItemResponse{
		ID:          updates.ID,
		Count:       int(response.ModifiedCount),
		PriorItem:   priorItem,
		UpdatedItem: updatedItem,
	}
	if err := services.EncodeData(result.Data, &b); err != nil {
		return err
	}
	result.Data = b.Bytes()
	return nil
}

func (c *ItemCollection) DeleteItem(id string, result *config.JSONResponse) error {
	var (
		item models.Items
		b    bytes.Buffer
	)
	err := c.services.GetItemCollection(id).Decode(&item)
	if err != nil {
		return err
	}
	deleteResp, err := c.services.DeleteItemCollection(id)
	if err != nil {
		return err
	}
	result.Message = fmt.Sprintf("Item with id %v deleted successfully", id)
	item.StringID = id
	result.Data = models.DeleteItem{
		DeleteCount: int(deleteResp.DeletedCount),
		Item:        item,
	}
	if err := services.EncodeData(result.Data, &b); err != nil {
		return err
	}
	result.Data = b.Bytes()
	return nil
}
