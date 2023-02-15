package routes

import (
	"cloneUdemy/services"

	"github.com/gofiber/fiber/v2"
)

func R_course(router fiber.Router) {
	router.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("respond with routes courses")
	})

	router.Get("/course", func(c *fiber.Ctx) error {
		result, err := services.CourseServiceGet(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(result)
	})
}
