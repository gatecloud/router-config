package main

import (
	"fmt"
	"log"
	"net/http"
	"router-config/configs"
	"router-config/routers"
	"router-config/validations"
	"strconv"

	libRoute "github.com/gatecloud/webservice-library/route"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	if err := configs.InitConfig(); err != nil {
		log.Fatal(err)
	}

	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	validator := validations.InitValidation()

	if configs.Configuration.Production {
		gin.SetMode(gin.ReleaseMode)
		db.LogMode(false)
	} else {
		gin.SetMode(gin.DebugMode)
		db.LogMode(true)
	}
	r := gin.Default()
	r.HandleMethodNotAllowed = true

	r.StaticFS("/public", http.Dir("public"))
	r.StaticFS("/templates", http.Dir("templates"))
	r.LoadHTMLGlob("templates/*")

	fmt.Println("v6")

	store := sessions.NewFilesystemStore("", []byte("roconfig-secret"))
	var roResource routers.RoResource
	roResource.DB = db
	roResource.Validator = validator
	roResource.Store = store

	apiRouter := r.Group("/api")
	libRoute.DistributeRouters(apiRouter, routers.RouteMap["api"], &roResource)

	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/home", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})

	r.GET("/project", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "project.html", nil)
	})

	r.GET("/template", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "template.html", nil)
	})

	r.GET("/error", func(ctx *gin.Context) {
		var statusCode int
		errMsg := ctx.Query("Error")
		code := ctx.Query("StatusCode")
		if code != "" {
			statusCode, err = strconv.Atoi(code)
			if err != nil {
				log.Fatal(err)
			}
		}

		ctx.HTML(statusCode, "error.html", gin.H{
			"Error":      errMsg,
			"StatusCode": statusCode,
		})
	})

	r.Run(configs.Configuration.Port)
}
