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
	"github.com/go-redis/redis"
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

	fmt.Println("v5")

	sr := &libRoute.Resource{
		DB:          db,
		Validator:   validator,
		RedisClient: &redis.Client{},
	}
	apiRouter := r.Group("/api")
	libRoute.DistributeRouters(apiRouter, routers.RouteMap["api"], sr)

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
