package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	FindByBrandAndYears(brand string, start, end int) (map[int]Vehicle, error)
	Add(v Vehicle) (Vehicle, error)
	AddBatch(v []Vehicle) (map[int]Vehicle, error)
	UpdateMaxSpeed(id int, maxSpeed float64) (Vehicle, error)
}
