package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func isAuthenticated(c *gin.Context) {
	session, err := store.Get(c.Request, "flash-session")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
	var username = session.Values["username"]
	if username == nil {
		session.AddFlash("Login failed!", "message")
		session.Save(c.Request, c.Writer)
		c.Fail(http.StatusUnauthorized, errors.New("Unauthorized")) // idk why this is needed but it is
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}
