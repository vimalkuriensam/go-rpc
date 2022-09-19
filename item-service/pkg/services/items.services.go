package services

import (
	"context"
	"time"

	"github.com/vimalkuriensam/item-service/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemService interface {
	InsertItem(models.Items) (*mongo.InsertOneResult, error)
	GetItem(string) *mongo.SingleResult
	UpdateItem(string, models.Items) (*mongo.UpdateResult, error)
	DeleteItem(string) (*mongo.DeleteResult, error)
}

type itemService struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) ItemService {
	return &itemService{
		collection: collection,
	}
}

func (s *itemService) InsertItem(item models.Items) (*mongo.InsertOneResult, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	newItem := bson.D{
		{Key: "name", Value: item.Name},
		{Key: "value", Value: item.Value},
		{Key: "created_at", Value: time.Now()},
		{Key: "updated_at", Value: time.Now()},
	}
	return s.collection.InsertOne(ctx, newItem)
}

func (s *itemService) GetItem(id string) *mongo.SingleResult {
	return nil
}

func (s *itemService) UpdateItem(id string, item models.Items) (*mongo.UpdateResult, error) {
	return nil, nil
}

func (s *itemService) DeleteItem(id string) (*mongo.DeleteResult, error) {
	return nil, nil
}
