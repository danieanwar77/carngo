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

func GetAllSeller(c echo.Context) error {

	var result interface{}

	sellers, err := repository.GetAllSeller(database.DbConnection)

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
			Data: &sellers,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func GetSeller(c echo.Context) error {
	var (
		result interface{}
		seller   structs.Seller
	)
	id, _ := strconv.Atoi(c.Param("id"))

	seller.ID = id

	sellers, err := repository.GetSeller(database.DbConnection, seller)

	fmt.Println("sellers comeback :", sellers)

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
			Data: &sellers,
		}
	}

	return c.JSON(http.StatusOK, result)
}

func InsertSeller(c echo.Context) error {
	var (
		result interface{}
		seller   structs.Seller
	)

	err := c.Bind(&seller)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {

		errUser := repository.InsertUserSeller(database.DbConnection, seller)

		errSeller := repository.InsertSeller(database.DbConnection, seller)

		if errUser != nil && errSeller == nil {
			result = structs.ResponseMetaObject{
				Status:  http.StatusBadRequest,
				Message: errUser.Error(),
			}
		} else if errUser == nil && errSeller != nil {
			result = structs.ResponseMetaObject{
				Status:  http.StatusBadRequest,
				Message: errSeller.Error(),
			}
		} else if errUser != nil && errSeller != nil {
			result = structs.ResponseMetaObject{
				Status:  http.StatusBadRequest,
				Message: errUser.Error() + "|" + errSeller.Error(),
			}
		} else {
			result = structs.ResponseObject{
				Meta: structs.ResponseMetaObject{
					Status:  http.StatusOK,
					Message: "Success",
				},
				Data: seller,
			}
		}
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateSeller(c echo.Context) error {
	var (
		result interface{}
		seller   structs.Seller
	)
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.Bind(&seller)
	if err != nil {
		result = structs.ResponseMetaObject{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		seller.ID = id

		err = repository.UpdateSeller(database.DbConnection, seller)
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
				Data: seller,
			}
		}
	}

	

	return c.JSON(http.StatusOK, result)
}

func DeleteSeller(c echo.Context) error {
	var (
		result interface{}
		seller   structs.Seller
	)
	id, _ := strconv.Atoi(c.Param("id"))

	seller.ID = id
	err := repository.DeleteSeller(database.DbConnection, seller)
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
			Data: seller,
		}
	}

	return c.JSON(http.StatusOK, result)
}


