package controllers

import (
	"carngo/database"
	repository "carngo/repositories"
	"carngo/structs"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {

	var result interface{}

	users, err := repository.GetAllUser(database.DbConnection)

	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		result = structs.ResponseObject{
			Meta: structs.ResponseMetaObject{
				Status:  http.StatusOK,
				Message: "Success",
			},
			Data: &users,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func GetUser(c echo.Context) error {
	var (
		result interface{}
		user   structs.User
	)
	id, _ := strconv.Atoi(c.Param("id"))

	user.ID = id

	users, err := repository.GetUser(database.DbConnection, user)

	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		result = structs.ResponseObject{
			Meta: structs.ResponseMetaObject{
				Status:  http.StatusOK,
				Message: "Success",
			},
			Data: &users,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func InsertUser(c echo.Context) error {
	var (
		result interface{}
		user   structs.User
	)

	err := c.Bind(&user)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {

		err = repository.InsertUser(database.DbConnection, user)
		if err != nil {
			result = structs.ResponseMetaObject{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
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
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUser(c echo.Context) error {
	var (
		result interface{}
		user   structs.User
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.Bind(&user)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		user.ID = id

		err = repository.UpdateUser(database.DbConnection, user)
		if err != nil {
			result = structs.ResponseMetaObject{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
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
	}

	

	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) error {
	var (
		result interface{}
		user   structs.User
	)
	id, _ := strconv.Atoi(c.Param("id"))

	user.ID = id
	err := repository.DeleteUser(database.DbConnection, user)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
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

	return c.JSON(http.StatusOK, result)
}


