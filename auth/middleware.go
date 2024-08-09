package auth

import (
	"carngo/structs"
	"fmt"
	"net/http"
	"time"

	"carngo/controllers"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var SecretKey = []byte("secret-key")

type User struct {
	Name  string `json:"name"`
    UUID  string `json:"uuid"`
    Admin bool   `json:"admin"`
    jwt.RegisteredClaims
}

 func LoginHandler(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		fmt.Println("user:",username, "password:", password)

		if username == "admin" {
			if password == "katasandi" {
		 
				claims:= jwt.RegisteredClaims{
					Issuer:    username,
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),

				}
				  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		 
				 //fmt.Println("token :", token)
		 
				 t, err := token.SignedString(SecretKey)
				 if err != nil {
					 var meta = structs.ResponseMetaObject{
						 Status: http.StatusBadRequest,
						 Message: err.Error(),
					 } 
					 return c.JSON(http.StatusBadRequest, meta)
				 }
		 
				 return c.JSON(http.StatusOK, map[string]string{
					 "token": t,
				 })
			 } else {
				 var meta = structs.ResponseMetaObject{
					 Status: http.StatusBadRequest,
					 Message: "Username and Password incorrect!",
				 } 
				 return c.JSON(http.StatusBadRequest, meta)
			 }
		} else {
			var getAuthStat structs.ResponseObject = controllers.GetUserAuth(c)

			if getAuthStat.Meta.Message == "Success" {
		
				claims:= jwt.RegisteredClaims{
					Issuer:    username,
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),

				}
				  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		
				t, err := token.SignedString(SecretKey)
				if err != nil {
					var meta = structs.ResponseMetaObject{
						Status: http.StatusBadRequest,
						Message: err.Error(),
					} 
					return c.JSON(http.StatusBadRequest, meta)
				} else {
					var changeUserToken structs.ResponseObject = controllers.ChangeUserToken(c, t)
					if changeUserToken.Meta.Message == "Success" {
						return c.JSON(http.StatusOK, map[string]string{
							"token": t,
						})
					} else {
						var meta = structs.ResponseMetaObject{
							Status: http.StatusBadRequest,
							Message: changeUserToken.Meta.Message,
						} 
						return c.JSON(http.StatusBadRequest, meta)
					}
				}
			} else {
				var meta = structs.ResponseMetaObject{
					Status: http.StatusBadRequest,
					Message: getAuthStat.Meta.Message,
				} 
				return c.JSON(http.StatusBadRequest, meta)
			}

		}
	
  }

  func ProtectedHandler(c echo.Context) structs.ResponseObject {
	//fmt.Println("This is the token", c.Request().Header.Get("Authorization"))
	user := c.Request().Header.Get("Authorization")

	var result structs.ResponseObject

	tokenString := user
	if tokenString == "" {
		result = structs.ResponseObject{
			Meta: structs.ResponseMetaObject{
				Status: http.StatusBadRequest,
				Message: "Missing Autohrization Header",
			},
			Data: "Empty",
		}
		return result
	}
	tokenString = tokenString[len("Bearer "):]
	
	err := VerifyToken(tokenString)
	if err != nil {
		result = structs.ResponseObject{
			Meta: structs.ResponseMetaObject{
				Status: http.StatusBadRequest,
				Message: "Invalid Token",
			},
			Data: "Empty",
		}
	} else {
		var getVerifiyToken structs.ResponseObject = controllers.GetVerifiyToken(c)

		if getVerifiyToken.Meta.Message == "Valid" {
			result = structs.ResponseObject{
				Meta: structs.ResponseMetaObject{
					Status: http.StatusOK,
					Message: "Success",
				},
				Data: getVerifiyToken.Data,
			}
		} else {

			result = structs.ResponseObject{
				Meta: structs.ResponseMetaObject{
					Status: http.StatusBadRequest,
					Message: err.Error(),
				},
				Data: "Empty",
			}
		}
	}

	return result
  }
  
  func VerifyToken(tokenString string) error {
	 token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	 })
	
	 if err != nil {
		return err
	 }
	
	 if !token.Valid {
		return fmt.Errorf("invalid token")
	 }
	
	 return nil
  }