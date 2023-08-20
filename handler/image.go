package handler

import (
	"fmt"
	"log"
	"strings"
	"path/filepath"
	"github.com/google/uuid"	
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kkkiikkk/go-jwt/model"
	"github.com/kkkiikkk/go-jwt/config"
	"github.com/kkkiikkk/go-jwt/database"
)

func isImageExtension(extension string) bool {
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	
	for _, ext := range imageExtensions {
		if ext == extension {
			return true
		}
	}

	return false
}

// Upload image
func UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	fileExt := strings.ToLower(filepath.Ext(file.Filename))

	if !isImageExtension(fileExt) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": 400, "message": "The file extension must be one of the following: jpeg, jpg, png, gif", "data": nil})
	}

	uniqueId := uuid.New()

    filename := strings.Replace(uniqueId.String(), "-", "", -1)

    image := fmt.Sprintf("%s%s", filename, fileExt)

    imageUrl := fmt.Sprintf("http://localhost:" + config.Config("PORT") + "/images/%s", image)

	db := database.DB
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))

	newImage := model.Image{
		ImageUrl: imageUrl,
		ImagePath: image,
		UserId: userId,
	}
	
	result := db.Create(&newImage)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to create image",
        })
    }

	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

    if err != nil {
        log.Println("image save error --> ", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
    }

    return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": newImage})
}

// Get all images
func GetImages(c *fiber.Ctx) error {
	db := database.DB

	var images []model.Image
	db.Find(&images)

	return c.JSON(fiber.Map{"status": 200, "message": "All images", "data": images})
}