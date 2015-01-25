package handles

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PairGetHandler(c *gin.Context) {
	ctx := pongo2.Context{
		"title": "Pair",
	}
	global_crap := c.MustGet("WTF_KEVIN").(string)
	println("------")
	println(global_crap)
	println("-------")
	c.HTML(http.StatusOK, "templates/pages/pair.html", ctx)
}
