package magazine

import (
	"magazine_api/lib"
	"magazine_api/models"
)

type AdvertBase struct {
	AdvertCode    *string        `json:"advert_code"`
	AdvertTitle   *string        `json:"advert_title"`
	AdvertContent *string        `json:"advert_content"`
	AdvertType    *string        `json:"advert_type"`
	AdvertURL     *lib.SignedURL `json:"url"`

	Remarks *string `json:"remarks"`
}

type Advert struct {
	models.Base
	models.BaseDate
	models.BaseCreatedBy
	AdvertBase
}
