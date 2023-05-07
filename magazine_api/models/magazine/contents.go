package magazine

import (
	"magazine_api/models"

	"github.com/google/uuid"
)

type ContentBase struct {
	ContentCode    *string    `json:"content_code"`
	StoryCode      *uuid.UUID `json:"story_code"`
	PhotographCode *string    `json:"photo_code"`
	Remarks        *string    `json:"remarks"`
}

type Content struct {
	ContentBase
	models.Base
	models.BaseDate
	models.BaseCreatedBy
}
