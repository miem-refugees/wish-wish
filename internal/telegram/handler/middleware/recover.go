package middleware

import (
	"fmt"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

func Recover(l *zap.Logger) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			defer func() {
				if r := recover(); r != nil {
					traceIdKey := GetTraceId(c)
					logger := LoggerWithTrace(l, c)
					err := c.Send(fmt.Sprintf("Something went wrong. TraceId: %s", traceIdKey))
					if err != nil {
						logger.Error("Error while sending recover message", zap.Error(err))
					}
				}
			}()
			return next(c)
		}
	}
}
