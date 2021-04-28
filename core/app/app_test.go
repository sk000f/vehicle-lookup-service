package app_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/sk000f/vinlookup/core/app"
	"github.com/sk000f/vinlookup/core/domain"
)

func TestAddVehicle(t *testing.T) {
	r := createMockRepository()

	a := app.New(r)

	v := domain.Vehicle{
		VIN:    12345,
		Make:   "Land Rover",
		Model:  "Discovery",
		Colour: "Silver",
	}

	a.AddVehicle(v)

	want := []domain.Vehicle{
		{
			VIN:    12345,
			Make:   "Land Rover",
			Model:  "Discovery",
			Colour: "Silver",
		},
	}

	got := r.Data

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestLookupVehicle(t *testing.T) {
	r := createMockRepository()

	v := domain.Vehicle{
		VIN:    12345,
		Make:   "Land Rover",
		Model:  "Discovery",
		Colour: "Silver",
	}

	r.Data = append(r.Data, v)

	a := app.New(r)

	want := v

	got, err := a.LookupVehicle(12345)
	if err != nil {
		t.Errorf("test failed with error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}

}

func TestPaintVehicle(t *testing.T) {
	r := createMockRepository()

	v := domain.Vehicle{
		VIN:    12345,
		Make:   "Land Rover",
		Model:  "Discovery",
		Colour: "Silver",
	}

	r.Data = append(r.Data, v)

	a := app.New(r)

	v.Colour = "Black"

	want := v

	got, err := a.PaintVehicle(12345, "Black")
	if err != nil {
		t.Errorf("test failed with error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}

}

type mockRepository struct {
	Data []domain.Vehicle
}

func (m *mockRepository) Save(v domain.Vehicle) error {
	m.Data = append(m.Data, v)
	return nil
}

func (m *mockRepository) Lookup(VIN int) (domain.Vehicle, error) {
	for _, v := range m.Data {
		if v.VIN == VIN {
			return v, nil
		}
	}
	return domain.Vehicle{}, errors.New("VIN not found")
}

func createMockRepository() *mockRepository {
	return &mockRepository{}
}
