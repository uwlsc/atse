package magazine

import "magazine_api/models"

type MagazineBase struct {
	MagazineCode *string `json:"magazine_code"`
	IssueCode    *string `json:"issue_code"`
	Placement    *string `json:"placement"`
	Remarks      *string `json:"remarks"`
}

type Magazine struct {
	MagazineBase
	models.Base
	models.BaseDate
	models.BaseCreatedBy
}
