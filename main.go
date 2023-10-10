package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jasonalansmith/maelstrom-platform-api/internal/pkg/services/issue"
	"github.com/jasonalansmith/maelstrom-platform-api/internal/pkg/util/database"
	"github.com/jasonalansmith/maelstrom-platform-api/internal/pkg/util/middleware"
)

func main() {
	logFile, err := os.OpenFile("maelstromapi.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	logger := slog.New(slog.NewJSONHandler(logFile, nil))
	slog.SetDefault(logger)

	logger.Info("The Maelstrom Platform API is starting.")

	route := gin.Default()
	route.Use(middleware.RequestLogger())
	// route.Use(ResponseLogger())
	database.ConnectDatabases()
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	route.POST("/issue", issue.PostIssue)
	route.GET("/issue", issue.GetIssues)
	route.GET("/issue/:sysid", issue.GetIssueBySysId)
	route.PUT("/issue/:sysid", issue.PutIssue)
	route.PATCH("/issue/:sysid", issue.PatchIssuePatchDoc)
	// route.PATCH("/issue/:sysid", issue.PatchIssueMergeIssue)
	route.DELETE("/issue/:sysid", issue.DeleteIssue)
	err = route.Run(":8080")
	if err != nil {
		slog.Error(err.Error())
		panic(err.Error())
	}
}
