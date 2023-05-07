package magazine

import "magazine_api/models"

type IssuseBase struct {
	IssueCode   *string `json:"issue_code"`
	ContentCode *string `json:"content_code"`
	AdvertCode  *string `json:"advert_code"`
	Remarks     *string `json:"remarks"`
}

type MagazineIssue struct {
	IssuseBase
	models.Base
	models.BaseDate
	models.BaseCreatedBy
}
