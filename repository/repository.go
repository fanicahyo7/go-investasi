package repository

import (
	"database/sql"
	model "fanitest/models"
	"fmt"
)

type InvestasiRepository interface {
	SaveTransaction(transaction model.Transaction) error
	CheckIfEmailExists(email string) (bool, error)
	GetLastTransactionNumber() (string, error)
	GetAllInvestment() ([]model.InvestasiOutputAll, error)
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

func (r *investmentRepository) GetLastTransactionNumber() (string, error) {
	var transactionNumber string
	err := r.db.QueryRow("SELECT no_transaksi FROM investasi ORDER BY id DESC LIMIT 1").Scan(&transactionNumber)
	if err == sql.ErrNoRows {
		return "TRX000000", nil
	}
	if err != nil {
		println(err.Error())
		return "", err
	}
	return transactionNumber, nil
}

func (r *investmentRepository) GetAllInvestment() ([]model.InvestasiOutputAll, error) {
	var result []model.InvestasiOutputAll
	query := `SELECT ID,tanggal_transaksi,no_transaksi ,nama,jenis_kelamin,usia,email,perokok,nominal,lama_investasi,periode_pembayaran,metode_pembayaran,total_bayar FROM investasi`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var investasi model.InvestasiOutputAll
		err := rows.Scan(
			&investasi.ID,
			&investasi.TglTransaksi,
			&investasi.NoTransaction,
			&investasi.Nama,
			&investasi.JenisKelamin,
			&investasi.Usia,
			&investasi.Email,
			&investasi.Perokok,
			&investasi.Nominal,
			&investasi.LamaInvestasi,
			&investasi.PeriodeBayar,
			&investasi.MetodeBayar,
			&investasi.TotalBayar,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, investasi)
	}
	return result, nil
}

func (r *investmentRepository) GetInvestmentByTransactionNumber(transactionNumber string) (model.InvestasiOutputAll, error) {
	var investasi model.InvestasiOutputAll
	err := r.db.QueryRow("SELECT * FROM investasi WHERE no_transaksi=?", transactionNumber).Scan(
		&investasi.ID,
		&investasi.TglTransaksi,
		&investasi.NoTransaction,
		&investasi.Nama,
		&investasi.JenisKelamin,
		&investasi.Usia,
		&investasi.Email,
		&investasi.Perokok,
		&investasi.Nominal,
		&investasi.LamaInvestasi,
		&investasi.PeriodeBayar,
		&investasi.MetodeBayar,
		&investasi.TotalBayar,
		&investasi.Status,
	)
	if err != nil {
		return investasi, err
	}
	return investasi, nil
}
