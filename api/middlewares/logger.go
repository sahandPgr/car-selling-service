package middlewares

import (
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
)

type ResponseLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w ResponseLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w ResponseLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func DefaultMiddlwareLogger(config *config.Config) gin.HandlerFunc {
	log := logger.NewLogger(config)
	return loggerMiddleware(log)
}

func loggerMiddleware(log logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.Contains(ctx.FullPath(), "swagger") {
			ctx.Next()
			return
		} else {
			responseWriter := &ResponseLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
			start := time.Now()
			path := ctx.FullPath()
			raw := ctx.Request.URL.RawQuery

			readBody, _ := io.ReadAll(ctx.Request.Body)
			ctx.Request.Body.Close()
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(readBody))
			ctx.Writer = responseWriter

			ctx.Next()
			param := gin.LogFormatterParams{}
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)
			param.ClientIP = ctx.ClientIP()
			param.Method = ctx.Request.Method
			param.StatusCode = ctx.Writer.Status()
			param.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
			param.BodySize = ctx.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}
			param.Path = path

			keys := map[logger.ExtraKey]interface{}{}
			keys[logger.Path] = param.Path
			keys[logger.Method] = param.Method
			keys[logger.StatusCode] = param.StatusCode
			keys[logger.Latency] = param.Latency
			keys[logger.BodySize] = param.BodySize
			keys[logger.Body] = string(readBody)
			keys[logger.ErrorMessage] = param.ErrorMessage
			keys[logger.ClientIP] = param.ClientIP
			keys[logger.ResponseBody] = responseWriter.body.String()

			log.Info(logger.RequestResponse, logger.Api, keys, "")
		}
	}
}
