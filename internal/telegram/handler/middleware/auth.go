package middleware

import (
	"context"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

type UserStorage interface {
	GetUserByTelegramID(ctx context.Context, id int64) (interface{}, error)
}

func Auth(storage UserStorage, key string, logger *zap.Logger) tele.MiddlewareFunc {
	return func(handler tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5)
			defer cancel()

			done := make(chan struct{})

			go func(ctx context.Context, userId int64) {
				if user, err := storage.GetUserByTelegramID(ctx, userId); err != nil {
					logger.Debug("User not found", zap.Error(err))
				} else {
					logger.Debug("User found", zap.Any("user", user))
					c.Set(key, user)
				}
				done <- struct{}{}
			}(ctxWithTimeout, c.Sender().ID)

			select {
			case <-ctxWithTimeout.Done():
				logger.Error("GetUserByTelegramID timeout exceeded")
			case <-done:
				break
			}

			return handler(c)
		}
	}
}
