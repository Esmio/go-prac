package controller

import (
	"context"
	"mongosteen/config"
	"mongosteen/config/queries"
	"mongosteen/internal"
	"mongosteen/internal/database"
	"mongosteen/internal/jwt_helper"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
	q *queries.Queries
	c context.Context
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	r = gin.Default()
	internal.InitRouter(r)

	config.LoadAppConfig()
	database.Connect()

	q = database.NewQuery()
	c = context.Background()
	if err := q.DeleteAllUsers(c); err != nil {
		t.Fatal(err)
	}
	if err := q.DeleteAllItems(c); err != nil {
		t.Fatal(err)
	}
	return func(t *testing.T) {
		database.Close()
	}
}

func signIn(t *testing.T, userID int32, req *http.Request) {
	jwtString, _ := jwt_helper.GenerateJWT(int(userID))
	req.Header = http.Header{
		"Authorization": []string{"Bearer " + jwtString},
	}
}
