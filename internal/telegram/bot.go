package telegram

import (
	"github.com/miem-refugees/wish-wish/internal/telegram/handler"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	bot    *tele.Bot
	logger *zap.Logger
}

func (b *Bot) Start() {
	b.bot.Start()
}

func (b *Bot) Stop() {
	b.bot.Stop()
}

func (b *Bot) Use(middlewareFunc ...tele.MiddlewareFunc) {
	b.bot.Use(middlewareFunc...)
	b.logger.Debug("Registered middlewares", zap.Int("middleware_len", len(middlewareFunc)))
}

func (b *Bot) RegisterHandler(handlers ...handler.Handler) {
	for _, h := range handlers {
		b.bot.Handle(h.Endpoint, h.HandlerFunc, h.Middleware...)
	}
	b.logger.Debug("Registered handlers", zap.Int("handlers_len", len(handlers)))
}

func NewBot(token string, logger *zap.Logger) (*Bot, error) {
	b, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10},
	})
	if err != nil {
		return nil, err
	}
	logger.Info("Authorized on account", zap.Any("username", b.Me))

	return &Bot{
		bot:    b,
		logger: logger,
	}, nil
}
