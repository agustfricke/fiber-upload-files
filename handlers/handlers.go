package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
  "path/filepath"
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
	fullPath := fmt.Sprintf("public/uploads/%s", newFilename)
	fmt.Printf("Archivo guardado en: %s\n", fullPath)
	return c.Render("home", fiber.Map{})
}

func Home(c *fiber.Ctx) error {
	  return c.Render("home", fiber.Map{})
}

