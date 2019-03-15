package controllers

import (
	"net/http"

	"gopkg.in/go-playground/validator.v8"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type Controller interface {
	Init(*gorm.DB, *validator.Validate)
	Post(ctx *gin.Context)
	Patch(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type Control struct {
	DB        *gorm.DB
	Validator *validator.Validate
	Model     interface{}
}

func (ctrl *Control) Init(db *gorm.DB, validate *validator.Validate, redis *redis.Client, model interface{}) {
	ctrl.DB = db
	ctrl.Validator = validate
	ctrl.Model = model
}

func (ctrl *Control) Post(ctx *gin.Context) {
	ctx.JSON(http.StatusMethodNotAllowed, nil)
	return
}

func (ctrl *Control) Patch(ctx *gin.Context) {
	ctx.JSON(http.StatusMethodNotAllowed, nil)
	return
}

func (ctrl *Control) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusMethodNotAllowed, nil)
	return
}

func (ctrl *Control) GetByID(ctx *gin.Context) {
	ctx.JSON(http.StatusMethodNotAllowed, nil)
	return
}

func (ctrl *Control) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusMethodNotAllowed, nil)
	return
}

// RedirectError redirects to the error page
func (ctrl *Control) RedirectError(ctx *gin.Context, statusCode int, err error) {
	ctx.HTML(statusCode, "error.html", gin.H{
		"Error": err.Error(),
	})
}
