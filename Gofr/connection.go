package main

import "gofr.dev/pkg/gofr"

type Student struct {
	ROLL  int    `json:"roll"`
	Name  string `json:"name"`
	Class int    `json:"class"`
}

func main() {
	// initialise gofr object
	app := gofr.New()

	app.GET("/welcome", func(ctx *gofr.Context) (interface{}, error) {

		return "Hello students! Welcome to the classroom", nil
	})

	app.POST("/student/{name}", func(ctx *gofr.Context) (interface{}, error) {
		name := ctx.PathParam("name")

		// Inserting a customer row in database using SQL
		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO customers (name) VALUES (?)", name)

		return nil, err
	})

	app.GET("/student", func(ctx *gofr.Context) (interface{}, error) {
		var students []Student

		// Getting the customer from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM students")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var student Student
			if err := rows.Scan(&student.ROLL, &student.Name, &student.Class); err != nil {
				return nil, err
			}

			students = append(students, student)
		}

		// return the customer
		return students, nil
	})

	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}
