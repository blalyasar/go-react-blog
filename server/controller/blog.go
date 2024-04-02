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

	// http://localhost:8000/2
	id := c.Params("id")
	var record model.Blog
	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found")
		context["msg"] = "Record not found"
		c.Status(400)
		return c.JSON(context)
	}
	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request")
	}
	result := database.DBConn.Save(record)

	if result.Error != nil {
		log.Println("error in saving data")
	}
	context["msg"] = "record updated successfully"
	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")
	var record model.Blog
	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not Found")
		context["msg"] = "Record not Found"
		c.Status(400)
		return c.JSON(context)
	}
	result := database.DBConn.Delete(record)
	if result.Error != nil {
		context["msg"] = "Something is wrong"
		context["statusText"] = "Something is wrong"
		c.Status(400)
		return c.JSON(context)
	}

	context["statusText"] = "Ok"
	context["msg"] = "Record deleted successfuly"
	c.Status(200)
	return c.JSON(context)
}
