package routes

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.dscinflux.xyz/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetPartners retrieves all partners from the database
func GetPartners(c *fiber.Ctx) error {
	db := c.Locals("db").(*mongo.Client)
	if db == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not available",
		})
	}

	partnerCollection := db.Database("dscinflux").Collection("Partners")

	// Define a slice to store partners
	var partners []types.Partner

	// Fetch all partners from the collection
	cursor, err := partnerCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error fetching partners:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch partners",
		})
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor and decode partners
	for cursor.Next(context.Background()) {
		var partner types.Partner
		if err := cursor.Decode(&partner); err != nil {
			log.Println("Error decoding partner:", err)
			continue // Skip this partner and proceed to the next one
		}
		partners = append(partners, partner)
	}

	// Check for errors during cursor iteration
	if err := cursor.Err(); err != nil {
		log.Println("Error iterating over partners:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to iterate over partners",
		})
	}

	// Return partners as JSON
	return c.JSON(partners)
}
