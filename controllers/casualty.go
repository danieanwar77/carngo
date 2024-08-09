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

func GetAllCasualty(c echo.Context) error {

	var result interface{}

	casualties, err := repository.GetAllCasualty(database.DbConnection)

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

func GetCasualty(c echo.Context) error {
	var (
		result interface{}
		casualtiy   structs.Casualty
	)
	id, _ := strconv.Atoi(c.Param("id"))

	casualtiy.ID = id

	casualties, err := repository.GetCasualty(database.DbConnection, casualtiy)

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

func InsertCasualty(c echo.Context) error {
	var (
		result interface{}
		casualtiy   structs.Casualty
	)

	err := c.Bind(&casualtiy)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {

		err = repository.InsertCasualty(database.DbConnection, casualtiy)
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

func UpdateCasualty(c echo.Context) error {
	var (
		result interface{}
		casualtiy   structs.Casualty
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

		err = repository.UpdateCasualty(database.DbConnection, casualtiy)
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

func DeleteCasualty(c echo.Context) error {
	var (
		result interface{}
		casualtiy   structs.Casualty
	)
	id, _ := strconv.Atoi(c.Param("id"))

	casualtiy.ID = id
	err := repository.DeleteCasualty(database.DbConnection, casualtiy)
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


