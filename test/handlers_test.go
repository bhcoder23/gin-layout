package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bhcoder23/gin-layout/internal/controllers"
	"github.com/gin-gonic/gin"
)

func TestAuthRegisterRejectsInvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/auth/register", controllers.Register)

	req := httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(`{"username":"ab"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if !strings.Contains(rec.Body.String(), "Passwd") {
		t.Fatalf("response body = %q, want validation context", rec.Body.String())
	}
}

func TestGoodsPageListRejectsOversizedPage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/goods", controllers.GoodsPageList)

	req := httptest.NewRequest(http.MethodGet, "/goods?pageSize=1001", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if !strings.Contains(rec.Body.String(), "page size cannot exceed 1000") {
		t.Fatalf("response body = %q, want page size error", rec.Body.String())
	}
}
