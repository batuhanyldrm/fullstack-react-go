package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	repository := NewRepository()
	service := NewService(repository)
	api := NewApi(&service)
	app := SetupApp(&api)
	app.Listen(":3001")

}

func SetupApp(api *Api) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Get("/blogs", api.GetBlogsHandler)
	app.Post("/blogs", api.PostBlogsHandler)
	app.Delete("/blogs/:id", api.DeleteBlogsHandler)
	app.Put("/blogs/:id", api.UpdateBlogsHandler)
	app.Get("/blog/:id", api.GetBlogHandler)

	return app
}
