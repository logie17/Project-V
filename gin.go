package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func main() {
	router := gin.New()
	router.Use(Logrus())
	router.HTMLRender = newPongoRender()
	router.Static("/public", "./public")

	router.GET("/", indexHandler)
	router.GET("/login", loginGetHandler)
	router.POST("/login", loginPostHandler)
	router.GET("/pair", isAuthenticated, pairGetHandler)
	router.Run(":3001")
}

var store = sessions.NewCookieStore([]byte("a-secret-string"))

type LoginForm struct {
	User     string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
