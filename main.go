package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonalansmith/maelstrom-platform-api/database"
)

type Issue struct {
	SysId        uint
	Identifier   string
	SummaryBrief string
	SummaryLong  string
}

func createIssue(ctx *gin.Context) {
	body := Issue{}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Issue is not defined.\n")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, "Bad input.\n")
		return
	}

	sql := "INSERT INTO Issue (sysid, identifier, summary_brief, "
	sql += "summary_long) VALUES ($1, $2, $3, $4)"

	_, err = database.Db.Exec(sql, body.SysId, body.Identifier,
		body.SummaryBrief, body.SummaryLong)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, "Could not create new issue.\n")
	} else {
		ctx.JSON(http.StatusOK, "Issue successfully created.\n")
	}
}

func main() {
	route := gin.Default()
	database.ConnectDatabase()
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	route.POST("/issue", createIssue)

	err := route.Run(":8080")
	if err != nil {
		panic(err)
	}
}
