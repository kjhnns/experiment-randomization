package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	"time"
)

var g *gin.Engine

func main() {

	fmt.Println("bootstrapping")
	// Init Gin-Gonic
	g = gin.Default()

	fmt.Println("Routing")

	g.GET("/", Home)

	port := "3000"
	//if PORT passed as env-variable, use the port
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	// Listen and serve on 0.0.0.0:3000
	g.Run(":" + port)
}

func Home(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	mode := rand.Intn(4)

	urls := []string{
		"https://juntoapp.github.io/0.html",
		"https://juntoapp.github.io/p.html",
		"https://juntoapp.github.io/pc.html",
		"https://juntoapp.github.io/c.html",
	}

	c.Redirect(302, urls[mode])
}
