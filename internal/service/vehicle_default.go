package service

import "app/internal"

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
	allVehicles, err := s.rp.FindAll()
	if err != nil {
		return map[int]internal.Vehicle{}, err
	}

	vehicles := make(map[int]internal.Vehicle)
	for key, value := range allVehicles {
		if value.Brand == brand &&
			value.FabricationYear >= start &&
			value.FabricationYear <= end {
			vehicles[key] = value
		}
	}

	return vehicles, nil
}
