package controllers

import (
	"fmt"
	"net/http"
	"router-config/models"
	"time"

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

	entity.DeletedAt = &time.Time{}
	if err := ctrl.DB.Create(&entity).Error; err != nil {
		ctrl.RedirectError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, entity)
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
	fmt.Println("---", entities)

	ctx.JSON(http.StatusOK, entities)
	return
}
