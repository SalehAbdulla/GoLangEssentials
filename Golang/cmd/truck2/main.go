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
	id    string
	cargo int
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo++
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo++
	e.battery--
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	return nil
}

var (
	ErrNotImplemnted = errors.New("not implemented")
	ErrTruckNotFound = errors.New("truck not found")
)

// processTruck handles the loading and unloading of a truck
func processTruck(truck Truck) error {
	fmt.Printf("processing truck %+v\n", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	return nil
}

func main() {

	t := NormalTruck{cargo: 0}

	fillTruckCargo(&t)

	log.Println(t.cargo)

}


func fillTruckCargo(t *NormalTruck) {
	t.cargo = 100
}

