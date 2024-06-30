package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UserNum(c *fiber.Ctx) error {
	db, ok := c.Locals("db").(*mongo.Client)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not available",
		})
	}

	userCollection := db.Database("dscinflux").Collection("users")

	countOptions := options.Count()
	totaluser, err := userCollection.CountDocuments(context.Background(), bson.M{}, countOptions)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error counting users",
		})
	}

	return c.JSON(fiber.Map{"total_user": totaluser})
}

func ProfileNum(c *fiber.Ctx) error {
	db, ok := c.Locals("db").(*mongo.Client)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not available",
		})
	}

	userCollection := db.Database("dscinflux").Collection("entities")

	countOptions := options.Count()
	ProfileNum, err := userCollection.CountDocuments(context.Background(), bson.M{}, countOptions)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error counting profiles",
		})
	}

	return c.JSON(fiber.Map{"total_profiles": ProfileNum})
}

func StaffNum(c *fiber.Ctx) error {
	db, ok := c.Locals("db").(*mongo.Client)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not available",
		})
	}

	userCollection := db.Database("dscinflux").Collection("entities")

	// Define a filter to count documents where staff is true
	filter := bson.M{"staff": "true"}

	// Count documents in the collection
	totalStaff, err := userCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error counting verified users",
		})
	}

	return c.JSON(fiber.Map{"total_staff": totalStaff})
}

func VerifiedNum(c *fiber.Ctx) error {
	db, ok := c.Locals("db").(*mongo.Client)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not available",
		})
	}

	userCollection := db.Database("dscinflux").Collection("entities")

	// Define a filter to count documents where staff is true
	filter := bson.M{"isVerified": true}

	// Count documents in the collection
	TotalVerified, err := userCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error counting verified users",
		})
	}

	return c.JSON(fiber.Map{"total_verified": TotalVerified})
}
