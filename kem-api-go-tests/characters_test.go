package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestGetAllCharactersReturnsTwoSeededCharacters(t *testing.T) {
	// Arrange
	baseUrl := "https://localhost:8080/"

	// Act
	response, err := http.Get(baseUrl + "characters")
	if err != nil {
		t.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	var characters []Character
	err = json.Unmarshal(responseData, &characters)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	if len(characters) != 2 {
		t.Fatalf(`GET /characters got %#v, want 2 length slice`, characters)
	}
}
