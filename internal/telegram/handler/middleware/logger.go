package middleware

import (
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

func Logger(logger *zap.Logger) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			LoggerWithTrace(logger, c).Info("IncomingEvent", zap.Any("data", c.Update()))
			return next(c)
		}
	}
}
