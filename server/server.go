package main

import (
	"log"

	"github.com/blalyasar/go-react-blog/database"
	"github.com/blalyasar/go-react-blog/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// https://www.youtube.com/watch?v=PPsTNQpT3e0&list=PL0wd2rfixHH23Ej7qVSrbgr-yhI6iSC05&index=1
// https://docs.gofiber.io/
// go get github.com/gofiber/fiber/v2

func init() {
	if err :=godotenv.Load(".env");err !=nil{
		log.Fatal("Error in loading .env file.")
	}
	database.ConnectDB()
}

func main() {
	database.ConnectDB()

	sqlDb, err := database.DBConn.DB()
	if err != nil {
		panic("Error in sql connection")
	}
	defer sqlDb.Close()

	app := fiber.New()
	// app.Get("/", func(c *fiber.Ctx) error {

	// 	return c.JSON(
	// 		fiber.Map{"message": "welcome first blog"},
	// 	)
	// return c.SendString("Hello World")
	// })
	router.SetupRoutes(app)
	app.Listen(":8000")
}
