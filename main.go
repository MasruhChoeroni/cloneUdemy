package main

import (
	"database/sql"
	"fmt"
	"log"

	"cloneUdemy/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

// Database instance
var db *sql.DB

// Database settings
const (
	host     = "localhost"
	port     = 5432 // Default port
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "udemy_clone"
)

// Course struct
type Course struct {
	ID          string `json:"id"`
	CreatorId   string `json:"creator_id"`
	Title       string `json:"title"`
	SubTitle    string `json:"subtitle"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Price       string `json:"price"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

// Courses struct
type Courses struct {
	Courses []Course `json:"courses"`
}

// Connect function
func Connect() error {
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

func main() {
	// Connect with database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})

	routes.R_course(app.Group("/course"))

	// Get all records from postgreSQL
	app.Get("/course", func(c *fiber.Ctx) error {
		// Select all Course(s) from database
		rows, err := db.Query("SELECT * FROM courses order by id")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer rows.Close()
		result := Courses{}

		for rows.Next() {
			course := Course{}
			if err := rows.Scan(&course.ID, &course.CreatorId, &course.Title, &course.SubTitle, &course.Description, &course.Language, &course.Price, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt); err != nil {
				return err // Exit if we get an error
			}

			// Append Course to Courses
			result.Courses = append(result.Courses, course)
		}
		// Return Courses in JSON format
		return c.JSON(result)
	})

	// // Add record into postgreSQL
	// app.Post("/course", func(c *fiber.Ctx) error {
	// 	// New Course struct
	// 	u := new(Courses)

	// 	// Parse body into struct
	// 	if err := c.BodyParser(u); err != nil {
	// 		return c.Status(400).SendString(err.Error())
	// 	}

	// 	// Insert Courses into database
	// 	res, err := db.Query("INSERT INTO courses (name, salary, age)VALUES ($1, $2, $3)", u.Name, u.Salary, u.Age)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Print result
	// 	log.Println(res)

	// 	// Return Courses in JSON format
	// 	return c.JSON(u)
	// })

	// // Update record into postgreSQL
	// app.Put("/course", func(c *fiber.Ctx) error {
	// 	// New Courses struct
	// 	u := new(Course)

	// 	// Parse body into struct
	// 	if err := c.BodyParser(u); err != nil {
	// 		return c.Status(400).SendString(err.Error())
	// 	}

	// 	// Update Course into database
	// 	res, err := db.Query("UPDATE courses SET name=$1,salary=$2,age=$3 WHERE id=$5", u.Name, u.Salary, u.Age, u.ID)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Print result
	// 	log.Println(res)

	// 	// Return Course in JSON format
	// 	return c.Status(201).JSON(u)
	// })

	// // Delete record from postgreSQL
	// app.Delete("/course", func(c *fiber.Ctx) error {
	// 	// New Course struct
	// 	u := new(Course)

	// 	// Parse body into struct
	// 	if err := c.BodyParser(u); err != nil {
	// 		return c.Status(400).SendString(err.Error())
	// 	}

	// 	// Delete Course from database
	// 	res, err := db.Query("DELETE FROM courses WHERE id = $1", u.ID)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Print result
	// 	log.Println(res)

	// 	// Return Course in JSON format
	// 	return c.JSON("Deleted")
	// })

	log.Fatal(app.Listen(":3000"))
}
