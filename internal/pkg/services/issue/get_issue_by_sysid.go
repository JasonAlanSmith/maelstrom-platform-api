package issue

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jasonalansmith/maelstrom-platform-api/internal/pkg/util/database"
)

func GetIssueBySysId(ctx *gin.Context) {
	id := ctx.Param("sysid")

	sqls := "SELECT * FROM select_issue_by_sysid($1);"

	var sysid uuid.UUID
	var identifier, summary_brief, summary_long string

	res := database.MaelstromDb.QueryRow(sqls, id)
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
