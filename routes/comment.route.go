package routes

import (
	"github.com/gin-gonic/gin"
	"pop-social-assignment/controllers"
	"pop-social-assignment/middleware"
)

type CommentRouteController struct {
	commentController controllers.CommentController
}

func NewCommentRouteController(commentController controllers.CommentController) CommentRouteController {
	return CommentRouteController{commentController}
}

func (rc *CommentRouteController) CommentRoute(rg *gin.RouterGroup) {
	 router := rg.Group("posts")
	 router.Use(middleware.DeserializeUser())
	 router.POST("/:postId/comment", rc.commentController.CreateComment)
	 router.PUT("/:postId/comment/:commentId/update", rc.commentController.UpdateComment)
	 router.DELETE("/:postId/comment/:commentId/delete", rc.commentController.DeleteComment)
}
