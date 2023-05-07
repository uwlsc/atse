package magazine

import (
	"magazine_api/lib"
	"magazine_api/models"

	"github.com/google/uuid"
)

type PhotographBase struct {
	PhotographID   *uuid.UUID `json:"photo_id"`
	PhotographCode *string    `json:"photo_code"`

	PhotographTitle *string        `json:"photo_title"`
	PhotographType  *string        `json:"photo_type"`
	DocumentURL     *lib.SignedURL `json:"url"`

	Remarks *string `json:"remarks"`
}

type Photograph struct {
	PhotographBase
	models.Base
	models.BaseDate
	models.BaseCreatedBy
}
