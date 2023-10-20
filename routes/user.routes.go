package routes

import (
	"github.com/gin-gonic/gin"
	"pop-social-assignment/controllers"
	"pop-social-assignment/middleware"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
	router.POST("/register", uc.userController.SignUpUser)
	router.POST("/login", uc.userController.SignInUser)
	router.GET("/refresh", uc.userController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(), uc.userController.LogoutUser)
}