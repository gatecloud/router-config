package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"router-config/configs"
	"router-config/models"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	Control
}

func (ctrl *FileController) Post(ctx *gin.Context) {
	var (
		entity  models.File
		chkFile models.File
		routers []models.Router
	)

	if err := ctx.BindJSON(&entity); err != nil {
		fmt.Println(err)
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}
	fmt.Println("---", entity)

	if err := ctrl.Validator.Struct(entity); err != nil {
		fmt.Println(err)
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if !ctrl.DB.Where("name = ?", entity.Name).
		Find(&chkFile).
		RecordNotFound() {
		err := fmt.Errorf("%s is existed", entity.Name)
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	for i, v := range entity.Templates {
		var chkTemplate models.Template
		if ctrl.DB.Where("id = ?", v.ID).Find(&chkTemplate).RecordNotFound() {
			err := fmt.Errorf("template %s not found", entity.Name)
			ctrl.RedirectError(ctx, http.StatusBadRequest, err)
			return
		}
		entity.Templates[i] = chkTemplate

		routerTemplate, err := chkTemplate.Convert2RouterTemplate()
		if err != nil {
			ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
			return
		}
		routers = append(routers, routerTemplate.GenerateRouters()...)
	}

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

func (ctrl *FileController) GetByID(ctx *gin.Context) {
	var (
		entities []models.File
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

func (ctrl *FileController) GetAll(ctx *gin.Context) {
	var (
		entities []models.File
	)

	if ctrl.DB.Find(&entities).RecordNotFound() {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	ctx.JSON(http.StatusOK, entities)
	return
}
