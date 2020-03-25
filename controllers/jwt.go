package controllers

import (
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ControllerWithUserId struct {
	beego.Controller
	UserId primitive.ObjectID
}

func (a *ControllerWithUserId) Prepare() {
	a.UserId, _ = primitive.ObjectIDFromHex(a.Ctx.Input.Param("userId"))
}
