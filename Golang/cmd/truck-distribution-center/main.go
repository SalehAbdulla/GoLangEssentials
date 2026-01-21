package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
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
	t.cargo += 1
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery = -1
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo += 1
	e.battery = -1
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

	nT := &NormalTruck{id:"1"}
	if err := processTruck(nT); err != nil {
		log.Fatalf("error processing truck: %s", err)
	}

	eT := &ElectricTruck{id:"2"}
	if err := processTruck(eT); err != nil {
		log.Fatalf("error processing truck: %s", err)
	}

	log.Println(nT.cargo)
	log.Println(eT.battery)

}
