package models

import (
	"github.com/google/uuid"
)

type SalaryFormat int

const (
	PerPiece SalaryFormat = iota + 1
	Hourly
	Daily
	Monthly
	Yearly
)

type TailorRate struct {
	ItemId *uuid.UUID `json:"product_id"`
	Rate   *float32   `json:"rate"`
}

type SalaryBase struct {
	EmployeeId *uuid.UUID `json:"employee_id"`
	CompanyId  *string    `json:"company_id"`

	Amount       *float32      `json:"amount"`
	SalaryFormat *SalaryFormat `json:"salary_format"`

	TailorRate []TailorRate `json:"tailor_rate"`

	EffectiveFrom *Month `json:"effective_from"`
	EffectiveTo   *Month `json:"effective_to"`
}

type Salary struct {
	Base
	BaseDate
	BaseCreatedBy
	SalaryBase
}

/*
WITH RECURSIVE item_view AS (
	SELECT it.id, it.item_name, it.item_code, 0 AS level
	FROM items it
	UNION ALL
		SELECT it2.id, it2.item_name, it2.item_code, item_view.level + 1
		FROM items it2
		JOIN item_view ON it2.id = item_view.id
)
SELECT *
FROM (SELECT item_view.*,
             MAX(level) OVER (PARTITION BY id) AS maxlevel
      FROM item_view
     ) item_view
WHERE level = maxlevel;
*/
