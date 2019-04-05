package controllers

import (
	"encoding/json"
	"errors"
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

	if err := ctrl.Validator.Struct(entity); err != nil {
		fmt.Println(err)
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
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

	sess, err := ctrl.CreateAWSSession()
	if err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	if err := entity.UploadFile(sess, configs.Configuration.AWSS3Domain, body); err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	entity.DeletedAt = nil

	if ctrl.DB.Where("name = ?", entity.Name).
		Find(&chkFile).
		RecordNotFound() {
		if err := ctrl.DB.Create(&entity).Error; err != nil {
			ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
			return
		}
	}

	ctx.JSON(http.StatusOK, entity)
	return
}

func (ctrl *FileController) Delete(ctx *gin.Context) {
	var (
		chkEntity models.File
	)

	idStr := ctx.Params.ByName("id")
	if idStr == "" {
		err := errors.New("id is required")
		fmt.Println(err)
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	tx := ctrl.DB.Begin()
	if tx.Where("id = ?", idStr).Find(&chkEntity).RecordNotFound() {
		tx.Rollback()
		err := errors.New("file not found")
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := tx.Unscoped().Delete(&chkEntity).Error; err != nil {
		tx.Rollback()
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	sess, err := ctrl.CreateAWSSession()
	if err != nil {
		tx.Rollback()
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	if err := chkEntity.DeleteFile(sess, configs.Configuration.AWSS3Domain); err != nil {
		tx.Rollback()
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	if err := tx.Exec("DELETE FROM file_templates WHERE file_id = ?", idStr).Error; err != nil {
		tx.Rollback()
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	tx.Commit()
	ctx.JSON(http.StatusOK, chkEntity)
	return
}

func (ctrl *FileController) GetByID(ctx *gin.Context) {
	var (
		chkEntity models.File
		// routers   []models.Router
	)

	idStr := ctx.Params.ByName("id")
	if idStr == "" {
		err := errors.New("name is required")
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if ctrl.DB.Where("id = ?", idStr).
		Preload("Templates").
		Find(&chkEntity).RecordNotFound() {
		err := errors.New("file not found")
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	// for _, v := range chkEntity.Templates {
	// 	routerTemplate, err := v.Convert2RouterTemplate()
	// 	if err != nil {
	// 		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
	// 		return
	// 	}
	// 	routers = append(routers, routerTemplate.GenerateRouters()...)
	// }

	// body, err := json.MarshalIndent(routers, "", "\t")
	// if err != nil {
	// 	ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
	// 	return
	// }

	// chkEntity.Preview = string(body)
	sess, err := ctrl.CreateAWSSession()
	if err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	if err := chkEntity.GetFileContent(sess, configs.Configuration.AWSS3Domain); err != nil {
		fmt.Println(err)
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, chkEntity)
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
