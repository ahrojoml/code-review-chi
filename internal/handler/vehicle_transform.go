package handler

import "app/internal"

func vehicleToVehicleJSON(v internal.Vehicle) VehicleJSON {
	vehicleJSON := VehicleJSON{
		ID:              v.Id,
		Brand:           v.Brand,
		Model:           v.Model,
		Registration:    v.Registration,
		Color:           v.Color,
		FabricationYear: v.FabricationYear,
		Capacity:        v.Capacity,
		MaxSpeed:        v.MaxSpeed,
		FuelType:        v.FuelType,
		Transmission:    v.Transmission,
		Weight:          v.Weight,
		Height:          v.Dimensions.Height,
		Length:          v.Dimensions.Length,
		Width:           v.Dimensions.Width,
	}
	return vehicleJSON
}

func vehicleJSONToVehicle(vj VehicleJSON) internal.Vehicle {
	vehicleDimensions := internal.Dimensions{
		Height: vj.Height,
		Length: vj.Length,
		Width:  vj.Width,
	}

	vehicleAttr := internal.VehicleAttributes{
		Brand:           vj.Brand,
		Model:           vj.Model,
		Registration:    vj.Registration,
		Color:           vj.Color,
		FabricationYear: vj.FabricationYear,
		Capacity:        vj.Capacity,
		MaxSpeed:        vj.MaxSpeed,
		FuelType:        vj.FuelType,
		Transmission:    vj.Transmission,
		Weight:          vj.Weight,
		Dimensions:      vehicleDimensions,
	}
	// ...
	newVehicle := internal.Vehicle{
		Id:                vj.ID,
		VehicleAttributes: vehicleAttr,
	}

	return newVehicle
}
