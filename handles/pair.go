package handles

import (
	"github.com/gin-gonic/gin"
	"github.com/flosch/pongo2"
	"net/http"
)

func PairGetHandler(c *gin.Context) {
	ctx := pongo2.Context{
		"title": "Pair",
	}
	c.HTML(http.StatusOK, "templates/pages/pair.html", ctx)
}
