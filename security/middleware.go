package security

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"ofcode.dev/labala-backend/security/SJWT"
)

func JWTBarrier(minimumLevel uint) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("TOKEN")
		if token == "" {
			c.JSON(http.StatusUnauthorized, "token not provided")
			c.Abort()
			return
		}

		jwtToken, err := SJWT.Validate(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		_, body := SJWT.Translate(jwtToken)

		if body.Acs < minimumLevel {
			c.JSON(http.StatusUnauthorized, "need higher access level")
			c.Abort()
			return
		}

		sess := sessions.Default(c)
		sess.Set("ACCESS", body)

		c.Next()
	}
}
