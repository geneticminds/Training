package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=admin dbname=User sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello Dhebug")
	})

	e.POST("/register", func(c echo.Context) error {
		var u User
		if err := c.Bind(&u); err != nil {
			return err
		}
		sqlStatement := "INSERT INTO users (email, fname, lname) VALUES ($1, $2, $3)"
		res, err := db.Query(sqlStatement, u.Email, u.Fname, u.Lname)
		if err != nil {
			errJson := ErrorStruct{
				Error: map[string]string{
					"status":  "400",
					"message": "Bad Request",
				},
			}
			return c.JSON(http.StatusBadRequest, errJson)
		} else {
			fmt.Println(res)
			c.Response().Header().Set("Content-Type", "application/json")
			return c.JSON(http.StatusOK, u)
		}
		return c.String(http.StatusOK, "ok")
	})

	log.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}

type User struct {
	UserID   string `json:"userID"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Password string `json:"password"`
}

type ErrorStruct struct {
	Error map[string]string
}
