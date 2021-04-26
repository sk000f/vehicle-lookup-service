package ports

import "github.com/sk000f/vinlookup/core/domain"

type Service interface {
	AddVehicle(v domain.Vehicle) error
	LookupVehicle(VIN int) (domain.Vehicle, error)
	PaintVehicle(VIN int, c string) (domain.Vehicle, error)
}
