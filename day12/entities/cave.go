package entities

type CaveType string

const (
	Start CaveType = "start"
	End   CaveType = "end"
	Large CaveType = "large"
	Small CaveType = "small"
)

type Cave struct {
	name        string
	caveType    CaveType
	linkedCaves []*Cave
}

func NewCave(name string, caveType CaveType) Cave {
	return Cave{name: name, caveType: caveType}
}

func (c *Cave) AddLinkedCave(cave *Cave) {
	if cave.caveType != Start && c.caveType != End {
		c.linkedCaves = append(c.linkedCaves, cave)
	}
}

func (c *Cave) GetLinkedCaves() []*Cave {
	return c.linkedCaves
}

func (c *Cave) GetName() string {
	return c.name
}

func (c *Cave) GetType() CaveType {
	return c.caveType
}
