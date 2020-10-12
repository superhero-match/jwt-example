package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctrl *Controller) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := ctrl.Service.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()

			return
		}

		c.Next()
	}
}

