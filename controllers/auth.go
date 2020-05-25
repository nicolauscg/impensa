package controllers

import (
	"encoding/json"
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
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	BaseController
}

// @Title login user
// @Param credential  body  dt.AuthLogin true  "credential"
// @router /login [post]
func (o *AuthController) Login(credential dt.AuthLogin) {
	user, err := o.Handler.Orms.User.GetOneByEmail(credential.Email)
	if err == mongo.ErrNoDocuments {
		o.ResponseBuilder.SetError(http.StatusUnauthorized, constants.ErrorEmailNotRegistered).ServeJSON()

		return
	} else if !comparePasswords(user.Password, credential.Password) {
		o.ResponseBuilder.SetError(http.StatusUnauthorized, constants.ErrorIncorrectPassword).ServeJSON()

		return
	} else if err != nil {
		o.ResponseBuilder.SetError(http.StatusUnauthorized, err.Error()).ServeJSON()

		return
	}
	var jwtExpiry time.Duration
	var cookieExpiry int
	if credential.RememberMe {
		jwtExpiry = time.Hour * 24 * 7 * 4
		cookieExpiry = 60 * 60 * 24 * 7 * 4
	} else {
		jwtExpiry = time.Hour * 24
	}
	token, err := createJwtToken(user.Id, jwtExpiry)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusUnauthorized, err.Error()).ServeJSON()

		return
	}
	authPayloadJson, _ := json.Marshal(dt.AuthPayload{user.Id, user.Username, token})
	if credential.RememberMe {
		o.Ctx.SetCookie("impensa", string(authPayloadJson), cookieExpiry)
	} else {
		o.Ctx.SetCookie("impensa", string(authPayloadJson))
	}
	o.ResponseBuilder.SetData(dt.AuthPayload{user.Id, user.Username, token}).ServeJSON()
}

// @Title register user
// @Param newUser  body  dt.AuthRegister true  "newUser"
// @router /register [post]
func (o *AuthController) Register(newUser dt.AuthRegister) {
	newUser.Email = strings.ToLower(newUser.Email)
	temp := false
	newUser.Verified = &temp

	user, err := o.Handler.Orms.User.GetOneByEmail(newUser.Email)
	if user != nil && err == nil {
		o.ResponseBuilder.SetError(http.StatusConflict, constants.ErrorEmailAlreadyRegistered).ServeJSON()

		return
	} else if err != nil && err != mongo.ErrNoDocuments {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	user, err = o.Handler.Orms.User.GetOneByUsername(newUser.Username)
	if user != nil && err == nil {
		o.ResponseBuilder.SetError(http.StatusConflict, constants.ErrorUsernameAlreadyTaken).ServeJSON()

		return
	} else if err != nil && err != mongo.ErrNoDocuments {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}

	newUserWithHashedPassword := newUser
	hashedPassword, err := hashAndSalt(newUser.Password)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}
	newUserWithHashedPassword.Password = hashedPassword
	insertResult, err := o.Handler.Orms.User.InsertOne(newUserWithHashedPassword)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

		return
	}

	o.Handler.Orms.VerifyAccount.InsertOne(insertResult.InsertedID.(primitive.ObjectID))
	o.Login(dt.AuthLogin{newUser.Email, newUser.Password, false})
}

func AuthFilter(ctx *context.Context) {
	responseBuilder := dt.NewResponseBuilder(ctx.ResponseWriter)
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

func createJwtToken(userId primitive.ObjectID, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(expiry).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv(constants.EnvApiSecret)))
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
		return []byte(os.Getenv(constants.EnvApiSecret)), nil
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

func hashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func comparePasswords(hashedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return false
	}

	return true
}

// @Title verify user
// @Param userId  query string true  "userId"
// @Param verifyKey  query  string true  "verifyKey"
// @router /verify [get]
func (o *AuthController) VerifyUser(userId *string, verifyKey *string) {
	userObjectId, err := primitive.ObjectIDFromHex(*userId)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusForbidden, err.Error()).ServeJSON()

		return
	}
	exist, err := o.Handler.Orms.VerifyAccount.Verify(userObjectId, *verifyKey)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusForbidden, err.Error()).ServeJSON()

		return
	}

	if exist {
		o.ResponseBuilder.SetData("account verified").ServeJSON()
	} else {
		o.ResponseBuilder.SetError(http.StatusForbidden, constants.ErrorIncorrectVerifyKey).ServeJSON()
	}
}

// @Title request reset password
// @Param requestResetUserPasswordBody  body dt.RequestResetUserPasswordBody true  "requestResetUserPasswordBody"
// @router /requestreset [post]
func (o *AuthController) RequestResetUserPassword(requestResetUserPasswordBody dt.RequestResetUserPasswordBody) {
	user, err := o.Handler.Orms.User.GetOneByEmail(*requestResetUserPasswordBody.Email)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusForbidden, err.Error()).ServeJSON()

		return
	}
	_, _, err = o.Handler.Orms.ResetUserPassword.InsertOne(user.Id)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusForbidden, err.Error()).ServeJSON()

		return
	}
	o.ResponseBuilder.SetData("password request sent").ServeJSON()
}

// @Title reset account password
// @Param resetUserPasswordBody  body dt.ResetUserPasswordBody true  "resetUserPasswordBody"
// @router /resetpassword [post]
func (o *AuthController) ResetUserPassword(resetUserPasswordBody dt.ResetUserPasswordBody) {
	user, err := o.Handler.Orms.User.GetOneByEmail(*resetUserPasswordBody.Email)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusForbidden, err.Error()).ServeJSON()

		return
	}
	userUpdateInModel := &dt.UserUpdateFieldsInModel{}
	exist, err := o.Handler.Orms.ResetUserPassword.Verify(user.Id, *resetUserPasswordBody.VerifyKey)
	if err != nil {
		o.ResponseBuilder.SetError(http.StatusForbidden, err.Error()).ServeJSON()

		return
	}

	if exist {
		if resetUserPasswordBody.OldPassword != nil && resetUserPasswordBody.NewPassword != nil {
			user, err := o.Handler.Orms.User.GetOneWithPasswordById(user.Id)
			if err != nil {
				o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

				return
			}
			if !comparePasswords(user.Password, *resetUserPasswordBody.OldPassword) {
				o.ResponseBuilder.SetError(http.StatusInternalServerError, constants.ErrorOldPasswordMismatch).ServeJSON()

				return
			} else {
				hashedPassword, err := hashAndSalt(*resetUserPasswordBody.NewPassword)
				if err != nil {
					o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

					return
				}

				userUpdateInModel.Password = &hashedPassword
			}
		}

		updateResult, err := o.Handler.Orms.User.UpdateOneById(
			user.Id,
			userUpdateInModel,
		)
		if err != nil {
			o.ResponseBuilder.SetError(http.StatusInternalServerError, err.Error()).ServeJSON()

			return
		}

		o.ResponseBuilder.SetData(updateResult).ServeJSON()
	} else {
		o.ResponseBuilder.SetError(http.StatusForbidden, constants.ErrorIncorrectVerifyKey).ServeJSON()
	}
}
