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

func TestGetMeSuccess(t *testing.T) {
    r := gin.Default()
    userController := NewUserController(testDB)
    r.GET("/api/users/me", userController.GetMe)

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
    w := httptest.NewRecorder()

    // Set the access token in the request header
    req, _ := http.NewRequest("GET", "/api/users/me", nil)
    req.Header.Set("Authorization", "Bearer "+access_token)
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestSignUpUserSuccess(t *testing.T) {
    r := gin.Default()
    userController := NewUserController(testDB)
    r.POST("/api/users/signup", userController.SignUpUser)

    cleanDatabase(testDB)

    // Create a request with valid user data
    payload := `{"name": "Test User", "email": "test@example.com", "password": "password", "passwordConfirm": "password"}`
    req, _ := http.NewRequest("POST", "/api/users/signup", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)
}

func TestSignUpUserPasswordMismatch(t *testing.T) {
    r := gin.Default()
    userController := NewUserController(testDB)
    r.POST("/api/users/signup", userController.SignUpUser)

    cleanDatabase(testDB)

    // Create a request with password and passwordConfirm mismatch
    payload := `{"name": "Test User", "email": "test@example.com", "password": "password", "passwordConfirm": "mismatched"}`
    req, _ := http.NewRequest("POST", "/api/users/signup", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSignInUserSuccess(t *testing.T) {
    r := gin.Default()
    userController := NewUserController(testDB)
    r.POST("/api/users/signin", userController.SignInUser)

    cleanDatabase(testDB)

    // Create a test user
    testUser := models.User{
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "password",
    }
    testDB.Create(&testUser)

    // Create a request with valid user credentials
    payload := `{"email": "test@example.com", "password": "password"}`
    req, _ := http.NewRequest("POST", "/api/users/signin", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestSignInUserInvalidCredentials(t *testing.T) {
    r := gin.Default()
    userController := NewUserController(testDB)
    r.POST("/api/users/signin", userController.SignInUser)

    cleanDatabase(testDB)

    // Create a request with invalid user credentials
    payload := `{"email": "test@example.com", "password": "invalid_password"}`
    req, _ := http.NewRequest("POST", "/api/users/signin", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRefreshAccessTokenSuccess(t *testing.T) {
    r := gin.Default()
    userController := NewUserController(testDB)
    r.GET("/api/users/refresh", userController.RefreshAccessToken)

    cleanDatabase(testDB)

    // Create a test user
    testUser := models.User{
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "password",
    }
    testDB.Create(&testUser)

    // Generate a refresh token for the user
    config, _ := initializers.LoadConfig(".")
    refresh_token, _ := utils.CreateToken(config.RefreshTokenExpiresIn, testUser.ID, config.RefreshTokenPrivateKey)

    // Set the refresh token in the request cookie
    req, _ := http.NewRequest("GET", "/api/users/refresh", nil)
    req.AddCookie(&http.Cookie{Name: "refresh_token", Value: refresh_token})
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestLogoutUser(t *testing.T) {
    r := gin.Default()
    userController := NewUserController(testDB)
    r.GET("/api/users/logout", userController.LogoutUser)

    // Create a request to log the user out
    req, _ := http.NewRequest("GET", "/api/users/logout", nil)
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}
