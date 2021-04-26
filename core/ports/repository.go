package ports

import (
	"github.com/sk000f/vinlookup/core/domain"
)

type Repository interface {
	Save(v domain.Vehicle) error
	Lookup(VIN string) (domain.Vehicle, error)
}
