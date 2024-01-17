package repository

import "fmt"

type VehicleAlreadyExistsError struct {
	id int
}

func (e *VehicleAlreadyExistsError) Error() string {
	return fmt.Sprintf("vehicle with id %d already exists", e.id)
}

func NewVehicleAlreadyExistsError(id int) error {
	return &VehicleAlreadyExistsError{id: id}
}

type QueryError struct {
	msg string
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("query error: %s", e.msg)
}

func NewQueryError(msg string) error {
	return &QueryError{msg: msg}
}

type VehicleNotFoundError struct {
	id int
}

func (e *VehicleNotFoundError) Error() string {
	return fmt.Sprintf("vehicle with id %d not found", e.id)
}

func NewVehicleNotFoundError(id int) error {
	return &VehicleNotFoundError{id: id}
}
