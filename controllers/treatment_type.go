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

func GetAllTreatmentType(c echo.Context) error {

	var result interface{}

	treatments, err := repository.GetAllTreatmentType(database.DbConnection)

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

func GetTreatmentType(c echo.Context) error {
	var (
		result interface{}
		treatment   structs.TreatmentType
	)
	id, _ := strconv.Atoi(c.Param("id"))

	treatment.ID = id

	treatments, err := repository.GetTreatmentType(database.DbConnection, treatment)

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

func InsertTreatmentType(c echo.Context) error {
	var (
		result interface{}
		treatment   structs.TreatmentType
	)

	err := c.Bind(&treatment)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {

		err = repository.InsertTreatmentType(database.DbConnection, treatment)
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

func UpdateTreatmentType(c echo.Context) error {
	var (
		result interface{}
		treatment   structs.TreatmentType
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

		err = repository.UpdateTreatmentType(database.DbConnection, treatment)
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

func DeleteTreatmentType(c echo.Context) error {
	var (
		result interface{}
		treatment   structs.TreatmentType
	)
	id, _ := strconv.Atoi(c.Param("id"))

	treatment.ID = id
	err := repository.DeleteTreatmentType(database.DbConnection, treatment)
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


