package component

import (
	"context"
	"errors"
	"magazine_api/infrastructure"
	"magazine_api/models"
	"time"

	"github.com/elgris/sqrl"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionComponent struct {
	infrastructure.Database
}

//NewTransactionComponent creates new transaction component
func NewTransactionComponent(db infrastructure.Database) TransactionComponent {
	return TransactionComponent{db}
}

// Creates the Transaction in our Database
func (t TransactionComponent) CreateTransaction(transaction models.Transaction) error {
	sql, args, err := sqrl.Insert("transactions").
		Columns("id", "title", "transaction_cost", "debit_amount", "credit_amount", "payment_to", "payment_from", "payment_date", "payment_month", "paid_type", "paid_medium", "bank_payment_from", "bank_payment_to", "bank_payment_transaction_id",
			"online_payment_name",
			"online_payment_from",
			"online_payment_to",
			"online_payment_transaction_id", "remarks",
			"created_on").
		Values(transaction.ID, transaction.Title, transaction.TransactionCost, transaction.DebitAmount, transaction.CreditAmount, transaction.PaymentTo, transaction.PaymentFrom, time.Now(), transaction.PaymentMonth, transaction.PaidType, transaction.PaidMedium, transaction.BankPaymentFrom, transaction.BankPaymentTo, transaction.BankPaymentTransactionId, transaction.OnlinePaymentName, transaction.OnlinePaymentFrom, transaction.OnlinePaymentTo, transaction.OnlinePaymentTransactionId, transaction.Remarks, transaction.CreatedOn).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return err
	}
	exec, err := t.Exec(context.Background(), sql, args[:]...)
	if err != nil {
		return err
	}

	if exec.RowsAffected() != 1 {
		return errors.New("not inserted")
	}

	return nil
}

//List deleted tailor assignments
func (t TransactionComponent) ListDeletedTransactions(limit int64, offset int64) ([]*models.Transaction, error) {
	var tailors []*models.Transaction
	sql, args, err := sqrl.Select("*").
		From("transactions"). // ?
		Where(sqrl.NotEq{"deleted_on": nil}).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), t, &tailors, sql, args...); err != nil {
		return nil, err
	}

	return tailors, nil
}

//Lists Transaction from database
func (t TransactionComponent) ListTransactions(limit int64, offset int64) ([]*models.Transaction, error) {
	var tailors []*models.Transaction
	sql, args, err := sqrl.Select("*").
		From("transactions").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), t, &tailors, sql, args...); err != nil {
		return nil, err
	}

	return tailors, nil
}

// Gets single transaction from Database using ID
func (t TransactionComponent) GetTransactionFromID(id uuid.UUID) (*models.Transaction, error) {
	var transaction models.Transaction
	sql, args, err := sqrl.Select("*").From("transactions").Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Get(context.Background(), t, &transaction, sql, args[:]...); err != nil {
		return nil, err
	}

	return &transaction, nil
}

// TODO
// Get transaction by Paid Type
// Get transaction by Paid Medium
// Get transaction by month
// Get all transaction performed from account id
// Get all transaction performed to account id
// Sum Debit and Credit for transaction by paid type (rewquirement 3)
// Get transaction by date (exact date, after date, before date all 3)
// Get Credit and debit sum for all seleceted transaction -- similar to requirement 3

// func (t TransactionComponent) SumDebitCreditByType(payment_type string) (*float32, error) {
// 	sql, args, err := sqrl.Select("SUM(credit)").
// 		FROM("transactions").
// 		Where(sqrl.Eq{"deleted_on": nil}).
// 		Where(sqrl.Eq{"payment_type": payment_type}).
// 		PlaceholderFormat(sqrl.Dollar).
// 		ToSql()

// 	return nil, nil
// }

//Get transaction from mediom
func (t TransactionComponent) GetTransactionBeforeDate(date time.Time, limit int64, offset int64) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	sql, args, err := sqrl.Select("*").
		From("transactions").
		Where(sqrl.Eq{"deleted_on": nil}).
		Where(sqrl.Expr("payment_date <= ?", date)).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), t, &transactions, sql, args...); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t TransactionComponent) GetTransactionAfterDate(date time.Time, limit int64, offset int64) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	sql, args, err := sqrl.Select("*").
		From("transactions").
		Where(sqrl.Eq{"deleted_on": nil}).
		Where(sqrl.Expr("payment_date >= ?", date)).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), t, &transactions, sql, args...); err != nil {
		return nil, err
	}

	return transactions, nil
}

//Get transaction from mediom
func (t TransactionComponent) GetTransactionFromDate(date time.Time, limit int64, offset int64) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	sql, args, err := sqrl.Select("*").
		From("transactions").
		Where(sqrl.Eq{"deleted_on": nil}).
		Where(sqrl.Eq{"payment_date": date}).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), t, &transactions, sql, args...); err != nil {
		return nil, err
	}

	return transactions, nil
}

//Get transaction from mediom
func (t TransactionComponent) GetTransactionFromMonth(month string, limit int64, offset int64) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	sql, args, err := sqrl.Select("*").
		From("transactions").
		Where(sqrl.Eq{"deleted_on": nil}).
		Where(sqrl.Eq{"payment_month": month}).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), t, &transactions, sql, args...); err != nil {
		return nil, err
	}

	return transactions, nil
}

//Get all transaction to account
func (t TransactionComponent) GetTransactionToAccount(account uuid.UUID, limit int64, offset int64) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	sql, args, err := sqrl.Select("*").
		From("transactions").
		Where(sqrl.Eq{"deleted_on": nil}).
		Where(sqrl.Eq{"payment_to": account}).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), t, &transactions, sql, args...); err != nil {
		return nil, err
	}

	return transactions, nil
}

//GET all transaction from account
func (t TransactionComponent) GetTransactionFromAccount(account uuid.UUID, limit int64, offset int64) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	sql, args, err := sqrl.Select("*").
		From("transactions").
		Where(sqrl.Eq{"deleted_on": nil}).
		Where(sqrl.Eq{"payment_from": account}).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), t, &transactions, sql, args...); err != nil {
		return nil, err
	}

	return transactions, nil
}

//Get transaction from mediom
func (t TransactionComponent) GetTransactionFromMedium(paid_medium string, limit int64, offset int64) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	sql, args, err := sqrl.Select("*").
		From("transactions").
		Where(sqrl.Eq{"deleted_on": nil}).
		Where(sqrl.Eq{"paid_medium": paid_medium}).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), t, &transactions, sql, args...); err != nil {
		return nil, err
	}

	return transactions, nil
}

//Get transaction from type
func (t TransactionComponent) GetTransactionFromType(paid_type string, limit int64, offset int64) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	sql, args, err := sqrl.Select("*").
		From("transactions").
		Where(sqrl.Eq{"deleted_on": nil}).
		Where(sqrl.Eq{"paid_type": paid_type}).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), t, &transactions, sql, args...); err != nil {
		return nil, err
	}

	return transactions, nil
}

// Updates transaction in our database
func (t TransactionComponent) PatchTransaction(id uuid.UUID, patch *map[string]interface{}) error {
	sql, args, err := sqrl.Update("transactions").SetMap(*patch).Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil
	}

	exec, err := t.Exec(context.Background(), sql, args[:]...)
	if err != nil {
		return err
	}

	if exec.RowsAffected() != 1 {
		return errors.New("not updated")
	}

	return nil
}

// Delete transaction in our database
func (t TransactionComponent) DeleteTransaction(id uuid.UUID) error {

	sql, arg, err := sqrl.Update("transactions").SetMap(gin.H{"deleted_on": time.Now()}).
		Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return err
	}

	exec, err := t.Exec(context.Background(), sql, arg[:]...)
	if err != nil {
		return err
	}

	count := exec.RowsAffected()
	if count != 1 {
		return err
	}

	return nil
}

// Permanent Delete Transaction permanently deletes the Taior Assing in our database
func (t TransactionComponent) PermanentDeleteTransaction(id uuid.UUID) error {
	sql, arg, err := sqrl.Delete("transactions").Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return err
	}

	exec, err := t.Exec(context.Background(), sql, arg[:]...)
	if err != nil {
		return err
	}

	count := exec.RowsAffected()
	if count != 1 {
		return err
	}

	return nil
}
