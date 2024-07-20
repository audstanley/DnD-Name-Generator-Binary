package species

import (
	"fmt"
	"math/rand"
	"time"
)

// Generic interface for species
type Species interface {
	Aasimar | Dragonborn | Dwarf | Elf | Gnome |
		HalfElf | HalfOrc | Halfling | Human | Tiefling
}

type GenericSpecies struct {
	GlobalVariableName string   `mapstructure:"global_variable_name"` // Name of the global variable
	FirstFNames        []string `mapstructure:"first_f_names"`
	FirstMNames        []string `mapstructure:"first_m_names"`
	FirstNBNames       []string `mapstructure:"first_nb_names"`
	FirstNames         []string `mapstructure:"first_names"`
	LastNames          []string `mapstructure:"last_names"`
	SpecialNames       []string `mapstructure:"special_names"`
	LocationNames      []string `mapstructure:"location_names"`
}

func (species GenericSpecies) GetRandomFirstName() error {
	if len(species.FirstNames) == 0 {
		return fmt.Errorf("FirstNames for Aasimar List is empty")
	}
	if len(species.LastNames) == 0 {
		return fmt.Errorf("LastNames for Aasimar List is empty")
	}
	// Seed the random number generator
	// Use a new source for randomness (recommended)
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Pick two random indices
	index1 := random.Intn(len(species.FirstNames))
	index2 := random.Intn(len(species.LastNames))

	// Ensure indices are unique (avoid duplicates)
	for index1 == index2 {
		index2 = random.Intn(len(species.LastNames))
	}
	fmt.Printf("%s %s\n", species.FirstNames[index1], species.LastNames[index2])
	return nil
}

// Speices types for getting First and Last names from a file.
type Aasimar struct{ GenericSpecies }
type Dragonborn struct{ GenericSpecies }
type Dwarf struct{ GenericSpecies }
type Elf struct{ GenericSpecies }
type Gnome struct{ GenericSpecies }
type HalfElf struct{ GenericSpecies }
type HalfOrc struct{ GenericSpecies }
type Halfling struct{ GenericSpecies }
type Human struct{ GenericSpecies }
type Tiefling struct{ GenericSpecies }

var PublicAasimar = Aasimar{}
var PublicDragonborn = Dragonborn{}
var PublicDwarf = Dwarf{}
var PublicElf = Elf{}
var PublicGnome = Gnome{}
var PublicHalfElf = HalfElf{}
var PublicHalfOrc = HalfOrc{}
var PublicHalfling = Halfling{}
var PublicHuman = Human{}
var PublicTiefling = Tiefling{}

var PublicLocationNames []string

func (species GenericSpecies) Init() {
	// Populate LocationNames from YAML file

	PublicAasimar.GlobalVariableName = "Aasimar"
	PublicAasimar.FirstFNames = []string{"aasimar_first_names.txt"}
	PublicAasimar.FirstMNames = []string{"aasimar_first_names.txt"}
	PublicAasimar.FirstNBNames = []string{"aasimar_first_names.txt"}
	PublicAasimar.LastNames = []string{"aasimar_last_names.txt"}
	PublicAasimar.SpecialNames = []string{"aasimar_special_names.txt"}

}

// generic species function for getting names
func GetFirstNames[T interface{}](s T) T {
	//var inc T
	return s
}
