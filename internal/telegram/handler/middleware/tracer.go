package middleware

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

func GetTraceId(c tele.Context) string {
	traceId, _ := c.Get(TraceIdKey).(string)
	return traceId
}

func LoggerWithTrace(logger *zap.Logger, c tele.Context) *zap.Logger {
	return logger.WithOptions(zap.Fields(zap.String("traceId", GetTraceId(c))))
}

func Tracer() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			traceId, err := uuid.NewUUID()
			if err != nil {
				return err
			}
			c.Set(TraceIdKey, traceId.String())
			return next(c)
		}
	}

}
