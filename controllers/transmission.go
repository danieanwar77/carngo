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

func GetAllTransmission(c echo.Context) error {

	var result interface{}

	transmissions, err := repository.GetAllTransmission(database.DbConnection)

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
			Data: &transmissions,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func GetTransmission(c echo.Context) error {
	var (
		result interface{}
		transmission   structs.Transmission
	)
	id, _ := strconv.Atoi(c.Param("id"))

	transmission.ID = id

	transmissions, err := repository.GetTransmission(database.DbConnection, transmission)

	fmt.Println("transmissions comeback :", transmissions)

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
			Data: &transmissions,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func InsertTransmission(c echo.Context) error {
	var (
		result interface{}
		transmission   structs.Transmission
	)

	err := c.Bind(&transmission)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {

		err = repository.InsertTransmission(database.DbConnection, transmission)
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
				Data: transmission,
			}
		}
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateTransmission(c echo.Context) error {
	var (
		result interface{}
		transmission   structs.Transmission
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.Bind(&transmission)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		transmission.ID = id

		err = repository.UpdateTransmission(database.DbConnection, transmission)
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
				Data: transmission,
			}
		}
	}

	

	return c.JSON(http.StatusOK, result)
}

func DeleteTransmission(c echo.Context) error {
	var (
		result interface{}
		transmission   structs.Transmission
	)
	id, _ := strconv.Atoi(c.Param("id"))

	transmission.ID = id
	err := repository.DeleteTransmission(database.DbConnection, transmission)
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
			Data: transmission,
		}
	}

	return c.JSON(http.StatusOK, result)
}


