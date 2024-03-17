package rest

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"server/internal/config"
	"server/internal/database"
	"server/internal/services"
	"server/internal/transport/rest/private"
	"server/internal/transport/rest/public"
)

func CreateRouter(Db *sql.DB) {
	usersTable := database.UsersTable{Db: Db}
	productsTable := database.ProductsTable{Db: Db}

	jwtDbConnector := services.JwtDbConnector{UsersTable: usersTable}
	publicRoutes := public.UserHandlers{UsersTable: usersTable}
	privateRoutes := private.ProductHandlers{ProductsTable: productsTable}

	ginMode := config.GetEnv(config.GinMode)
	if ginMode == config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.GET("/", publicRoutes.Index)
	router.POST("/login", publicRoutes.Login)

	privateRouter := router.Group("/products")
	privateRouter.Use(jwtDbConnector.JwtTokenCheck)
	privateRouter.POST("/", privateRoutes.CreateProduct)
	privateRouter.GET("/", privateRoutes.GetProducts)
	privateRouter.PUT("/:id", privateRoutes.UpdateProduct)
	privateRouter.GET("/:id", privateRoutes.GetProduct)
	privateRouter.DELETE("/:id", privateRoutes.DeleteProduct)

	err := router.Run(":" + config.GetEnv(config.GinPort))
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
