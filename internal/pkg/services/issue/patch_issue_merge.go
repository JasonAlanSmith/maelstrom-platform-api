package issue

import (
	"database/sql"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/gin-gonic/gin"
	"github.com/jasonalansmith/maelstrom-platform-api/internal/pkg/util/database"
)

func PatchIssueMerge(ctx *gin.Context) {
	id := ctx.Param("sysid")

	iss := Issue{}
	sqls := "SELECT * FROM select_issue_by_sysid($1);"
	res := database.MaelstromDb.QueryRow(sqls, id)
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

	_, err = database.MaelstromDb.Exec(sqlu, iss1.SysId, iss1.Identifier,
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
