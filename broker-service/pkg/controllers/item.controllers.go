package controllers

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vimalkuriensam/broker-service/pkg/config"
	"github.com/vimalkuriensam/broker-service/pkg/models"
	"github.com/vimalkuriensam/broker-service/pkg/services"
)

type ItemController interface {
	AddItem(http.ResponseWriter, *http.Request)
	GetItem(http.ResponseWriter, *http.Request)
	UpdateItem(http.ResponseWriter, *http.Request)
	DeleteItem(http.ResponseWriter, *http.Request)
}

type itemController struct {
	fields map[string][]string
}

func New() ItemController {
	return &itemController{
		fields: models.ItemAcceptableFields,
	}
}

func (c *itemController) AddItem(w http.ResponseWriter, req *http.Request) {
	cfg := config.GetConfig()
	data, err := services.ReadRequest(req, c.fields["create"])
	if err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusBadRequest)
		return
	}
	item := models.Items{}
	json.Unmarshal(data.B, &item)
	reply := config.JSONResponse{}
	if err := cfg.Client.Call("ItemCollection.AddItem", item, &reply); err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := gob.NewDecoder(bytes.NewBuffer(reply.Data.([]byte))).Decode(&item); err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
		return
	}
	reply.Data = item
	cfg.WriteJSON(w, http.StatusCreated, reply.Data, reply.Message)
}

func (c *itemController) GetItem(w http.ResponseWriter, req *http.Request) {
	item := models.Items{}
	reply := config.JSONResponse{}
	cfg := config.GetConfig()
	id := chi.URLParam(req, "id")
	if len(id) == 0 {
		cfg.ErrorJSON(w, req.URL.Path, "user id is not provided", http.StatusBadRequest)
		return
	}
	if err := cfg.Client.Call("ItemCollection.GetItem", id, &reply); err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := gob.NewDecoder(bytes.NewBuffer(reply.Data.([]byte))).Decode(&item); err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
		return
	}
	reply.Data = item
	cfg.WriteJSON(w, http.StatusOK, reply.Data, reply.Message)
}

func (c *itemController) UpdateItem(w http.ResponseWriter, req *http.Request) {
	item := models.Items{}
	reply := config.JSONResponse{}
	cfg := config.GetConfig()
	id := chi.URLParam(req, "id")
	if len(id) == 0 {
		cfg.ErrorJSON(w, req.URL.Path, "user id is not provided", http.StatusBadRequest)
	}
	data, err := services.ReadRequest(req, c.fields["update"])
	if err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusBadRequest)
	}
	if err := json.Unmarshal(data.B, &item); err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
	}
	updateInput := &models.UpdateItemInput{
		ID:         id,
		UpdateItem: item,
	}
	updateResult := models.UpdateItemResponse{}
	if err := cfg.Client.Call("ItemCollection.UpdateItem", updateInput, &reply); err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
	}
	if err := gob.NewDecoder(bytes.NewBuffer(reply.Data.([]byte))).Decode(&updateResult); err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
	}
	reply.Data = updateResult
	cfg.WriteJSON(w, http.StatusOK, reply.Data, reply.Message)
}

func (c *itemController) DeleteItem(w http.ResponseWriter, req *http.Request) {
	deleteResponse := models.DeleteItem{}
	reply := config.JSONResponse{}
	cfg := config.GetConfig()
	id := chi.URLParam(req, "id")
	if len(id) == 0 {
		cfg.ErrorJSON(w, req.URL.Path, "user id is not provided", http.StatusBadRequest)
	}
	if err := cfg.Client.Call("ItemCollection.DeleteItem", id, &reply); err != nil {
		fmt.Println("here1")
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := gob.NewDecoder(bytes.NewBuffer(reply.Data.([]byte))).Decode(&deleteResponse); err != nil {
		fmt.Println("here2")
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
		return
	}
	reply.Data = deleteResponse
	cfg.WriteJSON(w, http.StatusOK, reply.Data, reply.Message)
}
