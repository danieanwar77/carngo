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

func GetAllCasualtyType(c echo.Context) error {

	var result interface{}

	casualties, err := repository.GetAllCasualtyType(database.DbConnection)

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

func GetCasualtyType(c echo.Context) error {
	var (
		result interface{}
		casualty   structs.CasualtyType
	)
	id, _ := strconv.Atoi(c.Param("id"))

	casualty.ID = id

	casualties, err := repository.GetCasualtyType(database.DbConnection, casualty)

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

func InsertCasualtyType(c echo.Context) error {
	var (
		result interface{}
		casualty   structs.CasualtyType
	)

	err := c.Bind(&casualty)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {

		err = repository.InsertCasualtyType(database.DbConnection, casualty)
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
				Data: casualty,
			}
		}
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCasualtyType(c echo.Context) error {
	var (
		result interface{}
		casualty   structs.CasualtyType
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.Bind(&casualty)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		casualty.ID = id

		err = repository.UpdateCasualtyType(database.DbConnection, casualty)
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
				Data: casualty,
			}
		}
	}

	

	return c.JSON(http.StatusOK, result)
}

func DeleteCasualtyType(c echo.Context) error {
	var (
		result interface{}
		casualty   structs.CasualtyType
	)
	id, _ := strconv.Atoi(c.Param("id"))

	casualty.ID = id
	err := repository.DeleteCasualtyType(database.DbConnection, casualty)
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
			Data: casualty,
		}
	}

	return c.JSON(http.StatusOK, result)
}


