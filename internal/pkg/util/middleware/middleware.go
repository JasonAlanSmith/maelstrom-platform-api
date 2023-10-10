package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

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
