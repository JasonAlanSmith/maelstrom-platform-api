package issue

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonalansmith/maelstrom-platform-api/internal/pkg/util/database"
	"github.com/jasonalansmith/maelstrom-platform-api/internal/pkg/util/unmarshal"
)

func PostIssue(ctx *gin.Context) {
	body := Issue{}
	data, err := ctx.GetRawData()
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Issue is not defined.\n")
		return
	}

	var val_err []string
	err, val_err = unmarshal.Unmarshal(data, &body)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Bad input.\n")
		return
	} else {
		if len(val_err) > 0 {
			for i := 0; i < len(val_err); i++ {
				slog.Error(val_err[i])
			}
			ctx.AbortWithStatusJSON(400, "One or more validation errors found.")
			return
		}
	}

	sql := "INSERT INTO Issue (sysid, identifier, summary_brief, "
	sql += "summary_long) VALUES ($1, $2, $3, $4)"

	_, err = database.MaelstromDb.Exec(sql, body.SysId, body.Identifier,
		body.SummaryBrief, body.SummaryLong)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Could not create new issue.\n")
	} else {
		slog.Info("Successfully created issue.")
		ctx.JSON(http.StatusOK, "Issue successfully created.\n")
	}
}
