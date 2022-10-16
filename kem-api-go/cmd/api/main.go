package main

import (
	"github.com/gin-gonic/gin"

	"github.com/morbalint/kem-api-go/handlers"
	"github.com/morbalint/kem-api-go/repository"
)

func main() {

	_ = repository.SeedCharacters()

	router := gin.Default()

	router.GET("/characters", handlers.GetCharacters)
	router.GET("/characters?name=:name", handlers.GetCharacterByName)
	router.GET("/characters/:id", handlers.GetCharacterById)
	router.POST("/characters", handlers.PostCharacters)

	_ = router.SetTrustedProxies(nil)
	_ = router.Run(":8080")
}
