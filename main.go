package main

import (
	"net/http"
	"router-config/logic"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.StaticFS("/public", http.Dir("public"))
	r.StaticFS("/groups", http.Dir("groups"))
	r.LoadHTMLGlob("templates/*")

	r.GET("/index", func(ctx *gin.Context) {
		groups, err := logic.ToList()
		if err != nil {
			RedirectError(ctx, http.StatusInternalServerError, err)
			return
		}
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Groups": groups,
		})
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
		filename := strings.TrimSpace(ctx.Param("filename"))
		routerTemplate, err := logic.Load(filename)
		if err != nil {
			RedirectError(ctx, http.StatusInternalServerError, err)
			return
		}

		// ctx.JSON(http.StatusOK, gin.H{
		// 	"RouterTemplate": routerTemplate,
		// })
		ctx.JSON(http.StatusOK, gin.H{
			"resources":     routerTemplate.Resources,
			"methods":       routerTemplate.Methods,
			"version":       routerTemplate.Version,
			"proxyschema":   routerTemplate.ProxySchema,
			"proxypass":     routerTemplate.ProxyPass,
			"proxyversion":  routerTemplate.ProxyVersion,
			"customconfigs": routerTemplate.CustomConfigs,
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

	r.Run(":7000")
}

func RedirectError(ctx *gin.Context, statusCode int, err error) {
	ctx.HTML(statusCode, "error.html", gin.H{
		"Error": err.Error(),
	})
}
