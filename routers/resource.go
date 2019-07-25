package routers

import (
	"errors"
	"reflect"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v8"
)

type RoResource struct {
	DB        *gorm.DB                  `gatecloud:"initparam"`
	Validator *validator.Validate       `gatecloud:"initparam"`
	Store     *sessions.FilesystemStore `gatecloud:"initparam"`
}

// Resources returns all resources that tag as initparam
func (r *RoResource) Resources() ([]reflect.Value, error) {
	rf := reflect.ValueOf(r).Elem()
	if !rf.IsValid() {
		return []reflect.Value{}, errors.New("reflect error")
	}

	args := make([]reflect.Value, rf.NumField())
	for i := 0; i < rf.NumField(); i++ {
		if rf.Type().Field(i).Tag.Get("gatecloud") == "initparam" {
			args[i] = rf.Field(i)
		}
	}
	return args, nil
}
