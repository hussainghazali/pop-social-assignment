package controllers

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestCreateLikeSuccess(t *testing.T) {
    r := gin.Default()
    likeController := NewLikeController(testDB)
    r.POST("/likes/:postId", likeController.CreateLike)

    cleanDatabase(testDB)

    // Create a request with a valid payload
    payload := `{"user_id": "valid-user-id", "post_id": "valid-post-id"}`
    req, _ := http.NewRequest("POST", "/likes/valid-post-id", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateLikeError(t *testing.T) {
    r := gin.Default()
    likeController := NewLikeController(testDB)
    r.POST("/likes/:postId", likeController.CreateLike)

    cleanDatabase(testDB)

    // Create a request with an invalid payload
    payload := `{"incorrect_field": "Invalid Like Data"}`
    req, _ := http.NewRequest("POST", "/likes/invalid-post-id", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
    assert.Contains(t, w.Body.String(), "Invalid PostID format")
}

func TestRemoveLikeSuccess(t *testing.T) {
    r := gin.Default()
    likeController := NewLikeController(testDB)
    r.DELETE("/likes/:likeId", likeController.RemoveLike)

    cleanDatabase(testDB)

    // Create a request with a valid like ID
    req, _ := http.NewRequest("DELETE", "/likes/valid-like-id", nil)
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestRemoveLikeError(t *testing.T) {
    r := gin.Default()
    likeController := NewLikeController(testDB)
    r.DELETE("/likes/:likeId", likeController.RemoveLike)

    cleanDatabase(testDB)

    // Create a request with an invalid like ID
    req, _ := http.NewRequest("DELETE", "/likes/invalid-like-id", nil)
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusForbidden, w.Code)
    assert.Contains(t, w.Body.String(), "You are not authorized to remove this like")
}
