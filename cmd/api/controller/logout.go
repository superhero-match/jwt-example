package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctrl *Controller) Logout(c *gin.Context) {
	au, err := ctrl.Service.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	deleted, delErr := ctrl.Service.DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	c.JSON(http.StatusOK, "Successfully logged out")
}

