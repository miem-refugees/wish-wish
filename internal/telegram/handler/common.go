package handler

import (
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

type CommonHandler struct {
	logger *zap.Logger
}

func CommonHandlers(logger *zap.Logger) []*Handler {
	h := &CommonHandler{
		logger: logger.Named("CommonHandler"),
	}

	return []*Handler{
		h.handleStart(),
		h.handleHelp(),
	}
}

func (h *CommonHandler) handleStart() *Handler {
	return &Handler{
		Endpoint: "/start",
		HandlerFunc: func(c tele.Context) error {
			return c.Send("Hello, " + c.Sender().FirstName)
		},
	}
}

func (h *CommonHandler) handleHelp() *Handler {
	return &Handler{
		Endpoint: "/help",
		HandlerFunc: func(c tele.Context) error {
			return c.Send(c.Sender().FirstName, " this is help message")
		},
	}
}
