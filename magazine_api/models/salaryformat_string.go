// Code generated by "stringer -type SalaryFormat models/salary.go"; DO NOT EDIT.

package models

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PerPiece-1]
	_ = x[Hourly-2]
	_ = x[Daily-3]
	_ = x[Monthly-4]
	_ = x[Yearly-5]
}

const _SalaryFormat_name = "PerPieceHourlyDailyMonthlyYearly"

var _SalaryFormat_index = [...]uint8{0, 8, 14, 19, 26, 32}

func (i SalaryFormat) String() string {
	i -= 1
	if i < 0 || i >= SalaryFormat(len(_SalaryFormat_index)-1) {
		return "SalaryFormat(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _SalaryFormat_name[_SalaryFormat_index[i]:_SalaryFormat_index[i+1]]
}
