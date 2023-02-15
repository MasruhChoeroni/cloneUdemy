package services

import (
	"cloneUdemy/models"
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

var db *sql.DB

// Database settings
const (
	host     = "localhost"
	port     = 5432 // Default port
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "udemy_clone"
)

func connect() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func CourseServiceGet(c *fiber.Ctx) (*models.Courses, error) {
	if err := connect(); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM courses order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := models.Courses{}

	for rows.Next() {
		course := models.Course{}
		if err := rows.Scan(&course.ID, &course.CreatorId, &course.Title, &course.SubTitle, &course.Description, &course.Language, &course.Price, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt); err != nil {
			return nil, err // Exit if we get an error
		}
		// Append Course to Courses
		result.Courses = append(result.Courses, course)
	}
	// Return Courses in JSON format
	return &result, nil
}
