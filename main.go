package main

import (
	"database/sql"
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

func getIssues(ctx *gin.Context) {
	sql := "SELECT * FROM issue"

	results, err := database.Db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	issues := []Issue{}
	for results.Next() {
		var iss Issue
		err = results.Scan(&iss.SysId, &iss.Identifier,
			&iss.SummaryBrief, &iss.SummaryLong)
		if err != nil {
			panic(err.Error())
		}

		issues = append(issues, iss)
	}

	ctx.JSON(http.StatusOK, issues)
}

func updateIssue(ctx *gin.Context) {
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

	var sysid int
	var identifier, summary_brief, summary_long string

	sqls := "SELECT * FROM issue WHERE sysid = $1"
	res := database.Db.QueryRow(sqls, ctx.Param("sysid"))
	err = res.Scan(&sysid, &identifier, &summary_brief, &summary_long)
	if err == sql.ErrNoRows {
		ctx.AbortWithStatusJSON(400, "Issue does not exist.")
		return
	}
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, "An error occurred.")
		return
	}

	sql := "UPDATE issue SET sysid = $1, identifier = $2, "
	sql += "summary_brief = $3, summary_long = $4 WHERE sysid = $5"

	_, err = database.Db.Exec(sql, body.SysId, body.Identifier,
		body.SummaryBrief, body.SummaryLong, ctx.Param("sysid"))
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, "Could not update issue.")
	} else {
		ctx.JSON(http.StatusOK, "Successfully updated issue.")
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
	route.GET("/issue", getIssues)
	route.POST("/issue/:sysid", updateIssue)
	err := route.Run(":8080")
	if err != nil {
		panic(err)
	}
}
