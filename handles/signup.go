package handles

import (
	"net/http"
	"regexp"

	"github.com/Sirupsen/logrus"
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

// SignupGetHandler is the GET handler for the signup page.  I does not
// contain any logic but simply renders the template.
func SignupGetHandler(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.MustGet("global_data").(pongo2.Context)
		c.HTML(http.StatusOK, "templates/pages/signup.html", ctx)
	}
}

func SignupPostHandler(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.MustGet("global_data").(pongo2.Context)
		// lets checkout the form
		var form SignupForm
		c.BindWith(&form, binding.Form)

		println("12341234")

		errors := validate(form)
		if errors != nil {
			for k, v := range errors {
				ctx[k] = v
			}
			ctx["f_name_val"] = form.FirstName
			ctx["l_name_val"] = form.LastName
			ctx["organization_val"] = form.Organization
			ctx["email_val"] = form.Email
			ctx["password"] = form.Password
			ctx["password_confirm"] = form.PasswordConfirm
			println("sdfsdf")
			logrus.Info("error")
			c.HTML(http.StatusOK, "templates/pages/signup.html", ctx)
			return
		}
		// form is valid: logie needs to save this shit somewhere
		// also, you will need the below, thats their password
		/*
			key, err := scrypt.DerivePassphrase(form.Password, 32)
			if err != nil {
				fmt.Errorf("Scrypt error returned: %s\n", err)
			}
		*/
		c.HTML(http.StatusOK, "templates/pages/pair.html", ctx)
		logrus.Info("Yay! a new user signup")
	}
}

// validate is a private function that pimps out to the
// other validator fields to check each and returns a
// map of the error messages or nil if the form is valid
func validate(form SignupForm) map[string]string {
	var errors map[string]string = make(map[string]string)
	var valid bool = true
	if !stringValidator(form.FirstName, 1) {
		errors["f_name"] = "Please enter a first name"
		valid = false
	}
	if !stringValidator(form.LastName, 1) {
		errors["l_name"] = "Please enter a last name"
		valid = false
	}
	if !stringValidator(form.Organization, 1) {
		errors["organization"] = "Please enter an organization"
		valid = false
	}
	if !stringValidator(form.Email, 1) {
		errors["email"] = "Please enter a valid email address"
		valid = false
	}
	if !stringValidator(form.Password, 8) {
		errors["password"] = "Please enter at least 8 characters"
		valid = false
	} else {
		if form.PasswordConfirm != form.Password {
			errors["password_confirm"] = "Passwords must match"
			valid = false
		}
	}
	if valid {
		return nil
	} else {
		return errors
	}
}

// stringValidator is a private function that checks the
// length of a string
func stringValidator(str string, length int) bool {
	if len(str) >= length {
		return true
	} else {
		return false
	}
}

// emailValidator is a private function that
// checks emails against a simple regex.  Note this
// is not checking the RFC 5322 regex
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
