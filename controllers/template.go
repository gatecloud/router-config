package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"router-config/configs"
	"router-config/models"
	"strings"

	"github.com/gin-gonic/gin"
)

type TemplateController struct {
	Control
}

func (ctrl *TemplateController) Post(ctx *gin.Context) {
	var (
		entity         models.Template
		routerTemplate models.RouterTemplate
		chkTemplate    models.Template
	)

	if err := ctx.Bind(&entity); err != nil {
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ctrl.Validator.Struct(entity); err != nil {
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if !ctrl.DB.Where("project_name = ? and template_name = ?", entity.ProjectName, entity.TemplateName).
		Find(&chkTemplate).
		RecordNotFound() {
		err := fmt.Errorf("%s/%s is existed", entity.ProjectName, entity.TemplateName)
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	routerTemplate.Resources = strings.Split(filterString(entity.Resource), " ")
	routerTemplate.Methods = strings.Split(filterString(entity.Method), " ")
	routerTemplate.Version = entity.Version
	routerTemplate.ProxySchema = entity.ProxySchema
	routerTemplate.ProxyPass = entity.ProxyPass
	routerTemplate.ProxyVersion = entity.ProxyVersion
	if err := json.Unmarshal([]byte(entity.CustomConfig), &routerTemplate.CustomeConfig); err != nil {
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	body, err := json.MarshalIndent(routerTemplate, "", "\t")
	if err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	uploader, err := ctrl.CreateAWSUploader()
	if err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	if err := entity.UploadFile(uploader, configs.Configuration.AWSS3Domain, body); err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	if err := ctrl.DB.Create(&entity).Error; err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, entity)
	return
}

func (ctrl *TemplateController) GetByID(ctx *gin.Context) {
	var (
		entities []models.Project
	)

	if err := ctrl.DB.Find(&entities).Error; err != nil {
		// if ctrl.DB.Find(&entities).RecordNotFound() {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}
	fmt.Println("---", entities)

	ctx.JSON(http.StatusOK, entities)
	return
}

func filterString(src string) string {
	dst := strings.Replace(src, "\r\n", "", -1)
	dst = strings.Replace(dst, "\"", "", -1)
	dst = strings.Replace(dst, "\t", "", -1)
	dst = strings.Replace(dst, "\n", "", -1)
	dst = strings.Replace(dst, "\r", "", -1)
	dst = strings.Replace(dst, " ", "", -1)
	dst = strings.Replace(dst, ",", " ", -1)
	return dst
}
