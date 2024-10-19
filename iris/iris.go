package iris

import (
	"github.com/kataras/iris/v12"
	"github.com/mvrilo/go-redoc"
)

func New(doc redoc.Redoc) iris.Handler {
	handle := doc.Handler()
	return func(ctx iris.Context) {
		handle(ctx.ResponseWriter(), ctx.Request())
		ctx.Next()
	}
}
