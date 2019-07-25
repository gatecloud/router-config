package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"router-config/configs"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type CallbackController struct {
	Control
}

func (ctrl *CallbackController) GetAll(ctx *gin.Context) {
	domain := configs.Configuration.Auth0Domain

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

	state := ctx.Request.URL.Query().Get("state")
	session, err := ctrl.Store.Get(ctx.Request, "state")
	if err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	if state != session.Values["state"] {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, errors.New("Invalid state parameter"))
		return
	}

	code := ctx.Request.URL.Query().Get("code")

	token, err := conf.Exchange(context.TODO(), code)
	if err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	// Getting now the userInfo
	client := conf.Client(context.TODO(), token)
	resp, err := client.Get("https://" + domain + "/userinfo")
	if err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	defer resp.Body.Close()

	var profile map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	session, err = ctrl.Store.Get(ctx.Request, "auth-session")
	if err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	session.Values["id_token"] = token.Extra("id_token")
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile
	err = session.Save(ctx.Request, ctx.Writer)
	if err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	// Redirect to logged in page
	ctx.Redirect(http.StatusSeeOther, configs.Configuration.RedirectHomePage)
}
