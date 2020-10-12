package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jwt-example/cmd/api/model"
	cm "github.com/jwt-example/internal/cache/model"
	"net/http"
)

func (ctrl *Controller) CreateTodo(c *gin.Context) {
	var td *model.Todo

	if err := c.ShouldBindJSON(&td); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	tokenAuth, err := ctrl.Service.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	userId, err := ctrl.Service.Cache.FetchAuth(&cm.AccessDetails{
		AccessUuid: tokenAuth.AccessUuid,
		UserId:     tokenAuth.UserId,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	td.UserID = userId

	c.JSON(http.StatusCreated, td)
}

