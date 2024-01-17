package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	GetById(id int) (Vehicle, error)
	FindByBrandAndYears(brand string, start, end int) (map[int]Vehicle, error)
	UpdateMaxSpeed(id int, maxSpeed float64) (Vehicle, error)
	UpdateFuel(id int, fuel string) (Vehicle, error)
	Add(v Vehicle) (Vehicle, error)
}
