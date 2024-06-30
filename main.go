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
		Prefork:       false,
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
			"version": "2.0.0",
			"author":  "DscInflux",
			"links": fiber.Map{
				"status":  "https://dscinflux.instatus.com/",
				"docs":    "https://docs.dscinflux.xyz/",
				"support": "https://discord.gg/invite/RPCtG7Em8g",
			},
		})
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://dscinflux.xyz,http://localhost:3000",
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

	store := session.New()

	app.Use(func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		c.Locals("session", sess)
		return c.Next()
	})
	// Middleware: OAuth2 Configuration
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("authConfig", &oauth2.Config{
			RedirectURL:  config.Client.Callback,
			ClientID:     config.Client.Id,
			ClientSecret: os.Getenv("ClientSecret"),
			Scopes:       []string{discord.ScopeIdentify},
			Endpoint:     discord.Endpoint,
		})

		return c.Next()
	})

	//Routes
	app.Get("/v1/auth/login", routes.Login)
	app.Get("/v1/auth/callback", routes.Callback)
	app.Get("/v1/auth/logout", routes.Logout)
	app.Get("/v1/auth/@me", routes.GetCurrentUser)
	app.Get("/v1/socials", routes.GetSocials)
	app.Get("/v1/sort", routes.GetSortings)
	app.Get("/v1/entity/:id", routes.GetEntity)
	app.Get("/v1/entities", routes.Entities)
	app.Get("/v1/roles", routes.GetRoles)
	app.Get("/v1/skills", routes.GetSkills)
	app.Get("/v1/languages", routes.GetLanguages)
	app.Post("/v1/entity", routes.NewEntity)
	app.Post("/v1/entity/:id/like", routes.Like)
	app.Post("/v1/entity/:id/unlike", routes.Unlike)
	app.Get("/v1/random", routes.GetRandomEntity)
	app.Get("/staff", routes.GetStaffUsers)
	app.Get("/partner", routes.GetPartners)
	app.Get("/banners/:id", routes.GetBanner)
	app.Get("/avatars/:id", routes.GetAvatar)
	app.Get("/verifiednum", routes.VerifiedNum)
	app.Get("/staffnum", routes.StaffNum)
	app.Get("/usernum", routes.UserNum)
	app.Get("/profilenum", routes.ProfileNum)

	app.Listen(":" + config.Web.Port)
}
