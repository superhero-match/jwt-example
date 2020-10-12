package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jwt-example/cmd/api/model"
	cm "github.com/jwt-example/internal/cache/model"
	"net/http"
)

//A sample use
var user = model.User{
	ID:       1,
	Username: "username",
	Password: "password",
}

func (ctrl *Controller) Login(c *gin.Context) {
	var u model.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	token, err := ctrl.Service.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := ctrl.Service.CreateAuth(user.ID, cm.TokenDetails{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		AccessUuid:   token.AccessUuid,
		RefreshUuid:  token.RefreshUuid,
		AtExpires:    token.AtExpires,
		RtExpires:    token.RtExpires,
	})
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}

	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}

	c.JSON(http.StatusOK, tokens)
}
