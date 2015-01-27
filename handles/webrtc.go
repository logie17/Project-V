package handles

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebrtcGetHandler(c *gin.Context) {
	ctx := pongo2.Context{
		"title": "pairparty.io",
		//"roomname": c.Params.ByName("roomname"),
		"room": c.Request.URL.Query().Get("room"),
		"username": c.Request.URL.Query().Get("username"),
	}
	c.HTML(http.StatusOK, "templates/pages/webrtc.html", ctx)
}

