package database

import (
	"context"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	once     sync.Once
	instance *mongo.Database
	err      error
)

func GetClient() (*mongo.Database, error) {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			beego.Error(err)
			return
		}

		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err = client.Ping(ctx, readpref.Primary()) // check if MongoDB server found
		if err != nil {
			beego.Error(err)
			return
		}

		beego.Info("[database/client] database found")
		instance = client.Database("test")
	})

	return instance, err
}
