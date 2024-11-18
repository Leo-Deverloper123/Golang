package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"hospital-middleware/internal/handlers"
	"hospital-middleware/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPatientSearch(t *testing.T) {
	router := gin.New()
	patientService := services.NewPatientService(db)
	handler := handlers.NewPatientHandler(patientService)

	router.GET("/patient/search", func(c *gin.Context) {
		c.Set("hospital_id", uint(1))
		handler.Search(c)
	})

	t.Run("successful patient search", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/patient/search?national_id=1234567890123", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("search with invalid parameters", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/patient/search?invalid_param=value", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}