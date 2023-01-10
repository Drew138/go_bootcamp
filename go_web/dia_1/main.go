package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.POST("/saludo", func(c *gin.Context) {
		var persona Persona
		bodyAsByteArray, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		json.Unmarshal(bodyAsByteArray, &persona)

		response := fmt.Sprintf("Hola %s %s", persona.Nombre, persona.Apellido)
		c.String(http.StatusOK, response)
	})
	r.Run()
}
