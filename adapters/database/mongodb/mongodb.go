package mongodb

import "github.com/sk000f/vinlookup/core/domain"

func Save(v domain.Vehicle) error {
	return nil
}

func Lookup(VIN int) (domain.Vehicle, error) {
	return domain.Vehicle{}, nil
}
