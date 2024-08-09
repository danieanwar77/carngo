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

func GetAllBrand(c echo.Context) error {

	var result interface{}

	brands, err := repository.GetAllBrand(database.DbConnection)

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
			Data: &brands,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func GetBrand(c echo.Context) error {
	var (
		result interface{}
		brand   structs.Brand
	)
	id, _ := strconv.Atoi(c.Param("id"))

	brand.ID = id

	brands, err := repository.GetBrand(database.DbConnection, brand)

	fmt.Println("brands comeback :", brands)

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
			Data: &brands,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func InsertBrand(c echo.Context) error {
	var (
		result interface{}
		brand   structs.Brand
	)

	err := c.Bind(&brand)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {

		err = repository.InsertBrand(database.DbConnection, brand)
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
				Data: brand,
			}
		}
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateBrand(c echo.Context) error {
	var (
		result interface{}
		brand   structs.Brand
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.Bind(&brand)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		brand.ID = id

		err = repository.UpdateBrand(database.DbConnection, brand)
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
				Data: brand,
			}
		}
	}

	

	return c.JSON(http.StatusOK, result)
}

func DeleteBrand(c echo.Context) error {
	var (
		result interface{}
		brand   structs.Brand
	)
	id, _ := strconv.Atoi(c.Param("id"))

	brand.ID = id
	err := repository.DeleteBrand(database.DbConnection, brand)
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
			Data: brand,
		}
	}

	return c.JSON(http.StatusOK, result)
}


