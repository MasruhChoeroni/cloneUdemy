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
		// 	// Select all Course(s) from database
		// 	rows, err := db.Query("SELECT * FROM courses order by id")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		// 	defer rows.Close()
		// 	result := Courses{}

		// 	for rows.Next() {
		// 		course := Course{}
		// 		if err := rows.Scan(&course.ID, &course.CreatorId, &course.Title, &course.SubTitle, &course.Description, &course.Language, &course.Price, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt); err != nil {
		// 			return err // Exit if we get an error
		// 		}

		// 		// Append Course to Courses
		// 		result.Courses = append(result.Courses, course)
		// 	}
		// 	// Return Courses in JSON format
		return c.JSON(result)
	})
}
