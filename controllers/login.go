package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"router-config/configs"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type LoginController struct {
	Control
}

func (ctrl *LoginController) GetAll(ctx *gin.Context) {
	domain := configs.Configuration.Auth0Domain
	aud := configs.Configuration.Auth0Audience

	conf := &oauth2.Config{
		ClientID:     configs.Configuration.Auth0ClientID,
		ClientSecret: configs.Configuration.Auth0ClientSecret,
		RedirectURL:  configs.Configuration.Auth0CallbackURL,
		Scopes:       []string{"openid", "profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + domain + "/authorize",
			TokenURL: "https://" + domain + "/oauth/token",
		},
	}

	if aud == "" {
		aud = "https://" + domain + "/userinfo"
	}

	fmt.Println("1")
	// Generate random state
	b := make([]byte, 32)
	rand.Read(b)
	state := base64.StdEncoding.EncodeToString(b)

	session, err := ctrl.Store.Get(ctx.Request, "state")
	if err != nil {
		fmt.Println(err)
		// ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		// return
	}
	fmt.Println("2")
	session.Values["state"] = state
	err = session.Save(ctx.Request, ctx.Writer)
	if err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("3")
	audience := oauth2.SetAuthURLParam("audience", aud)
	url := conf.AuthCodeURL(state, audience)

	fmt.Println("====", url)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}
