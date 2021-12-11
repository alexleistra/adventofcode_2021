package entities

type Squid struct {
	neighbours []*Squid
	energy     int
}

func NewSquid(energy int) Squid {
	return Squid{energy: energy}
}

func (s *Squid) AddNeighbour(squid *Squid) {
	s.neighbours = append(s.neighbours, squid)
}

func (s *Squid) IncreaseEnergy() {
	s.energy++
	if s.energy == 10 {
		for _, n := range s.neighbours {
			n.IncreaseEnergy()
		}
	}
}

func (s *Squid) FlashIfCanFlash() int {
	if s.energy > 9 {
		s.energy = 0
		flashes := 1
		for _, n := range s.neighbours {
			flashes += n.FlashIfCanFlash()
		}
		return flashes
	}
	return 0
}
