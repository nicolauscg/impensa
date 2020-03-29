package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/nicolauscg/impensa/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Handler struct {
	db   *mongo.Database
	Orms *Entity
}

type Entity struct {
	User        models.UserOrmer
	Transaction models.TransactionOrmer
}

func NewHandler(databaseName string, connString string) (handler *Handler, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		err = errors.New(fmt.Sprintf("[handler/handler] mongo.Connect error. %+v", err))
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// check if MongoDB server found
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		err = errors.New(fmt.Sprintf("[handler/handler] mongo client ping error. %+v", err))
		return
	}

	handler = &Handler{db: client.Database(databaseName)}
	handler.Orms = &Entity{
		models.NewUserOrm(handler.db),
		models.NewTransactionOrm(handler.db),
	}

	if handler == nil {
		err = errors.New("[handler/handler] database not found. %+v")
	} else {
		beego.Info("[handler/handler] database successfully connected.")
	}

	return handler, err
}
