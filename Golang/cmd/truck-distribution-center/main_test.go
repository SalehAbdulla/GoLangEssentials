package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	t.Run("processTruck", func(t *testing.T) {
		nt := &NormalTruck{id: "2", cargo: 0}
		err := processTruck(nt)
		if err != nil {
			t.Fatalf("Error processing truck: %s", err)
		}

		// asserting
		if nt.cargo != 1 {
			t.Error("normal truck cargo must be 1")
		}

	})

}

