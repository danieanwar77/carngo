package controllers

import (
	"carngo/database"
	repository "carngo/repositories"
	"carngo/structs"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserAuth(c echo.Context) structs.ResponseObject {
	var (
		result structs.ResponseObject
		user   structs.User
	)

	user.UserName = c.FormValue("username")
	user.Password = c.FormValue("password")

	_, err := repository.GetUserAuth(database.DbConnection, user)

	if err != nil {
		result = structs.ResponseObject{
			Meta: structs.ResponseMetaObject{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			},
			Data: "Empty",
		}
	} else {
		result = structs.ResponseObject{
			Meta: structs.ResponseMetaObject{
				Status:  http.StatusOK,
				Message: "Success",
			},
			Data: user,
		}
	}

	return result
}

func ChangeUserToken(c echo.Context, token string) structs.ResponseObject {
	var (
		result structs.ResponseObject
		user   structs.User
	)

	user.UserName = c.FormValue("username")
	user.Token = sql.NullString{String: token, Valid: true}

	err := repository.ChangeUserToken(database.DbConnection, user)

	if err != nil {
		result = structs.ResponseObject{
			Meta: structs.ResponseMetaObject{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			},
			Data: "Empty",
		}
	} else {
		result = structs.ResponseObject{
			Meta: structs.ResponseMetaObject{
				Status:  http.StatusOK,
				Message: "Success",
			},
			Data: "Empty",
		}
	}

	return result
}

func GetVerifiyToken(c echo.Context) structs.ResponseObject {
	var (
		result structs.ResponseObject
		user   structs.User
	)

	_, err := repository.GetUserAuth(database.DbConnection, user)

	if err != nil {
		result = structs.ResponseObject{
			Meta: structs.ResponseMetaObject{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			},
			Data: "Empty",
		}
	} else {
		result = structs.ResponseObject{
			Meta: structs.ResponseMetaObject{
				Status:  http.StatusOK,
				Message: "Valid",
			},
			Data: user,
		}
	}

	return result
}

