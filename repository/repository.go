package repository

import (
	"database/sql"
	model "fanitest/models"
	"fmt"
	"log"
	"strconv"
)

type InvestasiRepository interface {
	SaveTransaction(transaction model.Transaction) error
	CheckIfEmailExists(email string) (bool, error)
	GetLastTransactionNumber() (string, error)
	GetGenerateTransactionNumber(prefix string, lastTransactionNumber string) string
}

type investmentRepository struct {
	db *sql.DB
}

func NewInvestmentRepository(db *sql.DB) InvestasiRepository {
	return &investmentRepository{db: db}
}

func (r *investmentRepository) SaveTransaction(transaction model.Transaction) error {
	query := `INSERT INTO investasi (tanggal_transaksi, no_transaksi, nama, jenis_kelamin, usia, email, perokok, nominal, lama_investasi, periode_pembayaran, metode_pembayaran, total_bayar) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(transaction.TglTransaksi, transaction.NoTransaction, transaction.Nama, transaction.JenisKelamin, transaction.Usia, transaction.Email, transaction.Perokok, transaction.Nominal, transaction.LamaInvestasi, transaction.PeriodeBayar, transaction.MetodeBayar, transaction.TotalBayar)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func (r *investmentRepository) CheckIfEmailExists(email string) (bool, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM investasi WHERE email = ?", email).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

// GenerateTransactionNumber generates a transaction number with the given prefix
func GenerateTransactionNumber(prefix string, lastTransactionNumber string) string {
	//extract the last transaction number from the given string
	lastTransactionNumberInt, err := strconv.Atoi(lastTransactionNumber[len(prefix):])
	if err != nil {
		log.Fatal(err)
	}

	//increment the last transaction number by 1
	lastTransactionNumberInt++

	//return the new transaction number with the same prefix and the incremented number
	return fmt.Sprintf("%s%06d", prefix, lastTransactionNumberInt)
}

func (r *investmentRepository) GetLastTransactionNumber() (string, error) {
	var transactionNumber string
	err := r.db.QueryRow("SELECT no_transaksi FROM investasi ORDER BY id DESC LIMIT 1").Scan(&transactionNumber)
	// if err == sql.ErrNoRows {
	// 	return "TRX000001", nil
	// }
	if err != nil {
		return "", err
	}
	return transactionNumber, nil
}

// GenerateTransactionNumber generates a transaction number with the given prefix
func (r *investmentRepository) GetGenerateTransactionNumber(prefix string, lastTransactionNumber string) string {
	//extract the last transaction number from the given string
	lastTransactionNumberInt, err := strconv.Atoi(lastTransactionNumber[len(prefix):])
	if err != nil {
		log.Fatal(err)
	}

	//increment the last transaction number by 1
	lastTransactionNumberInt++

	//return the new transaction number with the same prefix and the incremented number
	return fmt.Sprintf("%s%06d", prefix, lastTransactionNumberInt)
}
