package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgtype"

	"github.com/morbalint/kem-api-go/dtos"
	"github.com/morbalint/kem-api-go/repository"
)

func GetCharacters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, repository.GetAllCharacters())
}

func PostCharacters(c *gin.Context) {
	var newCharacter dtos.Character

	// Call BindJSON to bind the received JSON to
	// newCharacter.
	if err := c.ShouldBind(&newCharacter); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		_ = c.BindJSON(&newCharacter)

		// Add the new album to the slice.
		var dto = repository.AddCharacter(newCharacter)
		c.IndentedJSON(http.StatusCreated, dto)
	}
}

func GetCharacterByName(c *gin.Context) {
	name := c.Param("name")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range repository.GetAllCharacters() {
		if a.Name == name {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Character not found"})
}

func GetCharacterById(c *gin.Context) {
	idStr := c.Param("id")
	var uuid = pgtype.UUID{}
	err := uuid.Set(idStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range repository.GetAllCharacters() {
		if a.ID == uuid {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Character not found"})
}
