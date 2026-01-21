package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck interface {
	LoadCargo() error
	UnloadTruck() error
}

type NormalTruck struct {
	id string
	cargo int
}

type ElectricTruck struct {
	id string
	cargo int
	battery float64
}

func (t *NormalTruck) LoadCargo() error {
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	return nil
}

func (e *ElectricTruck) LoadCargo() error {
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	return nil
}

var (
	ErrNotImplemnted = errors.New("not implemented");
	ErrTruckNotFound = errors.New("truck not found")
)

// processTruck handles the loading and unloading of a truck
func processTruck(truck Truck) error {
	fmt.Printf("processing truck %+v\n", truck)
	
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	return ErrNotImplemnted
}

func main() {
	eTrucks := []ElectricTruck{
		{id: "Electric-truck-1"},
	}

	if err := processTruck(NormalTruck{id:"1"}); err != nil {
		log.Fatal("error processing truck: %s", err)
	}

	
	if err := processTruck(NormalTruck{id:"2"}); err != nil {
		log.Fatal("error processing truck: %s", err)
	}

}
