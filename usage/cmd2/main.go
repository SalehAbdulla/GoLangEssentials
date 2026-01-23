package main

import (
	"context"
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
func processTruck(ctx context.Context, truck Truck) error {
	fmt.Printf("processing truck %+v\n", truck)

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	// simulate same processing time
	time.Sleep(time.Second)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	return nil
}

func processFleet(ctx context.Context, trucks []Truck) error {
	
	var wg sync.WaitGroup
	wg.Add(len(trucks))

	errChan := make(chan error)

	for _, t := range trucks {

		go func(t Truck) {
			if err := processTruck(ctx, t); err != nil {log.Fatal(err.Error())}
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

	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", 442)

	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(ctx, fleet); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Done")

}
