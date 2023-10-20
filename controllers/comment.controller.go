package controllers

import (
	"net/http"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"pop-social-assignment/models"
	"gorm.io/gorm"
    "github.com/google/uuid"
)

type CommentController struct {
	DB *gorm.DB
}

// ErrorResponse represents an error response.
type ErrorResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

func NewCommentController(DB *gorm.DB) CommentController {
	return CommentController{DB}
}

// CreateComment creates a new comment for a post.
// @Summary Create a comment
// @Description Create a new comment for a post.
// @ID create-comment
// @Tags comments
// @Accept json
// @Produce json
// @Param postId path string true "Post ID" Format(uuid)
// @Param request body models.CreatePostComment true "Comment request"
// @Success 201 {object} models.Comment
// @Failure 400 {object} ErrorResponse
// @Router /posts/:postId/comment [post]
func (pc *CommentController) CreateComment(ctx *gin.Context) {
    currentUser := ctx.MustGet("currentUser").(models.User)
    var payload *models.CreatePostComment

    if err := ctx.ShouldBindJSON(&payload); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }

    // Extract PostID from the query parameters
    postId := ctx.Param("postId")

    postID, err := uuid.Parse(postId)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid PostID format"})
        return
    }

    now := time.Now()
    newComment := models.Comment{
        Text:        payload.Text,
        UserID:      currentUser.ID,
        PostID:      postID,
        CreatedAt: now,
        UpdatedAt: now,
    }

    result := pc.DB.Create(&newComment)
    if result.Error != nil {
        if strings.Contains(result.Error.Error(), "duplicate key") {
            ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Comment with that text already exists"})
            return
        }
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newComment})
}

// UpdateComment updates an existing comment.
// @Summary Update a comment
// @Description Update an existing comment if the user is the owner.
// @ID update-comment
// @Tags comments
// @Accept json
// @Produce json
// @Param commentId path string true "Comment ID" Format(uuid)
// @Param request body models.UpdateComment true "Comment update request"
// @Success 200 {object} models.Comment
// @Failure 400 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/posts/:postId/comment/:commentId/update [put]
func (pc *CommentController) UpdateComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdateComment
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedComment models.Comment
	result := pc.DB.First(&updatedComment, "id = ?", commentId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Comment with that ID exists"})
		return
	}

	// Check if the user is the owner of the comment
	if updatedComment.UserID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to update this comment"})
		return
	}

	now := time.Now()

	// User is authorized to update the comment
	commentToUpdate := models.Comment{
		Text:      payload.Text,
		UserID:    currentUser.ID,
		PostID:    updatedComment.PostID,
		CreatedAt: updatedComment.CreatedAt,
		UpdatedAt: now,
	}

	pc.DB.Model(&updatedComment).Updates(commentToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedComment})
}

// DeleteComment deletes a comment.
// @Summary Delete a comment
// @Description Delete an existing comment if the user is the owner.
// @ID delete-comment
// @Tags comments
// @Accept json
// @Produce json
// @Param commentId path string true "Comment ID" Format(uuid)
// @Success 204 "No Content"
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 502 {object} ErrorResponse
// @Router /api/posts/:postId/comment/:commentId/delete [delete]
func (pc *CommentController) DeleteComment(ctx *gin.Context) {
    commentId := ctx.Param("commentId")
    currentUser := ctx.MustGet("currentUser").(models.User)

    var comment models.Comment
    result := pc.DB.First(&comment, "id = ?", commentId)
    if result.Error != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Comment with that ID exists"})
        return
    }

    // Check if the user is the owner of the comment
    if comment.UserID != currentUser.ID {
        ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to delete this comment"})
        return
    }

    result = pc.DB.Delete(&comment)

    if result.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
        return
    }

    ctx.JSON(http.StatusNoContent, nil)
}
