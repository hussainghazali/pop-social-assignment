package controllers

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
	"pop-social-assignment/models"
	"pop-social-assignment/utils"
	"pop-social-assignment/initializers"
)

func TestCreatePostSuccess(t *testing.T) {
    r := gin.Default()
    postController := NewPostController(testDB)
    r.POST("/api/posts", postController.CreatePost)

    cleanDatabase(testDB)

    // Create a test user
    testUser := models.User{
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "password",
    }
    testDB.Create(&testUser)

    // Generate an access token for the user
    config, _ := initializers.LoadConfig(".")
    access_token, _ := utils.CreateToken(config.AccessTokenExpiresIn, testUser.ID, config.AccessTokenPrivateKey)

    // Create a request with a valid post payload
    payload := `{"title": "Test Post", "content": "This is a test post", "imagePath": "test.jpg", "videoPath": "test.mp4"}`
    req, _ := http.NewRequest("POST", "/api/posts", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+access_token)
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdatePostSuccess(t *testing.T) {
    r := gin.Default()
    postController := NewPostController(testDB)
    r.PUT("/api/posts/:postId", postController.UpdatePost)

    cleanDatabase(testDB)

    // Create a test user
    testUser := models.User{
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "password",
    }
    testDB.Create(&testUser)

    // Generate an access token for the user
    config, _ := initializers.LoadConfig(".")
    access_token, _ := utils.CreateToken(config.AccessTokenExpiresIn, testUser.ID, config.AccessTokenPrivateKey)

    // Create a test post
    testPost := models.Post{
        Title:     "Test Post",
        Content:   "This is a test post",
        ImagePath: "test.jpg",
        VideoPath: "test.mp4",
        UserID:    testUser.ID,
    }
    testDB.Create(&testPost)

    // Create a request with a valid post payload for updating
    payload := `{"title": "Updated Post", "content": "This is an updated post", "imagePath": "updated.jpg", "videoPath": "updated.mp4"}`
    req, _ := http.NewRequest("PUT", "/api/posts/"+testPost.ID.String(), strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+access_token)
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestFindPostByIdSuccess(t *testing.T) {
    r := gin.Default()
    postController := NewPostController(testDB)
    r.GET("/api/posts/:postId", postController.FindPostById)

    cleanDatabase(testDB)

    // Create a test user
    testUser := models.User{
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "password",
    }
    testDB.Create(&testUser)

    // Create a test post
    testPost := models.Post{
        Title:     "Test Post",
        Content:   "This is a test post",
        ImagePath: "test.jpg",
        VideoPath: "test.mp4",
        UserID:    testUser.ID,
    }
    testDB.Create(&testPost)

    req, _ := http.NewRequest("GET", "/api/posts/"+testPost.ID.String(), nil)
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestFindPostsSuccess(t *testing.T) {
    r := gin.Default()
    postController := NewPostController(testDB)
    r.GET("/api/posts", postController.FindPosts)

    cleanDatabase(testDB)

    // Create a test user
    testUser := models.User{
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "password",
    }
    testDB.Create(&testUser)

    // Generate an access token for the user
    config, _ := initializers.LoadConfig(".")
    access_token, _ := utils.CreateToken(config.AccessTokenExpiresIn, testUser.ID, config.AccessTokenPrivateKey)

    // Create test posts
    for i := 0; i < 10; i++ {
        testPost := models.Post{
            Title:     "Test Post",
            Content:   "This is test post number",
            ImagePath: "test.jpg",
            VideoPath: "test.mp4",
            UserID:    testUser.ID,
        }
        testDB.Create(&testPost)
    }

    req, _ := http.NewRequest("GET", "/api/posts", nil)
    req.Header.Set("Authorization", "Bearer "+access_token)
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeletePostSuccess(t *testing.T) {
    r := gin.Default()
    postController := NewPostController(testDB)
    r.DELETE("/api/posts/:postId", postController.DeletePost)

    cleanDatabase(testDB)

    // Create a test user
    testUser := models.User{
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "password",
    }
    testDB.Create(&testUser)

    // Generate an access token for the user
    config, _ := initializers.LoadConfig(".")
    access_token, _ := utils.CreateToken(config.AccessTokenExpiresIn, testUser.ID, config.AccessTokenPrivateKey)

    // Create a test post
    testPost := models.Post{
        Title:     "Test Post",
        Content:   "This is a test post",
        ImagePath: "test.jpg",
        VideoPath: "test.mp4",
        UserID:    testUser.ID,
    }
    testDB.Create(&testPost)

    req, _ := http.NewRequest("DELETE", "/api/posts/"+testPost.ID.String(), nil)
    req.Header.Set("Authorization", "Bearer "+access_token)
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestFindPostByUserIdSuccess(t *testing.T) {
    r := gin.Default()
    postController := NewPostController(testDB)
    r.GET("/api/posts/user", postController.FindPostByUserId)

    cleanDatabase(testDB)

    // Create a test user
    testUser := models.User{
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "password",
    }
    testDB.Create(&testUser)

    // Generate an access token for the user
    config, _ := initializers.LoadConfig(".")
    access_token, _ := utils.CreateToken(config.AccessTokenExpiresIn, testUser.ID, config.AccessTokenPrivateKey)

    // Create a test post
    testPost := models.Post{
        Title:     "Test Post",
        Content:   "This is a test post",
        ImagePath: "test.jpg",
        VideoPath: "test.mp4",
        UserID:    testUser.ID,
    }
    testDB.Create(&testPost)

    req, _ := http.NewRequest("GET", "/api/posts/user", nil)
    req.Header.Set("Authorization", "Bearer "+access_token)
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}
