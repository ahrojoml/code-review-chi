package service

import (
	"app/internal"
	"app/internal/repository"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) Add(vehicle internal.Vehicle) (internal.Vehicle, error) {
	if err := vehicle.IsValid(); err != nil {
		return internal.Vehicle{}, err
	}
	return s.rp.Add(vehicle)
}

func (s *VehicleDefault) FindByBrandAndYears(brand string, start, end int) (map[int]internal.Vehicle, error) {
	if brand == "" {
		return map[int]internal.Vehicle{}, NewFieldValidationError("brand must not be empty")
	}

	if start <= 0 {
		return map[int]internal.Vehicle{}, NewFieldValidationError("start year must be greater than zero")
	}

	if end <= 0 {
		return map[int]internal.Vehicle{}, NewFieldValidationError("end year must be greater than zero")
	}
	return s.rp.FindByBrandAndYears(brand, start, end)
}

func (s *VehicleDefault) UpdateMaxSpeed(id int, maxSpeed float64) (internal.Vehicle, error) {
	if maxSpeed <= 0 {
		return internal.Vehicle{}, NewFieldValidationError("max speed must be greater than zero")
	}

	return s.rp.UpdateMaxSpeed(id, maxSpeed)
}

func (s *VehicleDefault) AddBatch(vehicles []internal.Vehicle) (map[int]internal.Vehicle, error) {
	for _, vehicle := range vehicles {
		_, err := s.rp.GetById(vehicle.Id)
		if err == nil {
			return map[int]internal.Vehicle{}, repository.NewVehicleAlreadyExistsError(vehicle.Id)
		}

		if err := vehicle.IsValid(); err != nil {
			return map[int]internal.Vehicle{}, err
		}
	}

	newVehicles := make(map[int]internal.Vehicle)
	for _, vehicle := range vehicles {
		addedVehicle, err := s.rp.Add(vehicle)
		if err != nil {
			return map[int]internal.Vehicle{}, err
		}
		newVehicles[addedVehicle.Id] = addedVehicle
	}

	return newVehicles, nil

}
