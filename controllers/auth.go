package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	jwt "github.com/dgrijalva/jwt-go"
	dt "github.com/nicolauscg/impensa/datatransfers"
	handlerPkg "github.com/nicolauscg/impensa/handlers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	beego.Controller
	Handler *handlerPkg.Handler
}

// @Title login user
// @router /login [post]
func (o *AuthController) Login() {
	var credential dt.UserInsert
	json.Unmarshal(o.Ctx.Input.RequestBody, &credential)
	user, err := o.Handler.Orms.User.GetOneByEmailAndPassword(credential.Email, credential.Password)
	if err == mongo.ErrNoDocuments {
		o.Data["json"] = dt.NewErrorResponse(401, "incorrect username or password")
		o.ServeJSON()

		return
	} else if err != nil {
		o.Data["json"] = dt.NewErrorResponse(401, err.Error())
		o.ServeJSON()

		return
	}
	token, err := createJwtToken(user.Id)
	if err != nil {
		o.Data["json"] = dt.NewErrorResponse(401, err.Error())
	} else {
		o.Data["json"] = dt.NewSuccessResponse(map[string]string{"token": token})
	}
	o.ServeJSON()
}

// @Title register user
// @router /register [post]
func (o *AuthController) Register() {
	var user dt.UserInsert
	json.Unmarshal(o.Ctx.Input.RequestBody, &user)
	insertResult, err := o.Handler.Orms.User.InsertOne(user)
	if err != nil {
		o.Data["json"] = dt.NewErrorResponse(500, err.Error())
		o.ServeJSON()
	}
	o.Data["json"] = dt.NewSuccessResponse(insertResult)
	o.ServeJSON()
}

func AuthFilter(ctx *context.Context) {
	ctx.Output.Header("Content-Type", "application/json")
	tokenString, err := extractJwtToken(ctx)
	if err != nil {
		ctx.Output.SetStatus(403)
		errorResponseBody, _ := json.Marshal(dt.NewErrorResponse(403, err.Error()))
		ctx.Output.Body(errorResponseBody)
		return
	}

	err = validateJwtToken(tokenString)
	if err != nil {
		ctx.Output.SetStatus(403)
		errorResponseBody, _ := json.Marshal(dt.NewErrorResponse(403, err.Error()))
		ctx.Output.Body(errorResponseBody)
		return
	}
}

func createJwtToken(userId primitive.ObjectID) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func extractJwtToken(ctx *context.Context) (string, error) {
	bearerToken := ctx.Input.Header("Authorization")
	if splitRes := strings.Split(bearerToken, " "); len(splitRes) == 2 {
		return splitRes[1], nil
	}

	return "", errors.New("bearer token not found in Authorization header")
}

func validateJwtToken(tokenString string) (err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid || claims == nil {
		err = errors.New("jwt token not valid")
	}

	return
}
