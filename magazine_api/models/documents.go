package models

import (
	"magazine_api/lib"

	"github.com/google/uuid"
)

type DocumentBase struct {
	CreatorId *uuid.UUID `json:"creator_id"`
	CompanyId *string    `json:"company_id"`

	DocumentType *string        `json:"document_type"`
	DocumentURL  *lib.SignedURL `json:"url"`
}

type Document struct {
	DocumentBase
	Base
	BaseDate
}
