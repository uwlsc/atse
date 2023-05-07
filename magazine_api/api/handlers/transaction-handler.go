package handlers

import (
	"magazine_api/lib"
	"magazine_api/models"
	"magazine_api/services"
	"time"

	"github.com/danhper/structomap"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionHandler struct {
	logger  lib.Logger
	service services.TransactionService
}

func NewTransactionHandler(logger lib.Logger, service services.TransactionService) TransactionHandler {
	return TransactionHandler{logger: logger, service: service}
}

//CreateTransaction godoc
// @Summary      Create Transaction
// @Description  It creates assign for cutting.
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Param        transaction  body      models.TransactionBase  true  "Add Transaction"
// @Success      200          {object}  object{data=models.Transaction}
// @Router       /transaction [post]
//
//Creates Transaction
func (t TransactionHandler) CreateTransaction(c *gin.Context) {
	var assign *models.Transaction
	if err := c.ShouldBind(&assign); err != nil {
		handleError(t.logger, c, err)
		return
	}

	// Create Transaction in our Database
	assign, err := t.service.CreateTransaction(assign)
	if err != nil {
		handleError(t.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": assign})
}

// ListTransaction godoc
// @Summary      Lists Transaction
// @Description  List assigns for cutting.
// @Tags         Transaction
// @Produce      json
// @Success      200      {object}  object{data=[]models.Transaction}
// @Router       /transaction [get]
//
// List Transaction controller
func (t TransactionHandler) ListTransaction(c *gin.Context) {
	assigns, err := t.service.ListsTransaction(c)
	if err != nil {
		handleError(t.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": assigns})
}

// ListDeletedCuttingassign godoc
// @Summary      Lists Transaction
// @Description  List assigns for cutting.
// @Tags         Transaction
// @Produce      json
// @Success      200      {object}  object{data=[]models.Transaction}
// @Router       /transaction/deleted [get]
//
// List Transaction controller
func (t TransactionHandler) ListDeletedTransaction(c *gin.Context) {
	assigns, err := t.service.ListsDeletedTransactions(c)
	if err != nil {
		handleError(t.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": assigns})
}

// GetTransactionByID godoc
// @Summary      Gets One Transaction by ID
// @Description  Gets One Transactions by ID
// @Tags         Transaction
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=models.Transaction}
// @Router       /transaction/id/{id} [get]
// Gets Transaction By ID controller
func (t TransactionHandler) GetTransactionByID(c *gin.Context) {
	id := c.Param("id")

	assign, err := t.service.GetTransactionByID(uuid.MustParse(id))
	if err != nil {
		handleError(t.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": assign})
}

// GetTransactionByType godoc
// @Summary      Gets Transaction by Type
// @Description  Gets all Transactions by Type
// @Tags         Transaction
// @Produce      json
// @Param        paid_type  path      string  true  "Type"
// @Success      200        {object}  object{data=[]models.Transaction}
// @Router       /transaction/type/{paid_type} [get]
// Gets Transaction By Paid type controller
func (t TransactionHandler) GetTransactionByType(c *gin.Context) {
	paid_type := c.Param("paid_type")

	assign, err := t.service.GetTransactionByType(paid_type, c)
	if err != nil {
		handleError(t.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": assign})
}

// GetTransactionByMedium godoc
// @Summary      Gets Transactions by Medium
// @Description  Gets All Transactions by Paid Medium
// @Tags         Transaction
// @Produce      json
// @Param        paid_medium  path      string  true  "Type"
// @Success      200          {object}  object{data=[]models.Transaction}
// @Router       /transaction/medium/{paid_medium} [get]
// Gets Transaction By Paid medium controller
func (t TransactionHandler) GetTransactionByMedium(c *gin.Context) {
	paid_medium := c.Param("paid_medium")

	assign, err := t.service.GetTransactionByMedium(paid_medium, c)
	if err != nil {
		handleError(t.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": assign})
}

// GetTransactionByDate godoc
// @Summary      Gets Transactions by date
// @Description  Gets All Transactions by exact date
// @Tags         Transaction
// @Produce      json
// @Param        date  path      string  true  "Date"
// @Success      200   {object}  object{data=[]models.Transaction}
// @Router       /transaction/date/{date} [get]
// Gets Transaction By date controller
func (t TransactionHandler) GetTransactionByDate(c *gin.Context) {
	date := c.Param("date")

	assign, err := t.service.GetTransactionByMedium(date, c)
	if err != nil {
		handleError(t.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": assign})
}

// GetTransactionByMonth godoc
// @Summary      Gets Transactions by month
// @Description  Gets All Transactions by month
// @Tags         Transaction
// @Produce      json
// @Param        month  path      string  true  "Month"
// @Success      200    {object}  object{data=[]models.Transaction}
// @Router       /transaction/month/{month} [get]
// Gets Transaction By month controller
func (t TransactionHandler) GetTransactionByMonth(c *gin.Context) {
	month := c.Param("month")

	assign, err := t.service.GetTransactionByMedium(month, c)
	if err != nil {
		handleError(t.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": assign})
}

// GetTransactionFromAccount godoc
// @Summary      Gets Transactions by Account.
// @Description  Gets All Transactions done from Account.
// @Tags         Transaction
// @Produce      json
// @Param        account  path      string  true  "ID"
// @Success      200  {object}  object{data=[]models.Transaction}
// @Router       /transaction/fromaccount/{account} [get]
// Gets Transaction From Account controller
func (t TransactionHandler) GetTransactionFromAccount(c *gin.Context) {
	account := c.Param("account")

	assign, err := t.service.GetTransactionFromAccount(uuid.MustParse(account), c)
	if err != nil {
		handleError(t.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": assign})
}

// GetTransactionToAccount godoc
// @Summary      Gets Transactions to Account.
// @Description  Gets All Transactions done to the Account.
// @Tags         Transaction
// @Produce      json
// @Param        account  path      string  true  "ID"
// @Success      200  {object}  object{data=[]models.Transaction}
// @Router       /transaction/toaccount/{account} [get]
// Gets Transaction to the Account controller
func (t TransactionHandler) GetTransactionToAccount(c *gin.Context) {
	account := c.Param("account")

	assign, err := t.service.GetTransactionToAccount(uuid.MustParse(account), c)
	if err != nil {
		handleError(t.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": assign})
}

//UpdateTransaction godoc
// @Summary      Update Transaction
// @Description  Updates Transaction
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Update Transaction"
// @Success      200  {object}  object{data=models.Transaction}
// @Router       /transaction/{id} [patch]
// Patch Transaction by Id controller
func (t TransactionHandler) PatchTransaction(c *gin.Context) {
	id := c.Param("id")

	assign, err := t.service.GetTransactionByID(uuid.MustParse(id))
	if err != nil {
		handleError(t.logger, c, err)
		return
	}

	var newassign models.Transaction
	if err := c.ShouldBindJSON(&newassign); err != nil {
		handleError(t.logger, c, err)
		return
	}

	assignMap := structomap.New().UseSnakeCase().PickAll().
		OmitIf(func(ch interface{}) bool {
			return newassign.Title == nil
		}, "Title").
		OmitIf(func(ch interface{}) bool {
			return newassign.TransactionCost == nil
		}, "TransactionCost").
		OmitIf(func(ch interface{}) bool {
			return newassign.DebitAmount == nil
		}, "DebitAmount").
		OmitIf(func(ch interface{}) bool {
			return newassign.CreditAmount == nil
		}, "CreditAmount").
		OmitIf(func(ch interface{}) bool {
			return newassign.PaymentTo == nil
		}, "PaymentTo").
		OmitIf(func(ch interface{}) bool {
			return newassign.PaymentMonth == nil
		}, "PaymentMonth").
		OmitIf(func(ch interface{}) bool {
			return newassign.PaidType == nil
		}, "PaidType").
		OmitIf(func(ch interface{}) bool {
			return newassign.PaidMedium == nil
		}, "PaidMedium").
		OmitIf(func(ch interface{}) bool {
			return newassign.BankPaymentFrom == nil
		}, "ProductRecivedTime").
		OmitIf(func(ch interface{}) bool {
			return newassign.BankPaymentTo == nil
		}, "BankPaymentTo").
		OmitIf(func(ch interface{}) bool {
			return newassign.BankPaymentTransactionId == nil
		}, "BankPaymentTransactionId").
		OmitIf(func(ch interface{}) bool {
			return newassign.OnlinePaymentName == nil
		}, "OnlinePaymentName").
		OmitIf(func(ch interface{}) bool {
			return newassign.OnlinePaymentFrom == nil
		}, "OnlinePaymentFrom").
		OmitIf(func(ch interface{}) bool {
			return newassign.OnlinePaymentTo == nil
		}, "OnlinePaymentTo").
		OmitIf(func(ch interface{}) bool {
			return newassign.OnlinePaymentTransactionId == nil
		}, "OnlinePaymentTransactionId").
		OmitIf(func(ch interface{}) bool {
			return newassign.PaymentFrom == nil
		}, "PaymentFrom").
		OmitIf(func(ch interface{}) bool {
			return newassign.Remarks == nil
		}, "Remarks").
		Transform(newassign)

	if len(assignMap) > 0 {
		assignMap["updated_on"] = time.Now()
		assignMap["id"] = assign.ID

		err := t.service.UpdateTransaction(assign.ID, &assignMap)
		if err != nil {
			handleError(t.logger, c, err)
			return
		}

		c.JSON(200, gin.H{"data": assignMap})
		return
	}

	c.JSON(200, gin.H{"data": "nothing to update"})
}

// DeleteCategory godoc
// @Summary      Soft Delete Transaction
// @Description  Delete by Transaction ID
// @Tags         Transaction
// @Produce      json
// @Param        id   path      string  true  "assign Cutting ID"
// @Success      204  {object}  object{data=string}
// @Router       /transaction/{id} [delete]
//
// Delete Transaction By ID controller
func (t TransactionHandler) DeleteTransactionByID(c *gin.Context) {
	id := c.Param("id")

	err := t.service.DeleteTransaction(uuid.MustParse(id))
	if err != nil {
		handleError(t.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "successfully deleted"})
}

// PermanentDeleteassignCategory godoc
// @Summary      Permenant Delete an Transaction
// @Description  Delete by Transaction ID
// @Tags         Transaction
// @Produce      json
// @Param        id   path      string  true  "Transaction ID"
// @Success      204  {object}  object{data=string}
// @Router       /transaction/forcedelete/{id} [delete]
// Delete Transaction By ID controller
func (t TransactionHandler) PermanentDeleteTransactionByID(c *gin.Context) {
	id := c.Param("id")

	err := t.service.PermanentDeleteTransaction(uuid.MustParse(id))
	if err != nil {
		handleError(t.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "successfully deleted"})
}
