package services

import (
	"context"
	"time"

	"github.com/vimalkuriensam/item-service/pkg/config"
	"github.com/vimalkuriensam/item-service/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemService interface {
	InsertItemCollection(models.Items) (*mongo.InsertOneResult, error)
	GetItemCollection(string) *mongo.SingleResult
	UpdateItemCollection(string, models.Items) (*mongo.UpdateResult, error)
	DeleteItemCollection(string) (*mongo.DeleteResult, error)
}

type itemService struct{}

func New() ItemService {
	return &itemService{}
}

func (s *itemService) InsertItemCollection(item models.Items) (*mongo.InsertOneResult, error) {
	collection := config.GetConfig().Database.Collections["items"]
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	newItem := bson.D{
		{Key: "name", Value: item.Name},
		{Key: "value", Value: item.Value},
		{Key: "created_at", Value: time.Now()},
		{Key: "updated_at", Value: time.Now()},
	}
	return collection.InsertOne(ctx, newItem)
}

func (s *itemService) GetItemCollection(id string) *mongo.SingleResult {
	collection := config.GetConfig().Database.Collections["items"]
	docId, _ := primitive.ObjectIDFromHex(id)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	return collection.FindOne(ctx, bson.M{"_id": docId})
}

func (s *itemService) UpdateItemCollection(id string, item models.Items) (*mongo.UpdateResult, error) {
	collection := config.GetConfig().Database.Collections["items"]
	docId, _ := primitive.ObjectIDFromHex(id)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	return collection.UpdateByID(ctx, docId, bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "name", Value: item.Name},
			primitive.E{Key: "value", Value: item.Value},
			primitive.E{Key: "update_at", Value: time.Now()},
		}},
	})
}

func (s *itemService) DeleteItemCollection(id string) (*mongo.DeleteResult, error) {
	collection := config.GetConfig().Database.Collections["items"]
	docId, _ := primitive.ObjectIDFromHex(id)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	return collection.DeleteOne(ctx, bson.M{"_id": docId})
}
