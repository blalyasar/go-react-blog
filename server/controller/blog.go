package controller

import (
	"log"

	"github.com/blalyasar/go-react-blog/database"
	"github.com/blalyasar/go-react-blog/model"
	"github.com/gofiber/fiber/v2"
)

func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog List",
	}

	db := database.DBConn

	var records []model.Blog
	db.Find(&records)
	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)
	// c.JSON()
}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog Create",
	}

	record := new(model.Blog)
	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request")
		context["statusText"] = ""
		context["msg"] = "Something is wrong"
	}

	result := database.DBConn.Create(record)
	if result.Error != nil {
		log.Println("Error in saving data")
		context["statusText"] = ""
		context["msg"] = "Something is wrong"
	}

	context["msg"] = "Record is saved successfuly."
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog Update",
	}
	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog Delete from given id",
	}
	c.Status(200)
	return c.JSON(context)
}
