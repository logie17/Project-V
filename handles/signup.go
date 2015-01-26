package handles

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/Sirupsen/logrus"
	"github.com/agnivade/easy-scrypt"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/sessions"
)

type SignupForm struct {
	FirstName       string `form:"f_name" binding:"required"`
	LastName        string `form:"l_name" binding:"required"`
	Organization    string `form:"organization" binding:"required"`
	Email           string `form:"email" binding:"required"`
	Password        string `form:"password" binding:"required"`
	PasswordConfirm string `form:"password_confirm" binding:"required"`
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
		key, err := scrypt.DerivePassphrase(form.Password, 32)
		if err != nil {
			fmt.Errorf("Error returned: %s\n", err)
		}
		logrus.WithFields(logrus.Fields{
			"FirstName":    form.FirstName,
			"LastName":     form.LastName,
			"Organization": form.Organization,
			"Email":        form.Email,
			"Password":     form.Password,
			"Hash":         key,
		}).Info("User signup")
		// logie should do something about this
		ctx := pongo2.Context{
			"title": "Pair",
		}
		c.HTML(http.StatusOK, "templates/pages/signup.html", ctx)
	}
}

func stringValidator(str string, length int) bool {
	if len(str) >= length {
		return true
	} else {
		return false
	}
}

func emailValidator(email string) bool {
	// did we get a string worth checking?
	if len(email) > 0 {
		// check it
		match, err := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, email)
		if err != nil {
			panic(err)
		}
		if match {
			// passed regex
			return true
		} else {
			// falied regex
			return false
		}
	} else {
		// not a string worth checking
		return false
	}
}
