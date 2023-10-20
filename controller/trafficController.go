package controller

import (
	"embedded/infrastructure"
	"embedded/model"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type TrafficController interface {
	GetTrafficLight(w http.ResponseWriter, r *http.Request)
	UpdateTrafficLight(w http.ResponseWriter, r *http.Request)
}

type trafficController struct {
	db *gorm.DB
}

func (c *trafficController) GetTrafficLight(w http.ResponseWriter, r *http.Request) {
	var res model.Response
	var trafficLight model.TrafficLight

	// Get the traffic light info from the database (id = 1)
	if err := c.db.Where("id = 1").First(&trafficLight).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		res = model.Response{
			Data:    nil,
			Message: "Internal server error: " + err.Error(),
			Success: false,
		}
		return
	}

	// Return the response
	res = model.Response{
		Data:    trafficLight,
		Message: "Get traffic light successfully",
		Success: true,
	}
	render.JSON(w, r, res)
	return
}

func (c *trafficController) UpdateTrafficLight(w http.ResponseWriter, r *http.Request) {
	var res model.Response
	var trafficLight model.TrafficLight
	// Get the traffic light info from the request
	if err := json.NewDecoder(r.Body).Decode(&trafficLight); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		res = model.Response{
			Data:    nil,
			Message: "Bad request",
			Success: false,
		}
		return
	}

	// Update the traffic light info to the database
	if err := c.db.Model(&trafficLight).Where("id = 1").Updates(trafficLight).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		res = model.Response{
			Data:    nil,
			Message: "Internal server error: " + err.Error(),
			Success: false,
		}
		return
	}

	// Return the response
	res = model.Response{
		Data:    trafficLight,
		Message: "Update traffic light successfully",
		Success: true,
	}
	render.JSON(w, r, res)
	return
}

func NewTrafficController() TrafficController {
	return &trafficController{
		db: infrastructure.GetDB(),
	}
}
