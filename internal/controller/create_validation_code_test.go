package controller

import (
	"context"
	"mongosteen/config"
	"mongosteen/internal/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestCreateValidationCode(t *testing.T) {
	r := gin.Default()
	config.LoadAppConfig()

	database.Connect()

	vcc := ValidationCodeController{}
	vcc.RegisterRoutes(r.Group("/api"))

	viper.Set("email.smtp.host", "localhost")
	viper.Set("email.smtp.port", "1025")
	email := "hy44jt@gmail.com"
	c := context.Background()
	q := database.NewQuery()
	count1, _ := q.CountValidationCodes(c, email)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/validation_codes",
		strings.NewReader(`{"email":"`+email+`"}`),
	)
	req.Header.Set(
		"Content-Type",
		"application/json",
	)
	r.ServeHTTP(w, req)
	count2, _ := q.CountValidationCodes(c, email)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, count2-1, count1)
	// assert.Equal(t, "pong", w.Body.String())
}
