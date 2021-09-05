package main

import (
	"go-gearbox/middleware"

	"github.com/gogearbox/gearbox"
)

var (
	newMiddleware middleware.Middleware = middleware.ImplMiddleware()
)

func main() {
	// Setup gearbox
	server := gearbox.New()

	// create a logger middleware
	server.Use(newMiddleware.Logger)

	// Define your handlers
	server.Get("/", func(ctx gearbox.Context) {
		ctx.SendString("Hello World!")
	})

	// handle file uploader
	server.Post("/hello", func(ctx gearbox.Context) {
		form, err := ctx.Context().Request.MultipartForm()
		if err != nil {
			ctx.SendJSON(map[string]interface{}{
				"error": err.Error(),
				"code":  500,
			})
		}

		fileHeader := form.File["sasa"][0]

		file, err := fileHeader.Open()
		if err != nil {
			ctx.SendJSON(map[string]interface{}{
				"error": err.Error(),
				"code":  500,
			})
		}
		_ = file
	})

	// Start service
	server.Start(":3000")
}

// logger code from : https://github.com/AubSs/fasthttplogger
