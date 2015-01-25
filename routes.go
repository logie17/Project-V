package main

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	//"github.com/gorilla/context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

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
