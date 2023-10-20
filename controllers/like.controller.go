package controllers

import (
	"net/http"
	"errors"
	"time"
	"github.com/gin-gonic/gin"
	"pop-social-assignment/models"
	"gorm.io/gorm"
    "github.com/google/uuid"
)

type LikeController struct {
	DB *gorm.DB
}

func NewLikeController(DB *gorm.DB) LikeController {
	return LikeController{DB}
}

// CreateLike creates a new like for a post.
// @Summary Create a new like
// @Description Create a new like for a post.
// @ID create-like
// @Tags likes
// @Produce json
// @Param postId path string true "Post ID"
// @Success 201 {object} models.Like
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 502 {object} ErrorResponse
// @Router /api/posts/:postId/like [post]
func (pc *LikeController) CreateLike(ctx *gin.Context) {
    currentUser := ctx.MustGet("currentUser").(models.User)
    var payload *models.CreatePostLike

    if err := ctx.ShouldBindJSON(&payload); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }

    // Extract PostID from the route parameter
    postId := ctx.Param("postId")

    postID, err := uuid.Parse(postId)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid PostID format"})
        return
    }

    // Check if a like with the same UserID and PostID already exists
    var existingLike models.Like
    result := pc.DB.Where("user_id = ? AND post_id = ?", currentUser.ID, postID).First(&existingLike)
    if result.Error == nil {
        // A like with the same UserID and PostID already exists
        ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Like already exists for this post against current user"})
        return
    } else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
        // An error occurred while checking for the existing like
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
        return
    }

    // Create a new like
    now := time.Now()
    newLike := models.Like{
        UserID:    currentUser.ID,
        PostID:    postID,
        CreatedAt: now,
        UpdatedAt: now,
    }

    result = pc.DB.Create(&newLike)
    if result.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newLike})
}

// RemoveLike removes a like for a post.
// @Summary Remove a like
// @Description Remove a like for a post.
// @ID remove-like
// @Tags likes
// @Produce json
// @Param likeId path string true "Like ID"
// @Success 204 "No Content"
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 502 {object} ErrorResponse
// @Router /api/posts/:postId/like/:likeId/delete} [delete]
func (pc *LikeController) RemoveLike(ctx *gin.Context) {
	likeID := ctx.Param("likeId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var like models.Like
	result := pc.DB.First(&like, "id = ?", likeID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Like with that ID exists"})
		return
	}

	// Check if the user is the owner of the like
	if like.UserID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to remove this like"})
		return
	}

	result = pc.DB.Delete(&like)

	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
