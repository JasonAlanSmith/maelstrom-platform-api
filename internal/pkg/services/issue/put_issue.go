package issue

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jasonalansmith/maelstrom-platform-api/internal/pkg/util/database"
)

func PutIssue(ctx *gin.Context) {
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
	res := database.MaelstromDb.QueryRow(sqls, ctx.Param("sysid"))
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

	_, err = database.MaelstromDb.Exec(sql, body.SysId, body.Identifier,
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
