package issue

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jasonalansmith/maelstrom-platform-api/internal/pkg/util/database"
)

func DeleteIssue(ctx *gin.Context) {
	id := ctx.Param("sysid")

	sqle := "SELECT * FROM select_issue_by_sysid($1);"
	sqld := "DELETE FROM issue WHERE sysid = $1"

	rese := database.MaelstromDb.QueryRow(sqle, id)
	var sysid uuid.UUID
	var identifier, summary_brief, summary_long string
	err := rese.Scan(&sysid, &identifier, &summary_brief, &summary_long)
	if err == sql.ErrNoRows {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Issue does not exist.")
		return
	}

	_, err = database.MaelstromDb.Exec(sqld, id)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Could not delete issue.")
		return
	} else {
		slog.Info("Successfully deleted an issue.")
		ctx.JSON(http.StatusOK, "Successfully deleted the issue.")
	}
}
