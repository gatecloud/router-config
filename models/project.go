package models

import "github.com/google/uuid"

type Project struct {
	Model
	Name         string `validate:"required"`
	RouterGroups string `validate:"required"`
	UserID       uuid.UUID
	User         User `json:"-" sql:"-"`
}
