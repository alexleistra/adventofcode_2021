package entities

import "testing"

func TestIncreaseEnergy(t *testing.T) {
	squid1 := NewSquid(7)
	squid2 := NewSquid(6)
	squid3 := NewSquid(8)
	squid3.AddNeighbour(&squid1)
	squid3.AddNeighbour(&squid2)
	squid3.IncreaseEnergy()

	if squid3.energy != 9 {
		t.Errorf("IncreaseEnergy did not increase energy")
	}
}

func TestFlashIfCanFlash(t *testing.T) {
	squid1 := NewSquid(6)
	squid2 := NewSquid(9)
	squid2.AddNeighbour(&squid1)
	squid2.IncreaseEnergy()
	flashes := squid2.FlashIfCanFlash()

	if flashes != 1 {
		t.Errorf("flash did not increase flash count")
	}

	if squid2.energy != 0 {
		t.Errorf("flash did not reset energy")
	}

	if squid1.energy != 7 {
		t.Errorf("flash did not increase neighbour's energy")
	}
}
