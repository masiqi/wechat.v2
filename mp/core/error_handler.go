package core

import (
	"log"
	"os"

	"github.com/valyala/fasthttp"
)

type ErrorHandler interface {
	ServeError(*fasthttp.RequestCtx, error)
}

var DefaultErrorHandler ErrorHandler = ErrorHandlerFunc(defaultErrorHandlerFunc)

type ErrorHandlerFunc func(*fasthttp.RequestCtx, error)

func (fn ErrorHandlerFunc) ServeError(ctx *fasthttp.RequestCtx, err error) {
	fn(ctx, err)
}

var errorLogger = log.New(os.Stderr, "[WECHAT_ERROR] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)

func defaultErrorHandlerFunc(ctx *fasthttp.RequestCtx, err error) {
	errorLogger.Output(3, err.Error())
}
