package controllers

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"net/http"

	"github.com/vimalkuriensam/broker-service/pkg/config"
	"github.com/vimalkuriensam/broker-service/pkg/models"
)

type ItemController interface {
	AddItem(http.ResponseWriter, *http.Request)
	GetItem(http.ResponseWriter, *http.Request)
	UpdateItem(http.ResponseWriter, *http.Request)
	DeleteItem(http.ResponseWriter, *http.Request)
}

type itemController struct{}

func New() ItemController {
	return &itemController{}
}

func (c *itemController) AddItem(w http.ResponseWriter, req *http.Request) {
	cfg := config.GetConfig()
	go cfg.ReadJSON(req)
	data := (<-cfg.DataChan).(config.ReadValue)
	item := models.Items{}
	json.Unmarshal(data.B, &item)
	reply := config.JSONResponse{}
	if err := cfg.Client.Call("ItemCollection.AddItem", item, &reply); err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
	}
	if err := gob.NewDecoder(bytes.NewBuffer(reply.Data.([]byte))).Decode(&item); err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
	}
	reply.Data = item
	cfg.WriteJSON(w, http.StatusCreated, reply.Data, reply.Message)
}

func (c *itemController) GetItem(w http.ResponseWriter, res *http.Request) {}

func (c *itemController) UpdateItem(w http.ResponseWriter, res *http.Request) {}

func (c *itemController) DeleteItem(w http.ResponseWriter, req *http.Request) {}
