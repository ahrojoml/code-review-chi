package handler

import (
	"app/internal"
	"app/internal/repository"
	"app/internal/service"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = vehicleToVehicleJSON(value)
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody VehicleJSON
		if err := request.JSON(r, &reqBody); err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		newVehicle := vehicleJSONToVehicle(reqBody)

		vehicle, err := h.sv.Add(newVehicle)
		if err != nil {
			switch err.(type) {
			case *internal.VehicleValidationError:
				response.JSON(w, http.StatusBadRequest, err.Error())
			case *repository.VehicleAlreadyExistsError:
				response.JSON(w, http.StatusConflict, err.Error())
			default:
				response.JSON(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := vehicleToVehicleJSON(vehicle)

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetByBrandAndYears() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		if brand == "" {
			response.Error(w, http.StatusBadRequest, "brand field empty")
			return
		}

		startYear, err := strconv.Atoi(chi.URLParam(r, "start_year"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid starting year")
			return
		}

		endYear, err := strconv.Atoi(chi.URLParam(r, "end_year"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid ending year")
			return
		}

		vehicles, err := h.sv.FindByBrandAndYears(brand, startYear, endYear)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range vehicles {
			data[key] = vehicleToVehicleJSON(value)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) UpdateMaxSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		var reqBody VehicleJSON
		if err := request.JSON(r, &reqBody); err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		maxSpeed := reqBody.MaxSpeed

		vehicle, err := h.sv.UpdateMaxSpeed(id, maxSpeed)
		if err != nil {
			switch err.(type) {
			case *repository.VehicleNotFoundError:
				response.Error(w, http.StatusNotFound, err.Error())
			case *service.FieldValidationError:
				response.Error(w, http.StatusBadRequest, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := vehicleToVehicleJSON(vehicle)

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
