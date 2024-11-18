package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"hospital-middleware/internal/handlers"
	"hospital-middleware/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStaffCreate(t *testing.T) {
	router := gin.New()
	staffService := services.NewStaffService(db)
	authService := services.NewAuthService("test-secret")
	handler := handlers.NewStaffHandler(staffService, authService)

	router.POST("/staff/create", handler.Create)

	t.Run("successful staff creation", func(t *testing.T) {
		input := map[string]interface{}{
			"username": "testuser",
			"password": "testpass",
			"hospital_id": 1,
		}
		body, _ := json.Marshal(input)
		
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/staff/create", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("missing required fields", func(t *testing.T) {
		input := map[string]interface{}{
			"username": "testuser",
		}
		body, _ := json.Marshal(input)
		
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/staff/create", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}