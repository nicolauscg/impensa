package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego/context"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nicolauscg/impensa/constants"
	dt "github.com/nicolauscg/impensa/datatransfers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	BaseController
}

// @Title login user
// @Param credential  body  dt.AuthLogin true  "credential"
// @router /login [post]
func (o *AuthController) Login(credential dt.AuthLogin) {
	user, err := o.Handler.Orms.User.GetOneByEmailAndPassword(credential.Email, credential.Password)
	if err == mongo.ErrNoDocuments {
		o.ResponseBuilder.SetError(http.StatusUnauthorized, constants.ErrorIncorrectCredential).ServeJSON()

		return
	} else if err != nil {
		o.ResponseBuilder.SetError(http.StatusUnauthorized, err.Error()).ServeJSON()

		return
	}

	token, err := createJwtToken(user.Id)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusUnauthorized, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(dt.AuthPayload{user.Id, user.Email, token}).ServeJSON()
}

// @Title register user
// @Param newUser  body  dt.AuthRegister true  "newUser"
// @router /register [post]
func (o *AuthController) Register(newUser dt.AuthRegister) {
	insertResult, err := o.Handler.Orms.User.InsertOne(newUser)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData(insertResult).ServeJSON()
}

func AuthFilter(ctx *context.Context) {
	responseBuilder := dt.NewResponseBuilder(ctx.ResponseWriter)
	ctx.Output.Header("Content-Type", "application/json")
	tokenString, err := extractJwtToken(ctx)
	if err != nil {
		responseBuilder.SetError(http.StatusUnauthorized, err.Error()).ServeJSON()

		return
	}

	claims, err := validateJwtToken(tokenString)
	if err != nil {
		responseBuilder.SetError(http.StatusUnauthorized, err.Error()).ServeJSON()

		return
	}

	ctx.Input.SetParam("userId", claims["userId"].(string))
}

func createJwtToken(userId primitive.ObjectID) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 1 hour
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

func validateJwtToken(tokenString string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		err = errors.New("jwt token not valid")
	}

	return
}
