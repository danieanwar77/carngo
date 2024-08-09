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

func GetAllTreatment(c echo.Context) error {

	var result interface{}

	treatments, err := repository.GetAllTreatment(database.DbConnection)

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
			Data: &treatments,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func GetTreatment(c echo.Context) error {
	var (
		result interface{}
		treatment   structs.Treatment
	)
	id, _ := strconv.Atoi(c.Param("id"))

	treatment.ID = id

	treatments, err := repository.GetTreatment(database.DbConnection, treatment)

	fmt.Println("treatments comeback :", treatments)

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
			Data: &treatments,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func InsertTreatment(c echo.Context) error {
	var (
		result interface{}
		treatment   structs.Treatment
	)

	err := c.Bind(&treatment)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {

		err = repository.InsertTreatment(database.DbConnection, treatment)
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
				Data: treatment,
			}
		}
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateTreatment(c echo.Context) error {
	var (
		result interface{}
		treatment   structs.Treatment
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.Bind(&treatment)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		treatment.ID = id

		err = repository.UpdateTreatment(database.DbConnection, treatment)
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
				Data: treatment,
			}
		}
	}

	

	return c.JSON(http.StatusOK, result)
}

func DeleteTreatment(c echo.Context) error {
	var (
		result interface{}
		treatment   structs.Treatment
	)
	id, _ := strconv.Atoi(c.Param("id"))

	treatment.ID = id
	err := repository.DeleteTreatment(database.DbConnection, treatment)
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
			Data: treatment,
		}
	}

	return c.JSON(http.StatusOK, result)
}


