package main

import (
	"fmt"
	"log"
	"os"
	"storage/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Buat folder uploads jika belum ada
	uploadsDir := "./uploads"
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		log.Fatal("Error creating uploads directory:", err)
	}

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // 100 MB max file size
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Simple Object Storage API",
			"version": "1.0.0",
		})
	})

	// Upload file endpoint
	app.Post("/upload", handlers.UploadFile)

	// Download file endpoint
	app.Get("/files/:filename", handlers.DownloadFile)

	// List files endpoint
	app.Get("/files", handlers.ListFiles)

	// Delete file endpoint
	app.Delete("/files/:filename", handlers.DeleteFile)

	// Serve static files (untuk akses langsung via browser)
	app.Static("/storage", "./uploads")

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "50001"
	}

	// Start server
	log.Printf("Server running on port %s", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
