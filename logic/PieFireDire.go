package logic

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"regexp"
	"strings"
)

func PieFireDire(c *fiber.Ctx) error {
	// สร้าง map สำหรับเก็บข้อมูล
	data := make(map[string]interface{})

	err := json.Unmarshal(c.Body(), &data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Error",
			"description": "Invalid JSON format",
		})
	}

	// validate body message
	if data["message"] == nil || data["message"] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Error",
			"description": "missing or invalid field (\"message\")",
		})
	}
	// logic of PieFireDire
	text := strings.ToLower(data["message"].(string))
	//text := strings.ToLower(string(data))
	//text = strings.ReplaceAll(text, ".", ",")
	//text = strings.ReplaceAll(text, " ", ",")
	//text = strings.ReplaceAll(text, "\n", ",")
	//text = strings.ReplaceAll(text, "\r", ",")
	re := regexp.MustCompile(`[.\s\r\n]+`)
	text = re.ReplaceAllString(text, ",")
	re = regexp.MustCompile(",+")
	text = re.ReplaceAllString(text, ",")
	text = strings.Trim(text, ",")

	words := strings.Split(text, ",")

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "Success",
		"description": "success",
		"result":      wordCount,
	})
}
