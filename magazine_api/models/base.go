package models

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID uuid.UUID `json:"id"`
}

type BaseDate struct {
	CreatedOn *time.Time `json:"created_on" form:"created_on"`
	UpdatedOn *time.Time `json:"updated_on" form:"updated_on"`
	DeletedOn *time.Time `json:"deleted_on" form:"deleted_on"`
}

type BaseCreatedBy struct {
	CreatedBy   *uuid.UUID `json:"created_by"`
	CreatorName *string    `json:"creator_name"`
	UpdatedBy   *uuid.UUID `json:"updated_by"`
	UpdatedName *string    `json:"updator_name"`
	DeletedBy   *uuid.UUID `json:"deleted_by"`
	DeletedName *string    `json:"deletor_name"`
}
