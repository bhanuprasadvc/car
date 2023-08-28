package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	ID          int    `json:"id"`
	Model       string `json:"model"`
	Make        string `json:"make"`
	Year        int    `json:"year"`
	Availability bool   `json:"availability"`
}

var cars = []Car{
	{ID: 1, Model: "Civic", Make: "Honda", Year: 2020, Availability: true},
	{ID: 2, Model: "Corolla", Make: "Toyota", Year: 2021, Availability: false},
	// Add more cars here
}

func listAvailableCars(c *gin.Context) {
	var availableCars []Car
	for _, car := range cars {
		if car.Availability {
			availableCars = append(availableCars, car)
		}
	}
	c.JSON(http.StatusOK, availableCars)
}

func main() {
	router := gin.Default()

	router.GET("/cars", listAvailableCars)

	router.Run(":8080")
}
