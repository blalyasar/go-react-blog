package router

import (
	"github.com/blalyasar/go-react-blog/controller"
	"github.com/gofiber/fiber/v2"
)

// SETUP ROUTE INFO
func SetupRoutes(app *fiber.App){
	// list get
	// add post
	// update put
	// delete delete
	app.Get("/",controller.BlogList)
	app.Post("/",controller.BlogCreate)
	app.Put("/:id",controller.BlogUpdate)
	app.Delete("/:id",controller.BlogDelete)

}