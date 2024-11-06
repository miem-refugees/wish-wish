package handler

import (
	"go.uber.org/zap"

	tele "gopkg.in/telebot.v3"
)

type CommonHandler struct {
	logger *zap.Logger
}

func CommonHandlers(logger *zap.Logger) []Handler {
	h := &CommonHandler{
		logger: logger.Named("CommonHandler"),
	}
	return []Handler{
		{
			Endpoint:    "/start",
			HandlerFunc: h.start,
		},
		{
			Endpoint:    "/help",
			HandlerFunc: h.help,
		},
		{
			Endpoint:    "/panic",
			HandlerFunc: h.panic,
		},
	}
}

func (h *CommonHandler) start(c tele.Context) error {
	return c.Send("Hello, " + c.Sender().FirstName)
}

func (h *CommonHandler) help(c tele.Context) error {
	return c.Send(c.Sender().FirstName, " this is help message")
}

func (h *CommonHandler) panic(c tele.Context) error {
	panic("Panic handler! Woo-hoo!")
}
