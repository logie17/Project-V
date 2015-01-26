package handles

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func PairGetHandler(c *gin.Context) {
	ctx := c.MustGet("global_data").(pongo2.Context)
	c.HTML(http.StatusOK, "templates/pages/pair.html", ctx)
}
