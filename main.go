package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"pop-social-assignment/controllers"
	"pop-social-assignment/initializers"
	"pop-social-assignment/routes"
	"pop-social-assignment/models"
    "pop-social-assignment/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

var (
	server              *gin.Engine

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	PostController      controllers.PostController
	PostRouteController routes.PostRouteController

	CommentController      controllers.CommentController
	CommentRouteController routes.CommentRouteController

	LikeController      controllers.LikeController
	LikeRouteController routes.LikeRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	PostController = controllers.NewPostController(initializers.DB)
	PostRouteController = routes.NewRoutePostController(PostController)

	CommentController = controllers.NewCommentController(initializers.DB)
	CommentRouteController = routes.NewCommentRouteController(CommentController)

	LikeController = controllers.NewLikeController(initializers.DB)
	LikeRouteController = routes.NewLikeRouteController(LikeController)

	server = gin.Default()

	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}
	initializers.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Like{})

	// Define SwaggerInfo
	docs.SwaggerInfo.Title = "POP SOCIAL ASSIGNMENT"
	docs.SwaggerInfo.Description = "HUSSAIN GHAZALI -  Build a Go-based REST API for an Instagram-like app"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")

	router.GET("/status", func(ctx *gin.Context) {
		message := "Build a Go-based REST API for an Instagram-like app by Hussain Ghazali"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	UserRouteController.UserRoute(router)
	PostRouteController.PostRoute(router)
	CommentRouteController.CommentRoute(router)
	LikeRouteController.LikeRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
