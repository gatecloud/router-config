package controllers

import (
	"net/http"
	"net/url"
	"router-config/configs"

	"github.com/gin-gonic/gin"
)

type LogoutController struct {
	Control
}

func (ctrl *LogoutController) GetAll(ctx *gin.Context) {
	domain := configs.Configuration.Auth0Domain

	var Url *url.URL
	Url, err := url.Parse("https://" + domain)

	if err != nil {
		panic("boom")
	}

	Url.Path += "/v2/logout"
	parameters := url.Values{}
	parameters.Add("returnTo", configs.Configuration.RedirectIndex)
	parameters.Add("client_id", configs.Configuration.Auth0ClientID)
	Url.RawQuery = parameters.Encode()

	ctx.Redirect(http.StatusTemporaryRedirect, Url.String())
}
