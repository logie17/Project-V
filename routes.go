package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/logie17/Project-V/config"
	m "github.com/logie17/Project-V/middleware"
	h "github.com/logie17/Project-V/handles"
)

var store *sessions.CookieStore = sessions.NewCookieStore([]byte("a-secret-string"))


func main() {
	router := gin.New()
	configuration := config.LoadConfig()
	router.Use(m.Logrus())
	router.Use(m.IsMobile()) // this doesnt work yet
	// this is how we can get global template data
	//set := pongo2.NewSet("our web templates") // The idea behind sets is that you can create another set with other globals and configurations for mail templates or other kind of templates
	//set.Globals["global_variable"] = "this is a test"
	// https://github.com/flosch/pongo2/issues/35
	router.HTMLRender = newPongoRender()
	router.Static("/public", "./public")

	router.GET("/", h.IndexHandler)

	router.GET("/login", h.LoginGetHandler(store))
	router.POST("/login", h.LoginPostHandler(store))

	router.GET("/signup", h.SignupGetHandler(store))
	router.POST("/signup", h.LoginPostHandler(store))

	router.GET("/pair", m.IsAuthenticated(store), h.PairGetHandler)

	router.Run(fmt.Sprintf("[::]:%s", configuration.Port))
}




