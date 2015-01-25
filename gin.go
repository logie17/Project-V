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
	router.Use(IsMobile()) // this doesnt work yet
	// this is how we can get global template data
	//set := pongo2.NewSet("our web templates") // The idea behind sets is that you can create another set with other globals and configurations for mail templates or other kind of templates
	//set.Globals["global_variable"] = "this is a test"
	// https://github.com/flosch/pongo2/issues/35
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
