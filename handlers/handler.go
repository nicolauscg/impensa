package handler

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/astaxie/beego"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/nicolauscg/impensa/constants"
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
	User              models.UserOrmer
	Transaction       models.TransactionOrmer
	Account           models.AccountOrmer
	Category          models.CategoryOrmer
	VerifyAccount     models.VerifyUserOrmer
	ResetUserPassword models.ResetUserPasswordOrmer
	MailGun           models.MailOrmer
}

func NewHandler(databaseName string, connString string) (handler *Handler, err error) {
	connStringNoCred := regexp.MustCompile(`://.*@`).ReplaceAllString(connString, `://<username>:<password>@`)
	beego.Info(fmt.Sprintf("creating handler with DBName %v and ConnString %v", databaseName, connStringNoCred))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		err = errors.New(fmt.Sprintf("[handler/handler] mongo.Connect error. %+v", err))
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// check if MongoDB server found
	for waitInterval := 5; ; {
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			beego.Error(fmt.Sprintf("[handler/handler] mongo client ping error. %+v", err))
			beego.Error(fmt.Sprintf("retrying in %v seconds..", waitInterval))
			time.Sleep(time.Duration(waitInterval) * time.Second)
			waitInterval += waitInterval / 2
		} else {
			beego.Info("ping success")
			break
		}
	}

	handler = &Handler{db: client.Database(databaseName)}
	handler.Orms = &Entity{
		models.NewUserOrm(handler.db),
		models.NewTransactionOrm(handler.db),
		models.NewAccountOrm(handler.db),
		models.NewCategoryOrm(handler.db),
		models.NewVerifyUserOrm(handler.db),
		models.NewResetUserPassword(handler.db),
		models.NewMailOrmer(mailgun.NewMailgun("mail.impensa.nicolauscg.me", os.Getenv(constants.EnvMailgunApi))),
	}

	if handler == nil {
		err = errors.New("[handler/handler] database not found. %+v")
	} else {
		beego.Info("[handler/handler] database successfully connected.")
	}

	return handler, err
}
