package magazine

import (
	"magazine_api/models"

	"github.com/google/uuid"
)

type StoryBase struct {
	StoryID   *uuid.UUID `json:"story_id"`
	StoryCode *string    `json:"story_code"`
	CreatorId *uuid.UUID `json:"creator_id"`

	StoryTitle   *string `json:"story_title"`
	StoryType    *string `json:"story_type"`
	StoryContent *string `json:"story_content"`

	Remarks *string `json:"remarks"`
}

type Story struct {
	StoryBase
	models.Base
	models.BaseDate
}
