package handlers

import (
	model "fanitest/models"
	service "fanitest/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type investasiHandler struct {
	investasiService service.InvestasiService
}

func NewInvestmentHandler(investasiService service.InvestasiService) *investasiHandler {
	return &investasiHandler{investasiService}
}

func PerhitunganInvestasi(c *fiber.Ctx) error {
	var request model.PerhitunganInvetasiRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"message": err,
			"status":  http.StatusBadGateway,
			"data":    nil,
		})
	}

	response, err := service.PerhitunganInvestasi(request)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
			"data":    response,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"data":    response,
	})
}

func (h *investasiHandler) SaveTransaction(c *fiber.Ctx) error {
	var input model.InvestasiInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  http.StatusBadRequest,
		})

	}

	exists, err := h.investasiService.CheckIfEmailExists(input.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal cek email",
			"status":  http.StatusInternalServerError,
		})

	}
	if exists {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Email Sudah ada",
			"status":  http.StatusBadRequest,
		})

	}

	transactionNumber, err := h.investasiService.GetLastTransactionNumber()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal generate no transaksi",
			"status":  http.StatusInternalServerError,
		})

	}
	transactionNumber = service.GenerateTransactionNumber("TRX", transactionNumber)

	totalBayar := service.CalculateTotalPayment(input.Nominal, input.LamaInvestasi, input.MetodeBayar)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to calculate total payment",
			"status":  http.StatusInternalServerError,
		})

	}

	investasi := model.Transaction{
		TglTransaksi:  time.Now(),
		NoTransaction: transactionNumber,
		Nama:          input.Nama,
		JenisKelamin:  input.JenisKelamin,
		Usia:          input.Usia,
		Email:         input.Email,
		Perokok:       input.Perokok,
		Nominal:       input.Nominal,
		LamaInvestasi: input.LamaInvestasi,
		PeriodeBayar:  input.PeriodeBayar,
		MetodeBayar:   input.MetodeBayar,
		TotalBayar:    totalBayar,
	}

	if err := h.investasiService.SaveTransaction(investasi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal simpan transaksi",
			"status":  http.StatusInternalServerError,
		})

	}

	investasiOutput := model.InvestasiOutput{
		NoTransaction: investasi.NoTransaction,
		TglTransaksi:  time.Now().Format("2006-01-02 15:04:05"),
		Nama:          investasi.Nama,
		JenisKelamin:  investasi.JenisKelamin,
		Usia:          strconv.Itoa(investasi.Usia) + " tahun",
		Nominal:       "Rp. " + strconv.Itoa(investasi.Nominal),
		LamaInvestasi: strconv.Itoa(investasi.LamaInvestasi) + " tahun",
		PeriodeBayar:  investasi.PeriodeBayar,
		MetodeBayar:   investasi.MetodeBayar,
		TotalBayar:    "Rp. " + strconv.Itoa(investasi.TotalBayar),
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success save data",
		"status":  http.StatusOK,
		"data":    investasiOutput,
	})
}

func (h *investasiHandler) GetInvestasiData(c *fiber.Ctx) error {
	investmentData, err := h.investasiService.GetInvestasiData()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Success fetching investment data",
		"data":    investmentData,
	})
}
