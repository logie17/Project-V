package main

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	//"github.com/gorilla/context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/sessions"
	"github.com/logie17/Project-V/config"
	m "github.com/logie17/Project-V/middleware"
	"net/http"
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

type LoginForm struct {
	User     string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type SignupForm struct {
	FirstName string `form:"f_name" binding:"required"`
	Password  string `form:"password" binding:"required"`
}

func indexHandler(c *gin.Context) {
	ctx := pongo2.Context{
		"title": "pairparty.io",
	}
	c.HTML(http.StatusOK, "templates/pages/index.html", ctx)
	/*
		ajax := r.Header.Get("X-PUSH")
		log.Println(context.Get(r, UserKey))
		if ajax != "" {
			// serve the partial
			err := index_partial_tmpl.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			// serve a page
			err := index_page_tmpl.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	*/
}

func loginGetHandler(c *gin.Context) {
	ctx := pongo2.Context{
		"title": "Login",
	}
	session, err := store.Get(c.Request, "flash-session")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
	// already logged in
	var username = session.Values["username"]
	if username != nil {
		c.Fail(http.StatusUnauthorized, errors.New("Unauthorized")) // idk why this is needed but it is
		c.Redirect(http.StatusMovedPermanently, "/pair")
		return
	}
	if flashes := session.Flashes(); len(flashes) > 0 {
		fmt.Fprint(c.Writer, "%v", flashes)
	}
	c.HTML(http.StatusOK, "templates/pages/login.html", ctx)
}

func loginPostHandler(c *gin.Context) {
	session, err := store.Get(c.Request, "flash-session")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
	// lets checkout the form
	var form LoginForm
	c.BindWith(&form, binding.Form)
	if form.User == "km@km.com" {
		session.Values["username"] = form.User
		c.Redirect(http.StatusMovedPermanently, "/pair")
	}
	session.AddFlash("Login failed!", "message")
	session.Save(c.Request, c.Writer)
}

func signupGetHandler(c *gin.Context) {
	ctx := pongo2.Context{
		"title": "Pair",
	}
	c.HTML(http.StatusOK, "templates/pages/signup.html", ctx)
}
func signupPostHandler(c *gin.Context) {
	// lets checkout the form
	var form SignupForm
	c.BindWith(&form, binding.Form)
	// logie should do something about this
}

func pairGetHandler(c *gin.Context) {
	ctx := pongo2.Context{
		"title": "Pair",
	}
	c.HTML(http.StatusOK, "templates/pages/pair.html", ctx)
}
