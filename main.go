package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Issue struct {
	SysId        uuid.UUID `json:"sysid,omitempty"`
	Identifier   string    `json:"identifier,omitempty"`
	SummaryBrief string    `json:"summary_brief,omitempty"`
	SummaryLong  string    `json:"summary_long,omitempty"`
}

func postIssue(ctx *gin.Context) {
	body := Issue{}
	data, err := ctx.GetRawData()
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Issue is not defined.\n")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Bad input.\n")
		return
	}

	sql := "INSERT INTO Issue (sysid, identifier, summary_brief, "
	sql += "summary_long) VALUES ($1, $2, $3, $4)"

	_, err = MaelstromDb.Exec(sql, body.SysId, body.Identifier,
		body.SummaryBrief, body.SummaryLong)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Could not create new issue.\n")
	} else {
		slog.Info("Successfully created issue.")
		ctx.JSON(http.StatusOK, "Issue successfully created.\n")
	}
}

func getIssues(ctx *gin.Context) {
	sql := "SELECT * FROM select_issue_all();"

	results, err := MaelstromDb.Query(sql)
	if err != nil {
		slog.Error("In getIssues, first call: " + err.Error())
		return
	}

	issues := []Issue{}
	for results.Next() {
		var iss Issue
		err = results.Scan(&iss.SysId, &iss.Identifier,
			&iss.SummaryBrief, &iss.SummaryLong)
		if err != nil {
			slog.Error(err.Error())
			return
		}

		issues = append(issues, iss)
	}

	slog.Info("Successfully returned all issues.")
	ctx.JSON(http.StatusOK, issues)
}

func getIssueById(ctx *gin.Context) {
	id := ctx.Param("sysid")

	sqls := "SELECT * FROM select_issue_by_sysid($1);"

	var sysid uuid.UUID
	var identifier, summary_brief, summary_long string

	res := MaelstromDb.QueryRow(sqls, id)
	err := res.Scan(&sysid, &identifier, &summary_brief, &summary_long)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Issue not found.")
		return
	}

	issue := Issue{}
	issue.SysId = sysid
	issue.Identifier = identifier
	issue.SummaryBrief = summary_brief
	issue.SummaryLong = summary_long

	slog.Info("Successfully returned one issue.")
	ctx.JSON(http.StatusOK, issue)
}

func putIssue(ctx *gin.Context) {
	body := Issue{}
	data, err := ctx.GetRawData()
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Issue is not defined.\n")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Bad input.\n")
		return
	}

	var sysid uuid.UUID
	var identifier, summary_brief, summary_long string

	sqls := "SELECT * FROM select_issue_by_sysid($1);"
	res := MaelstromDb.QueryRow(sqls, ctx.Param("sysid"))
	err = res.Scan(&sysid, &identifier, &summary_brief, &summary_long)
	if err == sql.ErrNoRows {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Issue does not exist.")
		return
	}
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "An error occurred.")
		return
	}

	sql := "UPDATE issue SET sysid = $1, identifier = $2, "
	sql += "summary_brief = $3, summary_long = $4 WHERE sysid = $5"

	_, err = MaelstromDb.Exec(sql, body.SysId, body.Identifier,
		body.SummaryBrief, body.SummaryLong, ctx.Param("sysid"))
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Could not update issue.")
		return
	} else {
		slog.Info("Successfully updated an issue.")
		ctx.JSON(http.StatusOK, "Successfully updated issue.")
	}
}

func patchIssuePatchDoc(ctx *gin.Context) {
	id := ctx.Param("sysid")

	iss := &Issue{}
	sqls := "SELECT * FROM select_issue_by_sysid($1);"
	res := MaelstromDb.QueryRow(sqls, id)
	err := res.Scan(&iss.SysId, &iss.Identifier, &iss.SummaryBrief,
		&iss.SummaryLong)
	if err == sql.ErrNoRows {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Issue does not exist.")
		return
	}

	issueBytes, err := json.Marshal(iss)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	PatchJSON, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	patch, err := jsonpatch.DecodePatch(PatchJSON)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	patchedIssue, err := patch.Apply(issueBytes)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	si := Issue{}
	err = json.Unmarshal(patchedIssue, &si)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Cannot unmarshal patchedIssue.")
		return
	}

	sqlu := "UPDATE issue SET sysid = $1, identifier = $2, "
	sqlu += "summary_brief = $3, summary_long = $4 "
	sqlu += "WHERE sysid = $5"

	_, err = MaelstromDb.Exec(sqlu, si.SysId, si.Identifier,
		si.SummaryBrief, si.SummaryLong, id)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Cannot patch issue.")
		return
	} else {
		slog.Info("Successfully patched an issue using PatchDoc.")
		ctx.JSON(http.StatusOK, "Successfully patched issue.")
		return
	}
}

func patchIssueMerge(ctx *gin.Context) {
	id := ctx.Param("sysid")

	iss := Issue{}
	sqls := "SELECT * FROM select_issue_by_sysid($1);"
	res := MaelstromDb.QueryRow(sqls, id)
	err := res.Scan(&iss.SysId, &iss.Identifier, &iss.SummaryBrief,
		&iss.SummaryLong)
	if err == sql.ErrNoRows {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Issue does not exist.")
		return
	}

	issueBytes, err := json.Marshal(iss)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	request, _ := io.ReadAll(ctx.Request.Body)
	patchedJSON, _ := jsonpatch.MergePatch(issueBytes, request)

	iss1 := Issue{}
	err = json.Unmarshal(patchedJSON, &iss1)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	sqlu := "UPDATE issue SET sysid = $1, identifier = $2, "
	sqlu += "summary_brief = $3, summary_long = $4 "
	sqlu += "WHERE sysid = $5"

	_, err = MaelstromDb.Exec(sqlu, iss1.SysId, iss1.Identifier,
		iss1.SummaryBrief, iss1.SummaryLong, id)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Cannot merge issue.")
		return
	} else {
		slog.Info("Successfully patched an issue using Merge method.")
		ctx.JSON(http.StatusOK, "Successfully merged issue.")
		return
	}
}

func deleteIssue(ctx *gin.Context) {
	id := ctx.Param("sysid")

	sqle := "SELECT * FROM select_issue_by_sysid($1);"
	sqld := "DELETE FROM issue WHERE sysid = $1"

	rese := MaelstromDb.QueryRow(sqle, id)
	var sysid uuid.UUID
	var identifier, summary_brief, summary_long string
	err := rese.Scan(&sysid, &identifier, &summary_brief, &summary_long)
	if err == sql.ErrNoRows {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Issue does not exist.")
		return
	}

	_, err = MaelstromDb.Exec(sqld, id)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Could not delete issue.")
		return
	} else {
		slog.Info("Successfully deleted an issue.")
		ctx.JSON(http.StatusOK, "Successfully deleted the issue.")
	}
}

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Next()

		latency := time.Since(t).Milliseconds()

		m := ctx.Request.Method
		rp := ctx.Request.URL.Path
		p := ctx.Request.Proto
		l := latency
		slog.Info("Request: ", "Method", m, "Request URL Path", rp, "Protocol", p, "Latency", l)
	}
}

/*
func ResponseLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("X-Content-Type-Options", "nosniff")

		ctx.Next()

		s := ctx.Writer.Status()
		m := ctx.Request.Method
		rp := ctx.Request.URL.Path
		slog.Info("Response: ", "Status", strconv.Itoa(s), "Method", m, "Request URL Path", rp)
	}
}
*/

func main() {
	logFile, err := os.OpenFile("maelstromapi.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	logger := slog.New(slog.NewJSONHandler(logFile, nil))
	slog.SetDefault(logger)

	slog.Info("The Maelstrom Platform API is starting.")

	route := gin.Default()
	route.Use(RequestLogger())
	// route.Use(ResponseLogger())
	ConnectDatabases()
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	route.POST("/issue", postIssue)
	route.GET("/issue", getIssues)
	route.GET("/issue/:sysid", getIssueById)
	route.PUT("/issue/:sysid", putIssue)
	route.PATCH("/issue/:sysid", patchIssuePatchDoc)
	// route.PATCH("/issue/:sysid", patchIssueMergeIssue)
	route.DELETE("/issue/:sysid", deleteIssue)
	err = route.Run(":8080")
	if err != nil {
		slog.Error(err.Error())
		panic(err.Error())
	}
}
