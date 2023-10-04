package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/gin-gonic/gin"
	"github.com/jasonalansmith/maelstrom-platform-api/database"
)

type Issue struct {
	SysId        uint   `json:"sysid,omitempty"`
	Identifier   string `json:"identifier,omitempty"`
	SummaryBrief string `json:"summary_brief,omitempty"`
	SummaryLong  string `json:"summary_long,omitempty"`
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

func getIssueById(ctx *gin.Context) {
	id := ctx.Param("sysid")

	sqls := "SELECT * FROM issue WHERE sysid = $1"

	var sysid uint
	var identifier, summary_brief, summary_long string

	res := database.Db.QueryRow(sqls, id)
	err := res.Scan(&sysid, &identifier, &summary_brief, &summary_long)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Issue not found.")
		return
	}

	issue := Issue{}
	issue.SysId = sysid
	issue.Identifier = identifier
	issue.SummaryBrief = summary_brief
	issue.SummaryLong = summary_long

	ctx.JSON(http.StatusOK, issue)
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

func patchIssue(ctx *gin.Context) {
	id := ctx.Param("sysid")

	iss := &Issue{}
	sqls := "SELECT * FROM issue WHERE sysid = $1"
	res := database.Db.QueryRow(sqls, id)
	err := res.Scan(&iss.SysId, &iss.Identifier, &iss.SummaryBrief,
		&iss.SummaryLong)
	if err == sql.ErrNoRows {
		ctx.AbortWithStatusJSON(400, "Issue does not exist.")
		return
	}

	issueBytes, err := json.Marshal(iss)
	if err != nil {
		fmt.Println("Error creating patch json ", err.Error())
		return
	}

	PatchJSON, err := io.ReadAll(ctx.Request.Body)
	fmt.Println(string(PatchJSON))
	if err != nil {
		fmt.Println(err)
	}

	patch, err := jsonpatch.DecodePatch(PatchJSON)
	if err != nil {
		fmt.Println("Error decoding patch json ", err.Error())
		return
	}

	patchedIssue, err := patch.Apply(issueBytes)
	if err != nil {
		fmt.Println("Error applying patch json ", err.Error())
		return
	}

	fmt.Println(string(patchedIssue))

	si := Issue{}
	err = json.Unmarshal(patchedIssue, &si)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, "Cannot unmarshal patchedIssue.")
		return
	}

	sqlu := "UPDATE issue SET sysid = $1, identifier = $2, "
	sqlu += "summary_brief = $3, summary_long = $4 "
	sqlu += "WHERE sysid = $5"

	fmt.Println(sqlu)

	_, err = database.Db.Exec(sqlu, si.SysId, si.Identifier,
		si.SummaryBrief, si.SummaryLong, id)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, "Cannot patch issue.")
		return
	} else {
		ctx.JSON(http.StatusOK, "Successfully patched issue.")
		return
	}
}

func mergeIssue(ctx *gin.Context) {
	id := ctx.Param("sysid")

	iss := Issue{}
	sqls := "SELECT * FROM issue WHERE sysid = $1"
	res := database.Db.QueryRow(sqls, id)
	err := res.Scan(&iss.SysId, &iss.Identifier, &iss.SummaryBrief,
		&iss.SummaryLong)
	if err == sql.ErrNoRows {
		ctx.AbortWithStatusJSON(400, "Issue does not exist.")
		return
	}

	issueBytes, err := json.Marshal(iss)
	if err != nil {
		fmt.Println("Error creating patch json ", err.Error())
		return
	}

	fmt.Println("issueBytes is: " + string(issueBytes))
	request, _ := io.ReadAll(ctx.Request.Body)
	fmt.Println("request is: " + string(request))
	patchedJSON, _ := jsonpatch.MergePatch(issueBytes, request)
	fmt.Println("patchedJSON is: " + string(patchedJSON))

	iss1 := Issue{}
	err = json.Unmarshal(patchedJSON, &iss1)

	sqlu := "UPDATE issue SET sysid = $1, identifier = $2, "
	sqlu += "summary_brief = $3, summary_long = $4 "
	sqlu += "WHERE sysid = $5"

	fmt.Println(sqlu)

	_, err = database.Db.Exec(sqlu, iss1.SysId, iss1.Identifier,
		iss1.SummaryBrief, iss1.SummaryLong, id)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, "Cannot merge issue.")
		return
	} else {
		ctx.JSON(http.StatusOK, "Successfully merged issue.")
		return
	}
}

func deleteIssue(ctx *gin.Context) {
	id := ctx.Param("sysid")

	sqle := "SELECT * FROM issue WHERE sysid = $1"
	sqld := "DELETE FROM issue WHERE sysid = $1"

	rese := database.Db.QueryRow(sqle, id)
	var sysid int
	var identifier, summary_brief, summary_long string
	err := rese.Scan(&sysid, &identifier, &summary_brief, &summary_long)
	if err == sql.ErrNoRows {
		ctx.AbortWithStatusJSON(400, "Issue does not exist.")
		return
	}

	_, err = database.Db.Exec(sqld, id)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Could not delete issue.")
		return
	} else {
		ctx.JSON(http.StatusOK, "Successfully deleted the issue.")
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
	route.GET("/issue/:sysid", getIssueById)
	route.POST("/issue/:sysid", updateIssue)
	route.PATCH("/issue/:sysid", patchIssue)
	// route.PATCH("/issue/:sysid", mergeIssue)
	route.DELETE("/issue/:sysid", deleteIssue)
	err := route.Run(":8080")
	if err != nil {
		panic(err)
	}
}
