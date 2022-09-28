package config

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func (cfg *Config) ReadJSON(req *http.Request) error {
	data := ReadValue{
		B: []byte(""),
		D: nil,
	}
	if body, err := ioutil.ReadAll(req.Body); err == nil {
		if err = json.Unmarshal([]byte(body), &data.D); err == nil {
			data.B = []byte(body)
			cfg.DataChan <- data
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}

func (cfg *Config) WriteJSON(w http.ResponseWriter, status int, data interface{}, msg string, headers ...http.Header) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	cfg.Response.Data = data
	cfg.Response.Message = msg
	if b_data, err := json.Marshal(cfg.Response); err == nil {
		w.Write(b_data)
	}
}

func (cfg *Config) ErrorJSON(w http.ResponseWriter, path string, reason string, status ...int) {
	errorStatus := http.StatusBadRequest
	if len(status) > 0 {
		errorStatus = status[0]
	}
	cfg.Logger.Println("error-reason: ", reason)
	cfg.ErrorResponse = &ErrorResponse{
		Code:      errorStatus,
		Path:      path,
		Reason:    reason,
		Timestamp: time.Now(),
	}
	cfg.WriteJSON(w, errorStatus, cfg.ErrorResponse, "Error")
}
