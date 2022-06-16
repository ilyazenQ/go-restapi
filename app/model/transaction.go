package model

import (
	"errors"
	"github.com/shopspring/decimal"
	"payment_gateway/actions"
	"time"
)

const closedTransactionId = 5

var Sign = []byte("sign")

type Transaction struct {
	Id          uint64          `json:"id" db:"id"`
	User_id     int             `json:"user_id" binding:"required"`
	User_email  string          `json:"user_email" binding:"required"`
	Total       decimal.Decimal `json:"total" binding:"required"`
	Currency_id int             `json:"currency_id" binding:"required"`
	Status_id   int             `json:"status_id"`
	Created_at  time.Time       `json:"created_at"`
	Updated_at  time.Time       `json:"updated_at"`
}

type User struct {
	ForToken int `json:"for_token"`
}

type Status struct {
	Id    uint64 `json:"id" db:"id"`
	Title string `json:"title" binding:"required"`
}

type ChangeStatusRequest struct {
	Transaction_id uint64 `json:"transaction_id"`
	Status_id      uint64 `json:"status_id"`
}

//TODO Создать Response структуру

var CorrectUser = User{
	1,
}

// TODO разбить на подметоды и экшены

func CheckLogin(u User) (string, error) {
	if CorrectUser.ForToken != u.ForToken {
		errors.New("Field not correct")
		return "Error", errors.New("Field not correct")
	}

	return actions.GenerateToken(Sign)
}

func CreateTransaction(t *Transaction) (string, error) {
	t.Status_id = actions.GetNewOrErrorStatus()
	t.Created_at = time.Now()
	t.Updated_at = time.Now()

	query := `insert into transactions(user_id, user_email, total, currency_id, status_id, created_at, updated_at) 
				values($1, $2, $3, $4, $5, $6, $7) RETURNING id;`

	_, err := db.Exec(query, t.User_id, t.User_email, t.Total, t.Currency_id, t.Status_id, t.Created_at, t.Updated_at)

	if err != nil {
		return "fail", err
	}

	return "success", nil
}

func GetTransaction(id uint64) (Transaction, error) {
	var transaction Transaction

	query := `select id, user_id, user_email, total, currency_id, status_id, created_at, updated_at from transactions where id=$1`
	row, err := db.Query(query, id)
	if err != nil {
		return transaction, err
	}

	defer row.Close()

	if row.Next() {
		var userId int
		var userEmail string
		var total decimal.Decimal
		var currencyId int
		var statusId int
		var createdAt time.Time
		var updatedAt time.Time

		err := row.Scan(&id, &userId, &userEmail, &total, &currencyId, &statusId, &createdAt, &updatedAt)
		if err != nil {
			return transaction, err
		}

		transaction = Transaction{
			Id:          id,
			User_id:     userId,
			User_email:  userEmail,
			Total:       total,
			Currency_id: currencyId,
			Status_id:   statusId,
			Created_at:  createdAt,
			Updated_at:  updatedAt,
		}
	}

	return transaction, nil
}

func GetTransactionByUserId(userId int64) ([]Transaction, error) {
	var transactions []Transaction

	query := `select id, user_id, user_email, total, currency_id, status_id, created_at, updated_at from transactions where user_id=$1;`

	rows, err := db.Query(query, userId)
	if err != nil {
		return transactions, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint64
		var userId int
		var userEmail string
		var total decimal.Decimal
		var currencyId int
		var statusId int
		var createdAt time.Time
		var updatedAt time.Time

		err := rows.Scan(&id, &userId, &userEmail, &total, &currencyId, &statusId, &createdAt, &updatedAt)
		if err != nil {
			return transactions, err
		}

		transaction := Transaction{
			Id:          id,
			User_id:     userId,
			User_email:  userEmail,
			Total:       total,
			Currency_id: currencyId,
			Status_id:   statusId,
			Created_at:  createdAt,
			Updated_at:  updatedAt,
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func GetTransactionByUserEmail(email string) ([]Transaction, error) {
	var transactions []Transaction

	query := `select id, user_id, user_email, total, currency_id, status_id, created_at, updated_at from transactions 
                                                                                      where user_email=$1;`

	rows, err := db.Query(query, email)
	if err != nil {
		return transactions, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint64
		var userId int
		var userEmail string
		var total decimal.Decimal
		var currencyId int
		var statusId int
		var createdAt time.Time
		var updatedAt time.Time

		err := rows.Scan(&id, &userId, &userEmail, &total, &currencyId, &statusId, &createdAt, &updatedAt)
		if err != nil {
			return transactions, err
		}

		transaction := Transaction{
			Id:          id,
			User_id:     userId,
			User_email:  userEmail,
			Total:       total,
			Currency_id: currencyId,
			Status_id:   statusId,
			Created_at:  createdAt,
			Updated_at:  updatedAt,
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func GetTransactionStatus(id uint64) (Status, error) {
	var transaction Transaction
	var status Status
	transaction, err := GetTransaction(id)
	if err != nil {
		return status, err
	}

	statusId := transaction.Status_id

	query := `select title from status where id=$1`
	row, err := db.Query(query, statusId)

	defer row.Close()

	if row.Next() {
		var title string

		err := row.Scan(&title)
		if err != nil {
			return status, err
		}

		status = Status{
			Title: title,
		}
	}

	return status, nil
}

func GetAllTransactions() ([]Transaction, error) {
	var transactions []Transaction

	query := `select * from transactions;`

	rows, err := db.Query(query)
	if err != nil {
		return transactions, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint64
		var userId int
		var userEmail string
		var total decimal.Decimal
		var currencyId int
		var statusId int
		var createdAt time.Time
		var updatedAt time.Time

		err := rows.Scan(&id, &userId, &userEmail, &total, &currencyId, &statusId, &createdAt, &updatedAt)
		if err != nil {
			return transactions, err
		}

		transaction := Transaction{
			Id:          id,
			User_id:     userId,
			User_email:  userEmail,
			Total:       total,
			Currency_id: currencyId,
			Status_id:   statusId,
			Created_at:  createdAt,
			Updated_at:  updatedAt,
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func CloseTransaction(id uint64) (Transaction, error) {
	var transaction Transaction

	transaction, err := GetTransaction(id)

	if err != nil {
		return transaction, err
	}

	if _, ok := actions.AvailableStatus()[transaction.Status_id]; !ok {
		return transaction, errors.New("Cant close this transaction")
	}

	query := `update transactions set status_id=$1, updated_at=$2 where id=$3;`

	_, err = db.Exec(query, closedTransactionId, time.Now(), transaction.Id)
	if err != nil {
		return transaction, err
	}
	transaction.Status_id = closedTransactionId
	return transaction, nil
}

func ChangeTransactionStatus(csr *ChangeStatusRequest) (Transaction, error) {
	var transaction Transaction

	transaction, err := GetTransaction(csr.Transaction_id)

	if err != nil {
		return transaction, err
	}

	if int(csr.Status_id) == transaction.Status_id {
		return transaction, errors.New("Cant change. It current status now")
	}

	if _, ok := actions.GetAllStatus()[int(csr.Status_id)]; !ok {
		return transaction, errors.New("Status not found")
	}

	query := `update transactions set status_id=$1, updated_at=$2 where id=$3;`

	_, err = db.Exec(query, csr.Status_id, time.Now(), transaction.Id)
	if err != nil {
		return transaction, err
	}
	transaction.Status_id = int(csr.Status_id)
	return transaction, nil
}
