package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"router-config/configs"
	"router-config/models"

	"github.com/gin-gonic/gin"
)

type TemplateController struct {
	Control
}

func (ctrl *TemplateController) Post(ctx *gin.Context) {
	var (
		entity      models.Template
		chkTemplate models.Template
	)

	if err := ctx.Bind(&entity); err != nil {
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ctrl.Validator.Struct(entity); err != nil {
		fmt.Println(err)
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if !ctrl.DB.Where("project_name = ? and router_group = ? and template_name = ?",
		entity.ProjectName, entity.RouterGroup, entity.TemplateName).
		Find(&chkTemplate).
		RecordNotFound() {
		err := fmt.Errorf("%s/%s is existed", entity.ProjectName, entity.TemplateName)
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	routerTemplate, err := entity.Convert2RouterTemplate()
	if err != nil {
		fmt.Println(err)
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	routers := routerTemplate.GenerateRouters()
	body, err := json.MarshalIndent(routers, "", "\t")
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

	entity.DeletedAt = nil
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

func (ctrl *TemplateController) GetAll(ctx *gin.Context) {
	var (
		entities []models.Template
	)

	if ctrl.DB.Find(&entities).RecordNotFound() {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	ctx.JSON(http.StatusOK, entities)
	return
}
