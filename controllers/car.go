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

func GetAllCar(c echo.Context) error {

	var result interface{}

	cars, err := repository.GetAllCar(database.DbConnection)

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
			Data: &cars,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func GetCar(c echo.Context) error {
	var (
		result interface{}
		car   structs.Car
	)
	id, _ := strconv.Atoi(c.Param("id"))

	car.ID = id

	cars, err := repository.GetCar(database.DbConnection, car)

	fmt.Println("cars comeback :", cars)

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
			Data: &cars,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func InsertCar(c echo.Context) error {
	var (
		result interface{}
		car   structs.Car
	)

	err := c.Bind(&car)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {

		err = repository.InsertCar(database.DbConnection, car)
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
				Data: car,
			}
		}
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCar(c echo.Context) error {
	var (
		result interface{}
		car   structs.Car
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.Bind(&car)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		car.ID = id

		err = repository.UpdateCar(database.DbConnection, car)
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
				Data: car,
			}
		}
	}

	

	return c.JSON(http.StatusOK, result)
}

func DeleteCar(c echo.Context) error {
	var (
		result interface{}
		car   structs.Car
	)
	id, _ := strconv.Atoi(c.Param("id"))

	car.ID = id
	err := repository.DeleteCar(database.DbConnection, car)
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
			Data: car,
		}
	}

	return c.JSON(http.StatusOK, result)
}


