package controllers

import (
	"carngo/database"
	repository "carngo/repositories"
	"carngo/structs"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUserLevel(c echo.Context) error {

	var result interface{}

	users, err := repository.GetAllUserLevel(database.DbConnection)

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

func GetUserLevel(c echo.Context) error {
	var (
		result interface{}
		level   structs.UserLevel
	)
	id, _ := strconv.Atoi(c.Param("id"))

	level.ID = id

	levels, err := repository.GetUserLevel(database.DbConnection, level)

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
			Data: &levels,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func InsertUserLevel(c echo.Context) error {
	var (
		result interface{}
		level   structs.UserLevel
	)

	err := c.Bind(&level)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {

		err = repository.InsertUserLevel(database.DbConnection, level)
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
				Data: level,
			}
		}
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUserLevel(c echo.Context) error {
	var (
		result interface{}
		level   structs.UserLevel
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.Bind(&level)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		level.ID = id

		err = repository.UpdateUserLevel(database.DbConnection, level)
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
				Data: level,
			}
		}
	}

	

	return c.JSON(http.StatusOK, result)
}

func DeleteUserLevel(c echo.Context) error {
	var (
		result interface{}
		level   structs.UserLevel
	)
	id, _ := strconv.Atoi(c.Param("id"))

	level.ID = id
	err := repository.DeleteUserLevel(database.DbConnection, level)
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
			Data: level,
		}
	}

	return c.JSON(http.StatusOK, result)
}



