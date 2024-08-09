package main

import (

	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"

	auth "carngo/auth"
	controller "carngo/controllers"
	"carngo/database"
	"carngo/structs"
)

var (
	DB  *sql.DB
	err error
)

func main() {

	err = godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)


	router := echo.New()

    router.Use(middleware.Logger())
    router.Use(middleware.Recover())

	publicRoutes := router.Group("/v1")
	{
		publicRoutes.GET("", func(c echo.Context) error {
			return c.String(http.StatusOK, "Welcome to Car n Go API v1")
		})
		publicRoutes.POST("/login", auth.LoginHandler)
	}
    
	protectedRoutes := router.Group("/v1")
	{
		protectedRoutes.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				var middlewareAuth structs.ResponseObject = auth.ProtectedHandler(c)
				if middlewareAuth.Meta.Message != "Success" {
				return c.String(http.StatusBadRequest, middlewareAuth.Meta.Message)
				}
				return next(c)
			}
		})
		user := protectedRoutes.Group("/users")
		{
			user.GET("", controller.GetAllUser)
			user.GET("/:id", controller.GetUser)
			user.POST("", controller.InsertUser)
			user.PUT("/:id", controller.UpdateUser)
			user.DELETE("/:id", controller.DeleteUser)
		}
		userlevel := protectedRoutes.Group("/userslevel")
		{
			userlevel.GET("", controller.GetAllUserLevel)
			userlevel.GET("/:id", controller.GetUserLevel)
			userlevel.POST("", controller.InsertUserLevel)
			userlevel.PUT("/:id", controller.UpdateUserLevel)
			userlevel.DELETE("/:id", controller.DeleteUserLevel)
		}
		brand := protectedRoutes.Group("/brand")
		{
			brand.GET("", controller.GetAllBrand)
			brand.GET("/:id", controller.GetBrand)
			brand.POST("", controller.InsertBrand)
			brand.PUT("/:id", controller.UpdateBrand)
			brand.DELETE("/:id", controller.DeleteBrand)
		}
		car := protectedRoutes.Group("/car")
		{
			car.GET("", controller.GetAllCar)
			car.GET("/:id", controller.GetCar)
			car.POST("", controller.InsertCar)
			car.PUT("/:id", controller.UpdateCar)
			car.DELETE("/:id", controller.DeleteCar)
		}
		casualtyLevel := protectedRoutes.Group("/casualtylevel")
		{
			casualtyLevel.GET("", controller.GetAllCasualtyLevel)
			casualtyLevel.GET("/:id", controller.GetCasualtyLevel)
			casualtyLevel.POST("", controller.InsertCasualtyLevel)
			casualtyLevel.PUT("/:id", controller.UpdateCasualtyLevel)
			casualtyLevel.DELETE("/:id", controller.DeleteCasualtyLevel)
		}
		casualtyType := protectedRoutes.Group("/casualtytype")
		{
			casualtyType.GET("", controller.GetAllCasualtyType)
			casualtyType.GET("/:id", controller.GetCasualtyType)
			casualtyType.POST("", controller.InsertCasualtyType)
			casualtyType.PUT("/:id", controller.UpdateCasualtyType)
			casualtyType.DELETE("/:id", controller.DeleteCasualtyType)
		}
		casualty := protectedRoutes.Group("/casualty")
		{
			casualty.GET("", controller.GetAllCasualty)
			casualty.GET("/:id", controller.GetCasualty)
			casualty.POST("", controller.InsertCasualty)
			casualty.PUT("/:id", controller.UpdateCasualty)
			casualty.DELETE("/:id", controller.DeleteCasualty)
		}
		condition := protectedRoutes.Group("/condition")
		{
			condition.GET("", controller.GetAllCondition)
			condition.GET("/:id", controller.GetCondition)
			condition.POST("", controller.InsertCondition)
			condition.PUT("/:id", controller.UpdateCondition)
			condition.DELETE("/:id", controller.DeleteCondition)
		}
		registrationStatus := protectedRoutes.Group("/registrationstatus")
		{
			registrationStatus.GET("", controller.GetAllRegistrationStatus)
			registrationStatus.GET("/:id", controller.GetRegistrationStatus)
			registrationStatus.POST("", controller.InsertRegistrationStatus)
			registrationStatus.PUT("/:id", controller.UpdateRegistrationStatus)
			registrationStatus.DELETE("/:id", controller.DeleteRegistrationStatus)
		}
		seller := protectedRoutes.Group("/seller")
		{
			seller.GET("", controller.GetAllSeller)
			seller.GET("/:id", controller.GetSeller)
			seller.POST("", controller.InsertSeller)
			seller.PUT("/:id", controller.UpdateSeller)
			seller.DELETE("/:id", controller.DeleteSeller)
		}
		status := protectedRoutes.Group("/status")
		{
			status.GET("", controller.GetAllStatus)
			status.GET("/:id", controller.GetStatus)
			status.POST("", controller.InsertStatus)
			status.PUT("/:id", controller.UpdateStatus)
			status.DELETE("/:id", controller.DeleteStatus)
		}
		transmission := protectedRoutes.Group("/transmission")
		{
			transmission.GET("", controller.GetAllTransmission)
			transmission.GET("/:id", controller.GetTransmission)
			transmission.POST("", controller.InsertTransmission)
			transmission.PUT("/:id", controller.UpdateTransmission)
			transmission.DELETE("/:id", controller.DeleteTransmission)
		}
		treatmentType := protectedRoutes.Group("/treatmenttype")
		{
			treatmentType.GET("", controller.GetAllTreatmentType)
			treatmentType.GET("/:id", controller.GetTreatmentType)
			treatmentType.POST("", controller.InsertTreatmentType)
			treatmentType.PUT("/:id", controller.UpdateTreatmentType)
			treatmentType.DELETE("/:id", controller.DeleteTreatmentType)
		}
		treatment := protectedRoutes.Group("/treatment")
		{
			treatment.GET("", controller.GetAllTreatment)
			treatment.GET("/:id", controller.GetTreatment)
			treatment.POST("", controller.InsertTreatment)
			treatment.PUT("/:id", controller.UpdateTreatment)
			treatment.DELETE("/:id", controller.DeleteTreatment)
		}
	}
    
	router.Start(":" + "8080")

	fmt.Println("Successfully Connected!")
}
