package main

import (
	"fmt"
	"idv/chris/errs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//@see github.com/gin-contrib/zap
//Recovery 暫時不需要改寫

// MiddlewareLogger
func MiddlewareLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		fields := []zapcore.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		}

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				comps.Logger().Error(e, fields...)
			}

			return
		}

	}
}

var allow_path = map[string]bool{
	"/api/user/login":    true,
	"/api/user/register": true,
}

func MiddlewareToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		const msg = "middleware.token"
		var _path = c.Request.URL.Path
		comps.Logger().Debug(msg, zap.String("request", _path))
		if allow_path[_path] {
			c.Next()
			return
		}

		var e error
		defer func() {
			if e != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Next()
		}()

		_token, flag := c.GetQuery("token")
		if len(_token) == 0 || !flag {
			e = fmt.Errorf(string(errs.E0005))
			return
		}
		comps.Logger().Debug(msg, zap.String("token", _token))

		// 驗證 Redis (已存在則加時)
		var ukey string
		ukey, e = comps.GetToken(_token)
		if e != nil {
			comps.Logger().Debug(msg, zap.Error(e))
			e = fmt.Errorf(string(errs.E0005))
			return
		}
		// 驗證 使用者權限
		user, e := comps.IsAdmin(ukey)
		if e != nil {
			comps.Logger().Debug(msg, zap.Error(e))
			return
		}
		if user.ULv != "0" {
			e = fmt.Errorf(string(errs.E0006))
			return
		}

		comps.SetToken(_token, ukey, time.Minute*5)
		return
	}
}
