package issue

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonalansmith/maelstrom-platform-api/internal/pkg/util/database"
)

func GetIssues(ctx *gin.Context) {
	sql := "SELECT * FROM select_issue_all();"

	results, err := database.MaelstromDb.Query(sql)
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
