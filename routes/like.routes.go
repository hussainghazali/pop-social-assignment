package routes

import (
	"github.com/gin-gonic/gin"
	"pop-social-assignment/controllers"
	"pop-social-assignment/middleware"
)

type LikeRouteController struct {
	likeController controllers.LikeController
}

func NewLikeRouteController(likeController controllers.LikeController) LikeRouteController {
	return LikeRouteController{likeController}
}

func (rc *LikeRouteController) LikeRoute(rg *gin.RouterGroup) {
	 router := rg.Group("posts")
	 router.Use(middleware.DeserializeUser())
	 router.POST("/:postId/like", rc.likeController.CreateLike)
	 router.DELETE("/:postId/like/:likeId/delete", rc.likeController.RemoveLike)
}