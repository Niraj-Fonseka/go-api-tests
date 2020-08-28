package controllers

import (
	"go-api-tests/models"
	"go-api-tests/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAllClassesFail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockClassService := new(services.MockClassService)

	mockClassService.On("GetAllClasses").Return([]models.Class{}, nil)

	controller := NewClassesController()
	controller.ClassServiceInterface = mockClassService

	r := gin.Default()

	r.GET("/v1/classes", controller.GetClasses)

	req, err := http.NewRequest(http.MethodGet, "/v1/classes", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
