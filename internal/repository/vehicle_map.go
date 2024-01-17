package repository

import (
	"app/internal"
	"fmt"
)

type VehicleAlreadyExistsError struct {
	id int
}

func (e *VehicleAlreadyExistsError) Error() string {
	return fmt.Sprintf("vehicle with id %d already exists", e.id)
}

func NewVehicleAlreadyExistsError(id int) error {
	return &VehicleAlreadyExistsError{id: id}
}

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) Add(v internal.Vehicle) (internal.Vehicle, error) {
	_, ok := r.db[v.Id]
	if ok {
		return internal.Vehicle{}, NewVehicleAlreadyExistsError(v.Id)
	}
	r.db[v.Id] = v
	return v, nil
}
