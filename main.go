package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {

	engine := html.New("./frontend", ".html")
	engine.Reload(true)

	app := fiber.New()
	app.Use(cors.New())

	// Pass the engine to the Views
	app = fiber.New(fiber.Config{
		Views: engine,
	})

	pusherClient := pusher.Client{
		AppID: "1821998",
		Key: "80fae0a189495728af7c",
		Secret: "1139aa484cfdbc03abfb",
		Cluster: "us2",
		Secure: true,
	}

	app.Get("/chat", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Post("/api/send-message", func(c *fiber.Ctx) (err error) {
        var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		err = pusherClient.Trigger("go-chat-app", "send-message", data)
		if err != nil {
			fmt.Println(err.Error())
		}

		return
    })

	log.Fatalln("Listen: 3000", app.Listen(":3000"))
}
