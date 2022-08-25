package repository

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/morbalint/kem-api-go/dtos"
)

func SeedCharacters() error {
	if err := helmar.ID.Set("a556fa85-19a2-408f-b248-45c04f682041"); err != nil {
		return err
	}
	if err := ayla.ID.Set("10e49375-5955-4318-9510-fb867489ea22"); err != nil {
		return err
	}
	characters = append(characters, helmar, ayla)
	return nil
}

func GetAllCharacters() []dtos.Character {
	return characters
}

func AddCharacter(character dtos.Character) dtos.Character {
	if character.ID.Status == pgtype.Null {
		var myid = uuid.New()
		_ = character.ID.Set(myid.String())
	}
	characters = append(characters, character)
	return character
}

var helmar = dtos.Character{
	ID:             pgtype.UUID{},
	Name:           "Helmar",
	Class:          "Íjász - 9, Tolvaj - 2",
	Level:          11,
	Race:           "Elf",
	Alignment:      "Kaotikus jó",
	Gender:         "Férfi",
	AgeInYears:     34,
	Religion:       "Istár",
	Strength:       15,
	Dexterity:      16,
	Constitution:   13,
	Intelligence:   13,
	Wisdom:         11,
	Charisma:       10,
	Endurance:      6,
	ReflexBase:     6,
	WillPowerBase:  3,
	HPMax:          96,
	HPCurrent:      96,
	ACArmor:        6,
	ACShield:       2,
	ACOther:        1,
	Initiative:     6,
	MovementInFeet: 20,
}

var ayla = dtos.Character{
	Name:           "Ayla",
	Class:          "Papnő",
	Level:          11,
	Race:           "Ember",
	Alignment:      "Semleges jó",
	Gender:         "Nő",
	AgeInYears:     29,
	Religion:       "Istár",
	Strength:       12,
	Dexterity:      12,
	Constitution:   15,
	Intelligence:   15,
	Wisdom:         17,
	Charisma:       16,
	Endurance:      7,
	ReflexBase:     3,
	WillPowerBase:  7,
	HPMax:          80,
	HPCurrent:      80,
	ACArmor:        11,
	ACShield:       5,
	ACOther:        1,
	Initiative:     1,
	MovementInFeet: 20,
}

var characters = []dtos.Character{}
