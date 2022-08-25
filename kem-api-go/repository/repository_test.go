package repository

import (
	"github.com/jackc/pgtype"
	"github.com/morbalint/kem-api-go/dtos"
	"golang.org/x/exp/slices"
	"testing"
)

// Given seeded characters
// When I call GetAllCharacters
// Then 2 characters are returned
func TestGetAllCharactersReturnsTwoSeededCharacters(t *testing.T) {
	// Arrange
	if err := SeedCharacters(); err != nil {
		t.Fatal(err)
	}

	// Act
	allCharacters := GetAllCharacters()

	// Assert
	if len(allCharacters) != 2 {
		t.Fatalf(`GetAllCharacters() = %#v, want 2 length slice`, allCharacters)
	}
}

// Given seeded characters
// When I call GetAllCharacters twice
// Then the returned characters are the same
func TestGetAllCharactersCalledTwiceReturnsSameCharacters(t *testing.T) {
	// Arrange
	if err := SeedCharacters(); err != nil {
		t.Fatal(err)
	}

	// Act
	firstBatch := GetAllCharacters()
	secondBatch := GetAllCharacters()

	// Assert
	if len(firstBatch) != len(secondBatch) || firstBatch[0] != secondBatch[0] || firstBatch[1] != secondBatch[1] {
		t.Fatalf(`GetAllCharacters() = %#v, want the same as GetAllCharacters() = %#v`, firstBatch, secondBatch)
	}
}

// Given a valid character with an ID
// When I call AddCharacter
// Then the same character is returned
func TestAddCharacterGivenAValidCharacterReturnsSameCharacter(t *testing.T) {
	// Arrange
	validCharacter := dtos.Character{
		ID:             pgtype.UUID{},
		Name:           "Test",
		Class:          "NPC",
		Level:          1,
		Race:           "Ember",
		Alignment:      "Semleges",
		Gender:         "",
		AgeInYears:     20,
		Religion:       "",
		Strength:       10,
		Dexterity:      10,
		Constitution:   10,
		Intelligence:   10,
		Wisdom:         10,
		Charisma:       10,
		Endurance:      0,
		ReflexBase:     0,
		WillPowerBase:  0,
		HPMax:          10,
		HPCurrent:      10,
		ACArmor:        10,
		ACShield:       0,
		ACOther:        0,
		Initiative:     0,
		MovementInFeet: 30,
	}
	if err := validCharacter.ID.Set("d49c951e-f288-4963-ac40-4170bf66552a"); err != nil {
		t.Fatal(err)
	}

	// Act
	actualCharacter := AddCharacter(validCharacter)

	// Assert
	if actualCharacter != validCharacter || actualCharacter.ID != validCharacter.ID {
		t.Fatalf(`AddCharacter(%#v) = %#v, want the returned character to be the same`, validCharacter, actualCharacter)
	}
}

// Given a valid character with an ID
// And I call AddCharacter
// When I call GetAllCharacters
// Then the newly added character is also returned
func TestAddCharacterThenGetAllCharactersReturnTheNewCharacter(t *testing.T) {
	// Arrange
	validCharacter := dtos.Character{
		ID:             pgtype.UUID{},
		Name:           "Test",
		Class:          "NPC",
		Level:          1,
		Race:           "Ember",
		Alignment:      "Semleges",
		Gender:         "",
		AgeInYears:     20,
		Religion:       "",
		Strength:       10,
		Dexterity:      10,
		Constitution:   10,
		Intelligence:   10,
		Wisdom:         10,
		Charisma:       10,
		Endurance:      0,
		ReflexBase:     0,
		WillPowerBase:  0,
		HPMax:          10,
		HPCurrent:      10,
		ACArmor:        10,
		ACShield:       0,
		ACOther:        0,
		Initiative:     0,
		MovementInFeet: 30,
	}
	if err := validCharacter.ID.Set("d49c951e-f288-4963-ac40-4170bf66552a"); err != nil {
		t.Fatal(err)
	}
	AddCharacter(validCharacter)

	// Act
	allCharacters := GetAllCharacters()

	// Assert
	if !slices.Contains(allCharacters, validCharacter) {
		t.Fatalf(`GetAllCharacters() = %#v, want %#v`, allCharacters, validCharacter)
	}
}

func TestGetAllCharactersGivenCharacterWithoutIdGeneratesAnId(t *testing.T) {
	// Arrange
	characterWithoutId := dtos.Character{
		ID:             pgtype.UUID{},
		Name:           "Test",
		Class:          "NPC",
		Level:          1,
		Race:           "Ember",
		Alignment:      "Semleges",
		Gender:         "",
		AgeInYears:     20,
		Religion:       "",
		Strength:       10,
		Dexterity:      10,
		Constitution:   10,
		Intelligence:   10,
		Wisdom:         10,
		Charisma:       10,
		Endurance:      0,
		ReflexBase:     0,
		WillPowerBase:  0,
		HPMax:          10,
		HPCurrent:      10,
		ACArmor:        10,
		ACShield:       0,
		ACOther:        0,
		Initiative:     0,
		MovementInFeet: 30,
	}

	// Act
	addedCharacter := AddCharacter(characterWithoutId)

	// Assert
	if addedCharacter.ID.Status == pgtype.Null {
		t.Fatalf(`AddCharacters(%#v) = %#v, want ID to be not Null`, characterWithoutId, addedCharacter)
	}
}
