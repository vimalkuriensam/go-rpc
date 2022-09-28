package controllers

import (
	"net/http"

	"github.com/vimalkuriensam/broker-service/pkg/config"
)

type ItemController interface {
	AddItem(http.ResponseWriter, *http.Request)
	GetItem(http.ResponseWriter, *http.Request)
	UpdateItem(http.ResponseWriter, *http.Request)
	DeleteItem(http.ResponseWriter, *http.Request)
}

type itemController struct {
	config *config.Config
}

func New() ItemController {
	return &itemController{
		config: config.GetConfig(),
	}
}

func (c *itemController) AddItem(w http.ResponseWriter, res *http.Request) {

}

func (c *itemController) GetItem(w http.ResponseWriter, res *http.Request) {}

func (c *itemController) UpdateItem(w http.ResponseWriter, res *http.Request) {}

func (c *itemController) DeleteItem(w http.ResponseWriter, req *http.Request) {}
