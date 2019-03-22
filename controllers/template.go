package controllers

import (
	"errors"
	"fmt"
	"net/http"
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

	entity.DeletedAt = nil
	if err := ctrl.DB.Create(&entity).Error; err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, entity)
	return
}

func (ctrl *TemplateController) Delete(ctx *gin.Context) {
	var (
		chkEntity models.Template
	)

	idStr := ctx.Params.ByName("id")
	if idStr == "" {
		err := errors.New("id is required")
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if ctrl.DB.Where("id = ?", idStr).Find(&chkEntity).RecordNotFound() {
		err := errors.New("template not found")
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ctrl.DB.Unscoped().Delete(&chkEntity).Error; err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, chkEntity)
	return
}

func (ctrl *TemplateController) GetByID(ctx *gin.Context) {
	var (
		chkEntity models.Template
	)

	name := ctx.Params.ByName("id")
	if name == "" {
		err := errors.New("id is required")
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if ctrl.DB.Where("id = ?", name).Find(&chkEntity).RecordNotFound() {
		err := errors.New("template not found")
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, chkEntity)
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
