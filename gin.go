package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/sessions"
	"net/http"
	"time"
)

func main() {
	router := gin.New()
	router.Use(Logrus())
	router.LoadHTMLGlob("./*")
	router.GET("/", func(c *gin.Context) {
		obj := gin.H{"title": "Main website"}
		c.HTML(200, "index.tmpl", obj)
	})
	router.POST("/login", set)
	router.GET("/get", get)
	router.Run(":8080")
}
func Logrus() gin.HandlerFunc {
	var log = logrus.New()
	log.Level = logrus.InfoLevel
	log.Formatter = &logrus.TextFormatter{}
	return func(c *gin.Context) {
		t := time.Now()
		log.WithFields(logrus.Fields{
			"method":  c.Request.Method,
			"request": c.Request.URL,
			"remote":  c.Request.RemoteAddr,
		}).Info("started handling request")
		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.WithFields(logrus.Fields{
			"status": c.Writer.Status(),
			"proto":  c.Request.Proto,
			"took":   latency,
		}).Info("completed handling request")
	}
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
