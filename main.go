package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
	"github.com/ravener/discord-oauth2"
	"go.dscinflux.xyz/configuration"
	"go.dscinflux.xyz/database"
	"go.dscinflux.xyz/routes"
	"golang.org/x/oauth2"
)

func main() {
	// Load ENV
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	//Fiber stuff
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "DscInflux",
		AppName:       "A website used to list a users profile.",
	})

	config := configuration.GetConfig()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("config", config)
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
			"version": "1.0.0",
			"author":  "DscInflux",
			"links": fiber.Map{
				"status":  "https://dscinflux.instatus.com/",
				"docs":    "https://docs.dscinflux.xyz/",
				"support": "https://discord.gg/invite/RPCtG7Em8g",
			},
		})
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://dscinflux.xyz/",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	// Middleware: Database Connection
	db, err := database.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Disconnect(nil)
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	v1 := app.Group("/v1")

	store := session.New()

	v1.Use(func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		c.Locals("session", sess)
		return c.Next()
	})

	v1.Use(func(c *fiber.Ctx) error {
		c.Locals("authConfig", &oauth2.Config{
			RedirectURL:  config.Client.Callback,
			ClientID:     config.Client.Id,
			ClientSecret: os.Getenv("ClientSecret"),
			Scopes:       []string{discord.ScopeIdentify},
			Endpoint:     discord.Endpoint,
		})

		return c.Next()
	})

	v1.Get("/socials", routes.GetSocials)
	v1.Get("/sort", routes.GetSortings)
	v1.Get("/auth/login", routes.Login)
	v1.Get("/auth/callback", routes.Callback)
	v1.Get("/auth/logout", routes.Logout)
	v1.Get("/auth/@me", routes.GetCurrentUser)
	v1.Get("/entity/:id", routes.GetEntity)
	v1.Get("/entities", routes.Entities)
	v1.Get("/roles", routes.GetRoles)
	v1.Get("/skills", routes.GetSkills)
	v1.Get("/languages", routes.GetLanguages)
	v1.Post("/entity", routes.NewEntity)
	v1.Post("/entity/:id/like", routes.Like)
	v1.Post("/entity/:id/unlike", routes.Unlike)
	v1.Get("/random", routes.GetRandomEntity)
	v1.Get("/banners/:id", routes.GetBanner)
	v1.Get("/avatars/:id", routes.GetAvatar)

	app.Listen(":" + config.Web.Port)
}
