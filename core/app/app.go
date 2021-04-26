package app

import (
	"github.com/sk000f/vinlookup/core/domain"
	"github.com/sk000f/vinlookup/core/ports"
)

type App struct {
	r ports.Repository
}

func New(r ports.Repository) *App {
	return &App{
		r: r,
	}
}

func (a *App) AddVehicle(v domain.Vehicle) error {
	err := a.r.Save(v)

	if err != nil {
		return err
	}

	return nil
}

func (a *App) LookupVehicle(VIN int) (domain.Vehicle, error) {
	return a.r.Lookup(VIN)
}

func (a *App) PaintVehicle(VIN int, c string) (domain.Vehicle, error) {
	v, err := a.r.Lookup(VIN)
	if err != nil {
		return domain.Vehicle{}, err
	}

	v2 := &v
	v2.Colour = c

	err = a.r.Save(*v2)
	if err != nil {
		return domain.Vehicle{}, err
	}
	return *v2, nil
}
