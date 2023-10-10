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

func PatchIssuePatchDoc(ctx *gin.Context) {
	id := ctx.Param("sysid")

	iss := &Issue{}
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

	_, err = database.MaelstromDb.Exec(sqlu, si.SysId, si.Identifier,
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
