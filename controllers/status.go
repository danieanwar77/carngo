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

func GetAllStatus(c echo.Context) error {

	var result interface{}

	statuses, err := repository.GetAllStatus(database.DbConnection)

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
			Data: &statuses,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func GetStatus(c echo.Context) error {
	var (
		result interface{}
		status   structs.Status
	)
	id, _ := strconv.Atoi(c.Param("id"))

	status.ID = id

	statuses, err := repository.GetStatus(database.DbConnection, status)

	fmt.Println("statuses comeback :", statuses)

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
			Data: &statuses,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func InsertStatus(c echo.Context) error {
	var (
		result interface{}
		status   structs.Status
	)

	err := c.Bind(&status)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {

		err = repository.InsertStatus(database.DbConnection, status)
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
				Data: status,
			}
		}
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateStatus(c echo.Context) error {
	var (
		result interface{}
		status   structs.Status
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.Bind(&status)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		status.ID = id

		err = repository.UpdateStatus(database.DbConnection, status)
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
				Data: status,
			}
		}
	}

	

	return c.JSON(http.StatusOK, result)
}

func DeleteStatus(c echo.Context) error {
	var (
		result interface{}
		status   structs.Status
	)
	id, _ := strconv.Atoi(c.Param("id"))

	status.ID = id
	err := repository.DeleteStatus(database.DbConnection, status)
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
			Data: status,
		}
	}

	return c.JSON(http.StatusOK, result)
}


