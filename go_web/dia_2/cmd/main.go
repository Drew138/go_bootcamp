package main

import (
	"dia_2/cmd/router"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	e := gin.Default()
	r := router.NewRouter(e)
	r.SetRoutes()
	if err := e.Run(); err != nil {
		log.Fatal(err)
	}
}
