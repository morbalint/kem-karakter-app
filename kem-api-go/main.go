package main

import (
	"github.com/gin-gonic/gin"
	"github.com/morbalint/kem-api-go/dtos"
	"github.com/morbalint/kem-api-go/repository"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/characters", getCharacters)
	router.GET("/characters/:name", getCharacterByName)
	router.POST("/characters", postCharacters)

	router.Run("localhost:8080")
}

func getCharacters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, repository.Characters)
}

func postCharacters(c *gin.Context) {
	var newCharacter dtos.Character

	// Call BindJSON to bind the received JSON to
	// newCharacter.
	if err := c.BindJSON(&newCharacter); err != nil {
		return
	}

	// Add the new album to the slice.
	repository.Characters = append(repository.Characters, newCharacter)
	c.IndentedJSON(http.StatusCreated, newCharacter)
}

func getCharacterByName(c *gin.Context) {
	name := c.Param("name")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range repository.Characters {
		if a.Name == name {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Character not found"})
}
