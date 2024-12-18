package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.dscinflux.xyz/configuration"
	"go.dscinflux.xyz/middlewares"
	"go.dscinflux.xyz/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRandomEntity(c *fiber.Ctx) error {
	DConfig := configuration.GetConfig()
	var collectionData = DConfig.Collection

	var db *mongo.Client = c.Locals("db").(*mongo.Client)
	users := db.Database("dscinflux").Collection(collectionData)

	var user types.Entity

	err := users.FindOne(context.Background(), bson.M{}).Decode(&user)

	if err != nil {
		return c.Status(404).JSON(middlewares.NewError("User not found"))
	}

	return FindUser(user.URL, c)
}
