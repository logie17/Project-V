package handles

import (
	"github.com/gin-gonic/gin"
	"github.com/flosch/pongo2"
	"github.com/gorilla/sessions"
	"net/http"
	"github.com/gin-gonic/gin/binding"
)

type SignupForm struct {
	FirstName string `form:"f_name" binding:"required"`
	Password  string `form:"password" binding:"required"`
}

func SignupGetHandler(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := pongo2.Context{
			"title": "Pair",
		}
		c.HTML(http.StatusOK, "templates/pages/signup.html", ctx)
	}
}

func SignupPostHandler(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		// lets checkout the form
		var form SignupForm
		c.BindWith(&form, binding.Form)
		// logie should do something about this
	}
}
