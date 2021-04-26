package mysql

import "github.com/sk000f/vinlookup/core/domain"

func Save(v domain.Vehicle) error {
	return nil
}

func Lookup(VIN string) (domain.Vehicle, error) {
	return domain.Vehicle{}, nil
}
