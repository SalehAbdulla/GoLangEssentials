package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
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

	// simulate same processing time
	time.Sleep(time.Second)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	return nil
}


func processFleet(trucks []Truck) error {
	var wg sync.WaitGroup
	wg.Add(len(trucks))

	for _, t := range trucks {

		go func(t Truck) {
			if err := processTruck(t); err != nil {log.Fatal(err.Error())}
			wg.Done()
		}(t)
			
	}

	wg.Wait()
	return nil
}

func fillTruckCargo(t *NormalTruck) {
	t.cargo = 100
}

func main() {

	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(fleet); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Done")

}
