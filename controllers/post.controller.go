package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"pop-social-assignment/models"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

// ResponseData represents the structure of a response data.
type ResponseData struct {
    Message string `json:"message"`
}


func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}

// CreatePost creates a new post.
// @Summary Create a post
// @Description Create a new post.
// @ID create-post
// @Tags posts
// @Produce json
// @Param currentUser header string true "Current User ID"
// @Param payload body models.CreatePostRequest true "Post data"
// @Success 201 {object} models.Post
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 502 {object} ErrorResponse
// @Router /api/post [post]
func (pc *PostController) CreatePost(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreatePostRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newPost := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		ImagePath: payload.ImagePath,
        VideoPath: payload.VideoPath,
		UserID:      currentUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := pc.DB.Create(&newPost)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPost})
}

// UpdatePost updates an existing post.
// @Summary Update a post
// @Description Update an existing post.
// @ID update-post
// @Tags posts
// @Produce json
// @Param postId path string true "Post ID"
// @Param currentUser header string true "Current User ID"
// @Param payload body models.UpdatePost true "Post data"
// @Success 200 {object} models.Post
// @Failure 404 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 502 {object} ErrorResponse
// @Router /api/post/{postId} [put]
func (pc *PostController) UpdatePost(ctx *gin.Context) {
	postID := ctx.Param("postId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdatePost
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedPost models.Post
	result := pc.DB.First(&updatedPost, "id = ?", postID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Post with that ID exists"})
		return
	}

	// Check if the user is the owner of the post
	if updatedPost.UserID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to update this post"})
		return
	}

	now := time.Now()
	postToUpdate := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		ImagePath: payload.ImagePath,
		VideoPath: payload.VideoPath,
		UserID:    currentUser.ID,
		CreatedAt: updatedPost.CreatedAt,
		UpdatedAt: now,
	}

	pc.DB.Model(&updatedPost).Updates(postToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedPost})
}

// FindPostById retrieves a post by its ID.
// @Summary Get a post by ID
// @Description Retrieve a post by its unique ID.
// @ID find-post-by-id
// @Tags posts
// @Produce json
// @Param postId path string true "Post ID"
// @Success 200 {object} models.Post
// @Failure 404 {object} ErrorResponse
// @Router /api/post/{postId} [get]
func (pc *PostController) FindPostById(ctx *gin.Context) {
    postId := ctx.Param("postId")

    var post models.Post
    result := pc.DB.Preload("Likes").Preload("Comments").First(&post, "id = ?", postId)
    if result.Error != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that ID exists"})
        return
    }

    // Calculate and update the counts
    post.LikeCount = len(post.Likes)
    post.CommentCount = len(post.Comments)

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": post})
}

// FindPosts retrieves a list of posts.
// @Summary Get a list of posts
// @Description Retrieve a list of posts with optional pagination.
// @ID find-post
// @Tags posts
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {object} ResponseData
// @Failure 502 {object} ErrorResponse
// @Router /api/posts [get]
func (pc *PostController) FindPosts(ctx *gin.Context) {
    var page = ctx.DefaultQuery("page", "1")
    var limit = ctx.DefaultQuery("limit", "10")

    intPage, _ := strconv.Atoi(page)
    intLimit, _ := strconv.Atoi(limit)
    offset := (intPage - 1) * intLimit

    var posts []models.Post
    results := pc.DB.Limit(intLimit).Offset(offset).Find(&posts)
    if results.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
        return
    }

    // Preload the Likes and Comments relationships
    pc.DB.Preload("Likes").Preload("Comments").Find(&posts)

	 // Calculate and update the counts
	 for i := range posts {
        posts[i].LikeCount = len(posts[i].Likes)
        posts[i].CommentCount = len(posts[i].Comments)
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(posts), "data": posts})
}

// DeletePost deletes a post by ID.
// @Summary Delete a post
// @Description Delete a post by its ID.
// @ID delete-post
// @Tags posts
// @Produce json
// @Param postId path string true "Post ID"
// @Success 204 {object} nil
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 502 {object} ErrorResponse
// @Router /api/posts/{postId} [delete]
func (pc *PostController) DeletePost(ctx *gin.Context) {
    postID := ctx.Param("postId")
    currentUser := ctx.MustGet("currentUser").(models.User)

    var post models.Post
    result := pc.DB.First(&post, "id = ?", postID)
    if result.Error != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Post with that ID exists"})
        return
    }

    // Check if the user is the owner of the post
    if post.UserID != currentUser.ID {
        ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to delete this post"})
        return
    }

    result = pc.DB.Delete(&post)
    if result.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
        return
    }

    ctx.JSON(http.StatusNoContent, nil)
}

// FindPostByUserId finds a post by the current user's ID.
// @Summary Find a post by the current user's ID.
// @Description Find a post by the ID of the currently authenticated user.
// @ID find-post-by-user
// @Tags posts
// @Produce json
// @Success 200 {object} models.Post
// @Failure 404 {object} ErrorResponse
// @Router /api/posts/user [get]
func (pc *PostController) FindPostByUserId(ctx *gin.Context) {
    currentUser := ctx.MustGet("currentUser").(models.User)

    var post models.Post
    result := pc.DB.Preload("Likes").Preload("Comments").First(&post, "user_id = ?", currentUser.ID)
    if result.Error != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that userId exists"})
        return
    }

    // Calculate and update the counts
    post.LikeCount = len(post.Likes)
    post.CommentCount = len(post.Comments)

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": post})
}
