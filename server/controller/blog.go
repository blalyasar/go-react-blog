package controller

import (
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

	c.Status(200)
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
