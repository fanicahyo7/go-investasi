package service

import (
	"errors"
	model "fanitest/models"
	"fanitest/repository"
	"strings"
)

type InvestasiService interface {
	CheckIfEmailExists(email string) (bool, error)
	GetLastTransactionNumber() (string, error)
	SaveTransaction(transaction model.Transaction) error
	CalculateTotalPayment(nominal int, investmentPeriod int, paymentPeriod string) int
	GenerateTransactionNumber(prefix string, lastTransactionNumber string) string
}

type investasiService struct {
	repo repository.InvestasiRepository
}

func NewInvestasiService(repo repository.InvestasiRepository) InvestasiService {
	return &investasiService{repo: repo}
}

func PerhitunganInvestasi(request model.PerhitunganInvetasiRequest) (map[int]map[string]float64, error) {

	if request.Usia == 0 {
		return nil, errors.New("masukkan Umur Terlebih dahulu")
	} else if request.JenisKelamin == "" || (strings.ToLower(request.JenisKelamin) != "wanita" && strings.ToLower(request.JenisKelamin) != "pria") {
		return nil, errors.New("masukkan jenis kelamin dengan benar (wanita/pria)")
	} else if request.Perokok == "" || (strings.ToLower(request.Perokok) != "ya" && strings.ToLower(request.Perokok) != "tidak") {
		return nil, errors.New("masukkan perokok dengan benar (ya/tidak)")
	}

	//initialize base investment rate
	persenTambahanInves := 0.0
	if strings.ToLower(request.JenisKelamin) == "pria" {
		if strings.ToLower(request.Perokok) == "ya" {
			persenTambahanInves = 0.01
		} else {
			persenTambahanInves = 0.02
		}
	} else {
		if strings.ToLower(request.Perokok) == "ya" {
			persenTambahanInves = 0.02
		} else {
			persenTambahanInves = 0.03
		}
	}

	//add additional investment rate based on age
	if request.Usia >= 0 && request.Usia <= 30 {
		persenTambahanInves += 0.01
	} else if request.Usia >= 31 && request.Usia <= 50 {
		persenTambahanInves += 0.005
	}

	//initialize map to store investment data for each year
	investasiData := make(map[int]map[string]float64)

	//calculate investment data for each year
	for i := 1; i <= request.LamaInvestasi; i++ {
		investasiData[i] = make(map[string]float64)
		investasiData[i]["awal"] = request.Nominal
		investasiData[i]["bunga"] = request.Nominal * persenTambahanInves
		investasiData[i]["akhir"] = request.Nominal + investasiData[i]["bunga"]
		request.Nominal = investasiData[i]["akhir"]
	}

	return investasiData, nil
}

func (s *investasiService) SaveTransaction(transaction model.Transaction) error {
	return s.repo.SaveTransaction(transaction)
}

// CheckIfEmailExists checks if the given email already exists in the database
func (s *investasiService) CheckIfEmailExists(email string) (bool, error) {
	return s.repo.CheckIfEmailExists(email)
}

// GetLastTransactionNumber gets the last transaction number from the database
func (s *investasiService) GetLastTransactionNumber() (string, error) {
	return s.repo.GetLastTransactionNumber()
}

func (s *investasiService) GenerateTransactionNumber(prefix string, lastTransactionNumber string) string {
	return s.repo.GetGenerateTransactionNumber(prefix, lastTransactionNumber)
}

func (s *investasiService) CalculateTotalPayment(nominal int, investmentPeriod int, paymentPeriod string) int {
	if paymentPeriod == "tahunan" {
		investmentPeriod--
	}

	return nominal / 12 * investmentPeriod
}