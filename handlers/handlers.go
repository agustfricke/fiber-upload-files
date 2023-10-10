package handlers

import (
	"fmt"

	"path/filepath"

	"github.com/agustfricke/fiber-upload-files/database"
	"github.com/agustfricke/fiber-upload-files/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


func Upload(c *fiber.Ctx) error {
    file, err := c.FormFile("upload")
    if err != nil {
        return err
    }
    id := uuid.New()
    ext := filepath.Ext(file.Filename)
    newFilename := fmt.Sprintf("%s%s", id, ext)
    c.SaveFile(file, fmt.Sprintf("public/uploads/%s", newFilename))
    fullPath := fmt.Sprintf("uploads/%s", newFilename)
    fmt.Printf("Archivo guardado en: %s\n", fullPath)

    newFile := models.File{
        FilePath: fullPath,
    }

    if err := database.DB.Create(&newFile).Error; err != nil {
        return err
    }

    return c.Render("home", fiber.Map{})
}

func Home(c *fiber.Ctx) error {
	files := []models.File{}
	database.DB.Find(&files)
	return c.Render("home", fiber.Map{
    "Files":    files,
	})
}

