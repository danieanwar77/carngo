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

func GetAllRegistrationStatus(c echo.Context) error {

	var result interface{}

	regs, err := repository.GetAllRegistrationStatus(database.DbConnection)

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
			Data: &regs,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func GetRegistrationStatus(c echo.Context) error {
	var (
		result interface{}
		reg   structs.RegistrationStatus
	)
	id, _ := strconv.Atoi(c.Param("id"))

	reg.ID = id

	regs, err := repository.GetRegistrationStatus(database.DbConnection, reg)

	fmt.Println("regs comeback :", regs)

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
			Data: &regs,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func InsertRegistrationStatus(c echo.Context) error {
	var (
		result interface{}
		reg   structs.RegistrationStatus
	)

	err := c.Bind(&reg)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {

		err = repository.InsertRegistrationStatus(database.DbConnection, reg)
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
				Data: reg,
			}
		}
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateRegistrationStatus(c echo.Context) error {
	var (
		result interface{}
		reg   structs.RegistrationStatus
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.Bind(&reg)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		reg.ID = id

		err = repository.UpdateRegistrationStatus(database.DbConnection, reg)
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
				Data: reg,
			}
		}
	}

	

	return c.JSON(http.StatusOK, result)
}

func DeleteRegistrationStatus(c echo.Context) error {
	var (
		result interface{}
		reg   structs.RegistrationStatus
	)
	id, _ := strconv.Atoi(c.Param("id"))

	reg.ID = id
	err := repository.DeleteRegistrationStatus(database.DbConnection, reg)
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
			Data: reg,
		}
	}

	return c.JSON(http.StatusOK, result)
}


