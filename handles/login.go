package handles

import (
	"errors"
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/sessions"
	"net/http"
)

type LoginForm struct {
	User     string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func LoginGetHandler(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func LoginPostHandler(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}