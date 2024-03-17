package rest

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"server/internal/database"
	"server/internal/services"
	"server/internal/transport/rest/private"
	"server/internal/transport/rest/public"
)

func CreateRouter(Db *sql.DB) {
	usersTable := database.UsersTable{Db: Db}
	productsTable := database.ProductsTable{Db: Db}
	publicRoutes := public.UserHandlers{UsersTable: usersTable}
	privateRoutes := private.ProductHandlers{ProductsTable: productsTable}

	router := gin.New()
	router.GET("/", publicRoutes.Index)
	router.POST("/login", publicRoutes.Login)

	privateRouter := router.Group("/products")
	privateRouter.Use(services.JwtTokenCheck)
	privateRouter.POST("/", privateRoutes.CreateProduct)
	//router.Use(AuthMiddleware())
	/*	router.POST("/products", CreateProduct)
		router.GET("/products", GetProducts)
		router.GET("/products/:id", GetProduct)
		router.PUT("/products/:id", UpdateProduct)
		router.DELETE("/products/:id", DeleteProduct)
		router.GET("/products/search", SearchProducts)*/
	router.Run()
}
