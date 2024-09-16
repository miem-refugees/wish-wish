package handler

import tele "gopkg.in/telebot.v3"

type Handler struct {
	Endpoint    string
	HandlerFunc tele.HandlerFunc
	Middleware  []tele.MiddlewareFunc
}
