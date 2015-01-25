package handles

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/sessions"
	"net/http"
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
