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

func GetAllCondition(c echo.Context) error {

	var result interface{}

	conditions, err := repository.GetAllCondition(database.DbConnection)

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
			Data: &conditions,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func GetCondition(c echo.Context) error {
	var (
		result interface{}
		condition   structs.Condition
	)
	id, _ := strconv.Atoi(c.Param("id"))

	condition.ID = id

	conditions, err := repository.GetCondition(database.DbConnection, condition)

	fmt.Println("conditions comeback :", conditions)

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
			Data: &conditions,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func InsertCondition(c echo.Context) error {
	var (
		result interface{}
		condition   structs.Condition
	)

	err := c.Bind(&condition)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {

		err = repository.InsertCondition(database.DbConnection, condition)
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
				Data: condition,
			}
		}
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCondition(c echo.Context) error {
	var (
		result interface{}
		condition   structs.Condition
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.Bind(&condition)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		condition.ID = id

		err = repository.UpdateCondition(database.DbConnection, condition)
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
				Data: condition,
			}
		}
	}

	

	return c.JSON(http.StatusOK, result)
}

func DeleteCondition(c echo.Context) error {
	var (
		result interface{}
		condition   structs.Condition
	)
	id, _ := strconv.Atoi(c.Param("id"))

	condition.ID = id
	err := repository.DeleteCondition(database.DbConnection, condition)
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
			Data: condition,
		}
	}

	return c.JSON(http.StatusOK, result)
}


