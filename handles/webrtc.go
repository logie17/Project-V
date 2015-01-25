package handles

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebrtcGetHandler(c *gin.Context) {
	ctx := pongo2.Context{
		"title": "pairparty.io",
	}
	c.HTML(http.StatusOK, "templates/pages/webrtc.html", ctx)
}

