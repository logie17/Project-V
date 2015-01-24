package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
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

func isAuthenticated(c *gin.Context) {
	session, err := store.Get(c.Request, "flash-session")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
	var username = session.Values["username"]
	if username == nil {
		session.AddFlash("Login failed!", "message")
		session.Save(c.Request, c.Writer)
		c.Fail(http.StatusUnauthorized, errors.New("Unauthorized")) // idk why this is needed but it is
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}
