package httpserver

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	core "github.com/elangreza14/minipulsa/api-gateway/core"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func HttpServer(core core.BaseApp) {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}:${ips}]:${port} - ${protocol} - ${time} - ${status} - ${method} - ${url} - ${body} - ${route} \n",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		Next:             nil,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowCredentials: true,
		ExposeHeaders:    "",
		MaxAge:           0}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	authRouter := app.Group("/authentication")
	NewPublicAuthenticationController(authRouter, core.As, core.Log)

	productRouter := app.Group("/product")
	NewPublicProductController(productRouter, core.Ps, core.Log)

	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("Gracefully shutting down...")

	if err := app.Shutdown(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fiber was successful shutdown.")

}
