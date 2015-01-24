package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/sessions"
	"net/http"
)

func main() {
	router := gin.New()
	router.Use(Logrus())
	router.HTMLRender = newPongoRender()
	router.Static("/public", "./public")

	router.GET("/", func(c *gin.Context) {
	})
	router.GET("/login", loginGetHandler)
	router.POST("/login", set)
	router.GET("/get", get)
	router.Run(":3001")
}

var store = sessions.NewCookieStore([]byte("a-secret-string"))

type LoginForm struct {
	User     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func set(c *gin.Context) {
	session, err := store.Get(c.Request, "flash-session")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
	var form LoginForm

	c.BindWith(&form, binding.Form) // You can also specify which binder to use. We support binding.Form, binding.JSON and binding.XML.
	fmt.Printf("%s\n", form.User)
	fmt.Printf("%s\n", form.Password)
	session.Values["username"] = form.User
	/*
		if form.user == "manu" && form.Password == "123" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	*/

	session.AddFlash("This is a flashed message!", "message")
	session.Save(c.Request, c.Writer)
}

func get(c *gin.Context) {
	session, err := store.Get(c.Request, "flash-session")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
	var username = session.Values["username"]
	if username == nil {
		c.String(http.StatusUnauthorized, "Not logged in")
	} else {
		c.String(http.StatusOK, "Logged in")
	}
	session.Save(c.Request, c.Writer)
}
