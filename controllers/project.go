package controllers

import (
	"errors"
	"net/http"
	"router-config/models"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	Control
}

func (ctrl *ProjectController) Post(ctx *gin.Context) {
	var (
		entity models.Project
	)

	if err := ctx.Bind(&entity); err != nil {
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ctrl.Validator.Struct(entity); err != nil {
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
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

func (ctrl *ProjectController) Patch(ctx *gin.Context) {
	var (
		entity     models.Project
		chkProject models.Project
	)

	if err := ctx.Bind(&entity); err != nil {
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ctrl.Validator.Struct(entity); err != nil {
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if ctrl.DB.Where("id = ?", entity.ID).Find(&chkProject).RecordNotFound() {
		err := errors.New("proejct not found")
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if entity.RouterGroups != "" {
		chkProject.RouterGroups = entity.RouterGroups
	}

	if err := ctrl.DB.Save(&chkProject).Error; err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, chkProject)
	return
}

func (ctrl *ProjectController) Delete(ctx *gin.Context) {
	var (
		chkEntity models.Project
	)

	idStr := ctx.Params.ByName("id")
	if idStr == "" {
		err := errors.New("id is required")
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if ctrl.DB.Where("id = ?", idStr).Find(&chkEntity).RecordNotFound() {
		err := errors.New("project not found")
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

func (ctrl *ProjectController) GetByID(ctx *gin.Context) {
	var (
		chkEntity models.Project
	)

	id := ctx.Params.ByName("id")
	if id == "" {
		err := errors.New("id is required")
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	if ctrl.DB.Where("id = ?", id).Find(&chkEntity).RecordNotFound() {
		err := errors.New("project not found")
		ctrl.RedirectError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, chkEntity)
	return
}

func (ctrl *ProjectController) GetAll(ctx *gin.Context) {
	var (
		entities []models.Project
	)

	if err := ctrl.DB.Find(&entities).Error; err != nil {
		// if ctrl.DB.Find(&entities).RecordNotFound() {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	ctx.JSON(http.StatusOK, entities)
	return
}
