package models

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/vimalkuriensam/item-service/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ItemModel interface {
	Connect() error
	Disconnect()
	InsertMongoCollections(...string)
}

type itemModel struct {
	config *config.Config
}

func Init() ItemModel {
	return &itemModel{
		config: config.GetConfig(),
	}
}

func (m *itemModel) Connect() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	user := m.config.Env["db_user"].(string)
	pass := m.config.Env["db_password"].(string)
	host := m.config.Env["dsn"].(string)
	dsn := fmt.Sprintf("mongodb://%v:%v@%v/?maxPoolSize=20&w=majority", user, pass, host)
	var (
		attempts int
		err      error
		client   *mongo.Client
	)
	for {
		attempts++
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(dsn))
		if err != nil {
			m.config.Logger.Printf("Db connection failed. Attempt %d...\n", attempts)
			if attempts > 10 {
				m.config.Logger.Printf("Exceeded maximum number of attempts of %d\n", attempts)
				return err
			}
			time.Sleep(time.Duration(math.Pow(float64(attempts), 2)) * time.Second)
			continue
		}
		m.config.Logger.Printf("Db connected in %d attempts\n", attempts)
		m.config.Database.Client = client
		break
	}
	return nil
}

func (m *itemModel) Disconnect() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	if err := m.config.Database.Client.Disconnect(ctx); err != nil {
		m.config.Logger.Fatal(err)
	}
}

func (m *itemModel) InsertMongoCollections(collections ...string) {
	for _, value := range collections {
		go InsertMongoCollection(value, m.config.Database, m.config.Env["db_database"].(string))
	}
}

func InsertMongoCollection(collection string, db *config.Database, dbName string) {
	col := db.Client.Database(dbName).Collection(collection)
	db.Collections[collection] = col
	fmt.Printf("Added collection %s\n", collection)
}
