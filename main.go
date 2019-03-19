package main

import (
	"encoding/json"
	"log"
	"net/http"
	"router-config/configs"
	"router-config/logic"
	"router-config/routers"
	"router-config/validations"
	"strings"
	"time"

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
	r.StaticFS("/groups", http.Dir("groups"))
	r.StaticFS("/files", http.Dir("files"))
	r.LoadHTMLGlob("templates/*")

	sr := &libRoute.Resource{
		DB:          db,
		Validator:   validator,
		RedisClient: &redis.Client{},
	}
	apiRouter := r.Group("/api")
	libRoute.DistributeRouters(apiRouter, routers.RouteMap["api"], sr)

	r.GET("/index", func(ctx *gin.Context) {
		// files, err := logic.ToList("files/")
		// if err != nil {
		// 	RedirectError(ctx, http.StatusInternalServerError, err)
		// 	return
		// }

		// groups, err := logic.ToList("groups/")
		// if err != nil {
		// 	RedirectError(ctx, http.StatusInternalServerError, err)
		// 	return
		// }

		ctx.HTML(http.StatusOK, "index.html", gin.H{
		// "Groups": groups,
		// "Files":  files,
		})
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
	r.POST("/AddTemplate", func(ctx *gin.Context) {
		r := logic.RouterTemplate{}
		resources := ctx.PostForm("resource")
		methods := ctx.PostForm("method")
		version := ctx.PostForm("version")
		proxySchema := ctx.PostForm("proxy_schema")
		proxyPass := ctx.PostForm("proxy_pass")
		proxyVersion := ctx.PostForm("proxy_version")
		customConfig := ctx.PostForm("custom_config")
		if err := r.Parse(resources, methods, version, proxySchema, proxyPass, proxyVersion, customConfig); err != nil {
			RedirectError(ctx, http.StatusInternalServerError, err)
			return
		}
		if err := r.Save(ctx.PostForm("filename")); err != nil {
			RedirectError(ctx, http.StatusInternalServerError, err)
			return
		}
		ctx.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.GET("/GetTemplate/:filename", func(ctx *gin.Context) {
		routerTemplate := logic.RouterTemplate{}
		filename := strings.TrimSpace(ctx.Param("filename"))
		err := routerTemplate.Load(filename)
		if err != nil {
			RedirectError(ctx, http.StatusInternalServerError, err)
			return
		}

		customConfigs, err := json.MarshalIndent(routerTemplate.CustomConfigs, "", "\t")
		if err != nil {
			RedirectError(ctx, http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"resources":     routerTemplate.Resources,
			"methods":       routerTemplate.Methods,
			"version":       routerTemplate.Version,
			"proxyschema":   routerTemplate.ProxySchema,
			"proxypass":     routerTemplate.ProxyPass,
			"proxyversion":  routerTemplate.ProxyVersion,
			"customconfigs": string(customConfigs),
		})
	})

	r.POST("/Export", func(ctx *gin.Context) {
		fileContent := ctx.PostForm("content")
		fileList := strings.Split(fileContent, ";")
		filename := ctx.PostForm("filename")
		if err := logic.Export(filename, fileList); err != nil {
			RedirectError(ctx, http.StatusInternalServerError, err)
			return
		}
		ctx.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.POST("/DeleteGroup", func(ctx *gin.Context) {
		fileContent := ctx.PostForm("content")
		fileList := strings.Split(fileContent, ";")
		if err := logic.DeleteGroup(fileList); err != nil {
			RedirectError(ctx, http.StatusInternalServerError, err)
			return
		}
		time.Sleep(1 * time.Second)
		ctx.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.Run(configs.Configuration.Port)
}

// RedirectError redirects to the error page
func RedirectError(ctx *gin.Context, statusCode int, err error) {
	ctx.HTML(statusCode, "error.html", gin.H{
		"Error": err.Error(),
	})
}
