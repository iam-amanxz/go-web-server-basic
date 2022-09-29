package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Animal struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type AddAnimalRequestBody struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func main() {
	router := gin.Default()

	animals := []Animal{{Id: "1", Name: "Dog"}, {Id: "2", Name: "Cat"}}

	router.GET("/animals", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"payload": animals,
		})
	})

	router.POST("/animals", func(c *gin.Context) {
		body := AddAnimalRequestBody{}

		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		animals = append(animals, Animal(body))
		c.JSON(http.StatusCreated, gin.H{
			"payload": animals,
		})
	})

	router.GET("/animals/:id", func(c *gin.Context) {
		id := c.Param("id")

		var animal Animal

		for _, v := range animals {
			if v.Id == id {
				animal = v
			}
		}

		c.JSON(http.StatusCreated, gin.H{
			"payload": animal,
		})
	})

	router.DELETE("/animals/:id", func(c *gin.Context) {
		id := c.Param("id")

		for i, v := range animals {
			if v.Id == id {
				animals[i] = animals[len(animals)-1]
				animals[len(animals)-1] = Animal{}
				animals = animals[:len(animals)-1]
			}
		}

		c.JSON(http.StatusCreated, gin.H{
			"payload": animals,
		})
	})

	router.Run()
}
