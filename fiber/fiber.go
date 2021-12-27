package fiberredoc

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/mvrilo/go-redoc"
)

func New(doc redoc.Redoc) fiber.Handler {
	return adaptor.HTTPHandlerFunc(doc.Handler())
}
