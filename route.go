package main

import (
	_ "github.com/lib/pq"
)

func route() {
	// Get all records from postgreSQL

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

}
