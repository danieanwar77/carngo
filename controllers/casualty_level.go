package controllers

import (
	"carngo/database"
	repository "carngo/repositories"
	"carngo/structs"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllCasualtyLevel(c echo.Context) error {

	var result interface{}

	casualties, err := repository.GetAllCasualtyLevel(database.DbConnection)

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
			Data: &casualties,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func GetCasualtyLevel(c echo.Context) error {
	var (
		result interface{}
		casualtiy   structs.CasualtyLevel
	)
	id, _ := strconv.Atoi(c.Param("id"))

	casualtiy.ID = id

	casualties, err := repository.GetCasualtyLevel(database.DbConnection, casualtiy)

	fmt.Println("casualties comeback :", casualties)

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
			Data: &casualties,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func InsertCasualtyLevel(c echo.Context) error {
	var (
		result interface{}
		casualtiy   structs.CasualtyLevel
	)

	err := c.Bind(&casualtiy)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {

		err = repository.InsertCasualtyLevel(database.DbConnection, casualtiy)
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
				Data: casualtiy,
			}
		}
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCasualtyLevel(c echo.Context) error {
	var (
		result interface{}
		casualtiy   structs.CasualtyLevel
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.Bind(&casualtiy)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		casualtiy.ID = id

		err = repository.UpdateCasualtyLevel(database.DbConnection, casualtiy)
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
				Data: casualtiy,
			}
		}
	}

	

	return c.JSON(http.StatusOK, result)
}

func DeleteCasualtyLevel(c echo.Context) error {
	var (
		result interface{}
		casualtiy   structs.CasualtyLevel
	)
	id, _ := strconv.Atoi(c.Param("id"))

	casualtiy.ID = id
	err := repository.DeleteCasualtyLevel(database.DbConnection, casualtiy)
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
			Data: casualtiy,
		}
	}

	return c.JSON(http.StatusOK, result)
}


