package ports

import "github.com/sk000f/vinlookup/core/domain"

type Service interface {
	AddVehicle(v domain.Vehicle) error
	LookupVehicle(VIN string) (domain.Vehicle, error)
}
