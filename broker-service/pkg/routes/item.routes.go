package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/vimalkuriensam/broker-service/pkg/controllers"
)

var itemController controllers.ItemController = controllers.New()

func itemRoutes(r chi.Router) {
	r.Post("/addItem", itemController.AddItem)
	r.Get("/getItem/{id}", itemController.GetItem)
	r.Patch("/updateItem/{id}", itemController.UpdateItem)
	r.Delete("/deleteItem/{id}", itemController.DeleteItem)
}
