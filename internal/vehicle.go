package internal

import "fmt"

// Dimensions is a struct that represents a dimension in 3d
type Dimensions struct {
	// Height is the height of the dimension
	Height float64
	// Length is the length of the dimension
	Length float64
	// Width is the width of the dimension
	Width float64
}

// VehicleAttributes is a struct that represents the attributes of a vehicle
type VehicleAttributes struct {
	// Brand is the brand of the vehicle
	Brand string
	// Model is the model of the vehicle
	Model string
	// Registration is the registration of the vehicle
	Registration string
	// Color is the color of the vehicle
	Color string
	// FabricationYear is the fabrication year of the vehicle
	FabricationYear int
	// Capacity is the capacity of people of the vehicle
	Capacity int
	// MaxSpeed is the maximum speed of the vehicle
	MaxSpeed float64
	// FuelType is the fuel type of the vehicle
	FuelType string
	// Transmission is the transmission of the vehicle
	Transmission string
	// Weight is the weight of the vehicle
	Weight float64
	// Dimensions is the dimensions of the vehicle
	Dimensions
}

// Vehicle is a struct that represents a vehicle
type Vehicle struct {
	// Id is the unique identifier of the vehicle
	Id int

	// VehicleAttribue is the attributes of a vehicle
	VehicleAttributes
}

type VehicleValidationError struct {
	field string
}

func (v *VehicleValidationError) Error() string {
	return fmt.Sprintf("field: %s can not be empty", v.field)
}

func NewVehicleValidationError(field string) error {
	return &VehicleValidationError{field: field}
}

func (v Vehicle) IsValid() error {
	switch {
	case v.Id <= 0:
		return NewVehicleValidationError("id")
	case v.Brand == "":
		return NewVehicleValidationError("brand")
	case v.Model == "":
		return NewVehicleValidationError("model")
	case v.Registration == "":
		return NewVehicleValidationError("registration")
	case v.Color == "":
		return NewVehicleValidationError("color")
	case v.FabricationYear <= 0:
		return NewVehicleValidationError("fabrication_year")
	case v.Capacity <= 0:
		return NewVehicleValidationError("capacity")
	case v.MaxSpeed <= 0:
		return NewVehicleValidationError("max_speed")
	case v.FuelType == "":
		return NewVehicleValidationError("fuel_type")
	case v.Transmission == "":
		return NewVehicleValidationError("transmission")
	case v.Weight <= 0:
		return NewVehicleValidationError("weight")
	case v.Height <= 0:
		return NewVehicleValidationError("height")
	case v.Length <= 0:
		return NewVehicleValidationError("length")
	case v.Width <= 0:
		return NewVehicleValidationError("width")
	}

	return nil
}
