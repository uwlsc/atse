package constants

var SalaryFormat = map[string]string{
	"Per Piece": "PerPiece",
	"Hourly":    "Hourly",
	"Daily":     "Daily",
	"Monthly":   "Monthly",
	"Yearly":    "Yearly",
}

var TailorAttribute = map[string]string{
	"Product Id":    "product_id",
	"Category Id":   "category_id",
	"SubCategoryId": "sub_category_id",
	"Rate":          "rate",
}

var Month = map[string]string{
	"Baisakh":  "Baisakh",
	"Jyestha":  "Jyestha",
	"Ashad":    "Ashad",
	"Shrawan":  "Shrawan",
	"Bhadra":   "Bhadra",
	"Ashoj":    "Ashoj",
	"Karthik":  "Karthik",
	"Manghsir": "Manghsir",
	"Poush":    "Poush",
	"Magh":     "Magh",
	"Falgun":   "Falgun",
	"Chaitra":  "Chaitra",
}

var PaidMedium = map[string]string{
	"Cash":           "Cash",
	"Bank":           "Bank",
	"Online":         "Online",
	"Bank To Online": "BankToOnline",
	"Online To Bank": "OnlineToBank",
	"Self":           "Self",
}

var PaidStatus = map[string]string{
	"Complete": "Complete",
	"Pending":  "Pending",
}

var PaidType = map[string]string{
	"Salary Payment":           "SalaryPayment",
	"Advance Salary Payment":   "AdvanceSalaryPayment",
	"Advance Supplier Payment": "AdvanceSupplierPayment",
	"Supplier Payment":         "SupplierPayment",
	"Essential Payment":        "EssentialPayment",
	"Customer Payment":         "CustomerPayment",

	"Advance Salary Payment Return": "AdvanceSalaryPaymentReturn",
	"Store Payment Recieve":         "StorePaymentRecieve",
}

var ItemType = map[string]string{
	"Raw Material":     "RawMaterial",
	"In House Product": "InHouseProduct",
	"Outside Product":  "OutsideProduct",
}

var Status = map[string]string{
	"Complete": "Complete",
	"OnGoing":  "OnGoing",
}
