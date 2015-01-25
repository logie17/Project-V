package middleware

import (
	"github.com/Shaked/gomobiledetect"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func IsMobile() gin.HandlerFunc {
	return func(c *gin.Context) {
		detect := mobiledetect.NewMobileDetect(c.Request, nil)
		if detect.IsMobile() {
			var ctx map[string]interface{} = pongo2.Context{
				"mobile": true,
			}
			c.Set("template_data", ctx)
		}
		c.Next()
	}
}
