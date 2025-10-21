package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const uploadsDir = "./uploads"

// UploadFile handles file upload
func UploadFile(c *fiber.Ctx) error {
	// Get file from form
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No file uploaded",
		})
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	filePath := filepath.Join(uploadsDir, filename)

	// Save file
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}

	// Generate URL untuk akses file
	protocol := "http"
	if c.Protocol() == "https" {
		protocol = "https"
	}

	host := c.Hostname()
	if host == "" {
		host = "localhost:3000"
	}

	fileURL := fmt.Sprintf("%s://%s/files/%s", protocol, host, filename)
	storageURL := fmt.Sprintf("%s://%s/storage/%s", protocol, host, filename)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":       "File uploaded successfully",
		"filename":      filename,
		"original_name": file.Filename,
		"size":          file.Size,
		"url":           fileURL,
		"storage_url":   storageURL,
		"uploaded_at":   time.Now().Format(time.RFC3339),
	})
}

// DownloadFile handles file download
func DownloadFile(c *fiber.Ctx) error {
	filename := c.Params("filename")
	filePath := filepath.Join(uploadsDir, filename)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "File not found",
		})
	}

	// Send file
	return c.SendFile(filePath)
}

// ListFiles lists all uploaded files
func ListFiles(c *fiber.Ctx) error {
	files, err := os.ReadDir(uploadsDir)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read directory",
		})
	}

	var fileList []fiber.Map
	protocol := "http"
	if c.Protocol() == "https" {
		protocol = "https"
	}

	host := c.Hostname()
	if host == "" {
		host = "localhost:3000"
	}

	for _, file := range files {
		if !file.IsDir() {
			info, err := file.Info()
			if err != nil {
				continue
			}

			filename := file.Name()
			fileURL := fmt.Sprintf("%s://%s/files/%s", protocol, host, filename)

			fileList = append(fileList, fiber.Map{
				"filename": filename,
				"size":     info.Size(),
				"modified": info.ModTime().Format(time.RFC3339),
				"url":      fileURL,
			})
		}
	}

	return c.JSON(fiber.Map{
		"total_files": len(fileList),
		"files":       fileList,
	})
}

// DeleteFile handles file deletion
func DeleteFile(c *fiber.Ctx) error {
	filename := c.Params("filename")
	filePath := filepath.Join(uploadsDir, filename)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "File not found",
		})
	}

	// Delete file
	if err := os.Remove(filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete file",
		})
	}

	return c.JSON(fiber.Map{
		"message":  "File deleted successfully",
		"filename": filename,
	})
}
