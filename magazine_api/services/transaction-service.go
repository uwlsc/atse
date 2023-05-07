package services

import (
	"magazine_api/component"
	"magazine_api/constants"
	"magazine_api/lib"
	"magazine_api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//TransactionService service layer
type TransactionService struct {
	logger lib.Logger
	comp   component.TransactionComponent
}

//NewTransactionService creates new instance of TransactionService
func NewTransactionService(logger lib.Logger, comp component.TransactionComponent) TransactionService {
	return TransactionService{logger: logger, comp: comp}
}

// Creates the Transaction in database
func (t TransactionService) CreateTransaction(assign *models.Transaction) (*models.Transaction, error) {
	assign = t.BeforeCreate(assign)

	err := t.comp.CreateTransaction(*assign)
	if err != nil {
		return nil, err
	}

	return assign, nil
}

// Lists the Transaction in database
func (t TransactionService) ListsTransaction(c *gin.Context) ([]*models.Transaction, error) {
	assigns, err := t.comp.ListTransactions(c.MustGet(constants.Limit).(int64), c.MustGet(constants.Page).(int64))
	if err != nil {
		return nil, err
	}

	return assigns, nil
}

// Lists Deleted Transactions from database
func (t TransactionService) ListsDeletedTransactions(c *gin.Context) ([]*models.Transaction, error) {

	assigns, err := t.comp.ListDeletedTransactions(c.MustGet(constants.Limit).(int64), c.MustGet(constants.Page).(int64))

	if err != nil {
		return nil, err
	}

	return assigns, nil
}

// Get transaction  by id from database
func (t TransactionService) GetTransactionByID(id uuid.UUID) (*models.Transaction, error) {

	assigns, err := t.comp.GetTransactionFromID(id)

	if err != nil {
		return nil, err
	}

	return assigns, nil
}

//Get transaction by type
func (t TransactionService) GetTransactionByType(paid_type string, c *gin.Context) ([]*models.Transaction, error) {
	assigns, err := t.comp.GetTransactionFromType(paid_type, c.MustGet(constants.Limit).(int64), c.MustGet(constants.Page).(int64))

	if err != nil {
		return nil, err
	}

	return assigns, nil
}

//Get transaction by medium
func (t TransactionService) GetTransactionByMedium(paid_medium string, c *gin.Context) ([]*models.Transaction, error) {
	assigns, err := t.comp.GetTransactionFromMedium(paid_medium, c.MustGet(constants.Limit).(int64), c.MustGet(constants.Page).(int64))

	if err != nil {
		return nil, err
	}

	return assigns, nil
}

//Get transaction by date
func (t TransactionService) GetTransactionByDate(date time.Time, c *gin.Context) ([]*models.Transaction, error) {
	assigns, err := t.comp.GetTransactionFromDate(date, c.MustGet(constants.Limit).(int64), c.MustGet(constants.Page).(int64))

	if err != nil {
		return nil, err
	}

	return assigns, nil
}

//Get transaction by month
func (t TransactionService) GetTransactionByMonth(month string, c *gin.Context) ([]*models.Transaction, error) {
	assigns, err := t.comp.GetTransactionFromMonth(month, c.MustGet(constants.Limit).(int64), c.MustGet(constants.Page).(int64))

	if err != nil {
		return nil, err
	}

	return assigns, nil
}

//Get transaction from account id
func (t TransactionService) GetTransactionFromAccount(account uuid.UUID, c *gin.Context) ([]*models.Transaction, error) {
	assigns, err := t.comp.GetTransactionFromAccount(account, c.MustGet(constants.Limit).(int64), c.MustGet(constants.Page).(int64))

	if err != nil {
		return nil, err
	}

	return assigns, nil
}

//Get transaction from account id
func (t TransactionService) GetTransactionToAccount(account uuid.UUID, c *gin.Context) ([]*models.Transaction, error) {
	assigns, err := t.comp.GetTransactionToAccount(account, c.MustGet(constants.Limit).(int64), c.MustGet(constants.Page).(int64))

	if err != nil {
		return nil, err
	}

	return assigns, nil
}

// Update user by in our database
func (t TransactionService) UpdateTransaction(id uuid.UUID, patch *map[string]interface{}) error {
	err := t.comp.PatchTransaction(id, patch)

	if err != nil {
		return err
	}

	return nil
}

// Delete Tailor Assign by in our database
func (o TransactionService) DeleteTransaction(id uuid.UUID) error {
	err := o.comp.DeleteTransaction(id)
	if err != nil {
		return err
	}

	return nil
}

// Permanent Delete Tailor Assign by ID in our database permanently
func (t TransactionService) PermanentDeleteTransaction(id uuid.UUID) error {
	err := t.comp.PermanentDeleteTransaction(id)

	if err != nil {
		return err
	}

	return nil
}

func (t TransactionService) BeforeCreate(assign *models.Transaction) *models.Transaction {
	assign.ID = uuid.New()
	create := time.Now()
	assign.CreatedOn = &create
	assign.UpdatedOn = &create

	return assign
}
