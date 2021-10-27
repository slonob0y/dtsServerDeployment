package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	UserID    string `json:"userId"`
	Username  string `json:"username"`
	Followers int    `json:"followers"`
}

var data = []User{
	{UserID: "sammy", Username: "SammyShark", Followers: 987},
	{UserID: "jesse", Username: "JesseOctopus", Followers: 432},
	{UserID: "drew", Username: "DrewSquid", Followers: 321},
	{UserID: "jamie", Username: "JamieMantisShrimp", Followers: 654},
}

func getData(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(data)
}

func getDataByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	for _, follow := range data {
		if follow.Username == username {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				username: fiber.Map{
					"followers": follow.Followers,
				},
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Record not found",
	})
}

//GetDataFollowerById
func getDataById(c *fiber.Ctx) error {
	userId := c.Params("userId")

	for _, follow := range data {
		if follow.UserID == userId {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				userId: follow,
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server UP!")
	})

	app.Get("/follower", getData)
	app.Get("/follower/:username", getDataByUsername)
	app.Get("/:userId/detail", getDataById)

	err := app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		fmt.Print("error")
	}

}
