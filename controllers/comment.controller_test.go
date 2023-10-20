package controllers

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "pop-social-assignment/models"
    "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func initializeTestDB() *gorm.DB {
    // Use SQLite for testing (it's lightweight)
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database: " + err.Error())
    }

    // Migrate your database schema here
    db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Like{})

    return db
}

func init() {
    testDB = initializeTestDB()
}

func cleanDatabase(db *gorm.DB) {
    db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Like{})
}

func TestCreateCommentSuccess(t *testing.T) {
    r := gin.Default()
    commentController := NewCommentController(testDB)
    r.POST("/comments/:postId", commentController.CreateComment)

    cleanDatabase(testDB)

    // Create a request with a valid payload
    payload := `{"text": "Test Comment Text"}`
    req, _ := http.NewRequest("POST", "/comments/valid-post-id", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateCommentError(t *testing.T) {
    r := gin.Default()
    commentController := NewCommentController(testDB)
    r.POST("/comments/:postId", commentController.CreateComment)

    cleanDatabase(testDB)

    // Create a request with an invalid payload
    payload := `{"incorrect_field": "Invalid Comment Text"}`
    req, _ := http.NewRequest("POST", "/comments/invalid-post-id", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
    assert.Contains(t, w.Body.String(), "Invalid PostID format")
}

func TestUpdateCommentSuccess(t *testing.T) {
    r := gin.Default()
    commentController := NewCommentController(testDB)
    r.PUT("/comments/:commentId", commentController.UpdateComment)

    cleanDatabase(testDB)

    // Create a request with a valid payload
    payload := `{"text": "Updated Comment Text"}`
    req, _ := http.NewRequest("PUT", "/comments/valid-comment-id", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateCommentError(t *testing.T) {
    r := gin.Default()
    commentController := NewCommentController(testDB)
    r.PUT("/comments/:commentId", commentController.UpdateComment)

    cleanDatabase(testDB)

    // Create a request with an invalid payload
    payload := `{"incorrect_field": "Invalid Comment Text"}`
    req, _ := http.NewRequest("PUT", "/comments/invalid-comment-id", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
    assert.Contains(t, w.Body.String(), "Invalid Comment ID format")
}

func TestDeleteCommentSuccess(t *testing.T) {
    r := gin.Default()
    commentController := NewCommentController(testDB)
    r.DELETE("/comments/:commentId", commentController.DeleteComment)

    cleanDatabase(testDB)

    // Create a request with a valid comment ID
    req, _ := http.NewRequest("DELETE", "/comments/valid-comment-id", nil)
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteCommentError(t *testing.T) {
    r := gin.Default()
    commentController := NewCommentController(testDB)
    r.DELETE("/comments/:commentId", commentController.DeleteComment)

    cleanDatabase(testDB)

    // Create a request with an invalid comment ID
    req, _ := http.NewRequest("DELETE", "/comments/invalid-comment-id", nil)
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusForbidden, w.Code)
    assert.Contains(t, w.Body.String(), "not authorized to delete this comment")
}
