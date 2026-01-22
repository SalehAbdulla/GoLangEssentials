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
	
	// How to use the type of interface{} or any
	person := make(map[string]interface{})
	person["name"] = "SALEH"
	person["age"] = 42

	age, exists := person["age"].(int) // Specifying the type to get the result
	if !exists {
		log.Fatal("age not exists")
	}
	println(age)


}
