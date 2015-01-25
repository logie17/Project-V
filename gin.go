package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/logie17/Project-V/config"
	m "github.com/logie17/Project-V/middleware"
)

func main() {
	router := gin.New()
	configuration := config.LoadConfig()
	router.Use(Logrus())
	router.HTMLRender = newPongoRender()
	router.Static("/public", "./public")

	router.GET("/", indexHandler)

	router.GET("/login", loginGetHandler)
	router.POST("/login", loginPostHandler)

	router.GET("/signup", signupGetHandler)
	router.POST("/signup", loginPostHandler)

	router.GET("/pair", m.IsAuthenticated(store), pairGetHandler)

	router.Run(fmt.Sprintf("[::]:%s", configuration.Port))
}

var store *sessions.CookieStore = sessions.NewCookieStore([]byte("a-secret-string"))
