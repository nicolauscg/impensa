package handler

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/nicolauscg/impensa/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	once    sync.Once
	handler *Handler
	err     error
)

type Handler struct {
	db   *mongo.Database
	Orms *Entity
}

type Entity struct {
	Transaction models.TransactionOrmer
}

func NewHandler(databaseName string) (*Handler, error) {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			err = errors.New(fmt.Sprintf("[handler/handler] mongo.Connect error. %+v", err))
			return
		}

		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		// check if MongoDB server found
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			err = errors.New(fmt.Sprintf("[handler/handler] mongo client ping error. %+v", err))
			return
		}

		handler = &Handler{db: client.Database(databaseName)}
		handler.Orms = &Entity{models.NewTransactionOrm(handler.db)}
		beego.Info("[handler/handler] database successfully connected.")
	})

	return handler, err
}
