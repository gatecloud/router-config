package controllers

import (
	"net/http"
	"router-config/configs"

	"gopkg.in/go-playground/validator.v8"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
)

type Controller interface {
	Init(db *gorm.DB, validate *validator.Validate, store *sessions.FilesystemStore, model interface{})
	Post(ctx *gin.Context)
	Patch(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	RedirectError(ctx *gin.Context, statusCode int, err error)
}

type Control struct {
	DB        *gorm.DB
	Validator *validator.Validate
	Model     interface{}
	Store     *sessions.FilesystemStore
}

func (ctrl *Control) Init(db *gorm.DB, validate *validator.Validate, store *sessions.FilesystemStore, model interface{}) {
	ctrl.DB = db
	ctrl.Validator = validate
	ctrl.Store = store
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

func (ctrl *Control) CreateAWSSession() (*session.Session, error) {
	var err error
	creds := credentials.NewStaticCredentials(configs.Configuration.AWSS3Key,
		configs.Configuration.AWSS3Secret,
		"")
	if _, err = creds.Get(); err != nil {
		return nil, err
	}

	config := aws.NewConfig().
		WithRegion(configs.Configuration.AWSS3Region).
		WithCredentials(creds)

	return session.NewSession(config)
}

// RedirectError redirects to the error page
func (ctrl *Control) RedirectError(ctx *gin.Context, statusCode int, err error) {
	// ctx.HTML(statusCode, "error.html", gin.H{
	// 	"StatusCode": statusCode,
	// 	"Error":      err.Error(),
	// })

	ctx.JSON(statusCode, gin.H{
		"StatusCode": statusCode,
		"Error":      err.Error(),
	})
	ctx.Next()
}

func (ctrl *Control) IsAuthenticated(ctx *gin.Context) error {
	session, err := ctrl.Store.Get(ctx.Request, "auth-session")
	if err != nil {
		return err
	}

	if _, ok := session.Values["profile"]; !ok {
		ctx.Redirect(http.StatusSeeOther, "/index")
		ctx.Abort()
		// return errors.New("no profiles session")
	} else {
		ctx.Next()
	}
	return nil
}
