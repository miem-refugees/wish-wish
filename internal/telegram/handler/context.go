package handler

import "go.uber.org/zap"

type Context struct {
	logger  *zap.Logger
	traceID string
}

func NewContext(logger *zap.Logger, traceID string) *Context {
	return &Context{
		logger:  logger,
		traceID: traceID,
	}
}
