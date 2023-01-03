package main

import (
	"database/sql"
	"fanitest/handlers"
	handler "fanitest/handlers"
	"fanitest/repository"
	service "fanitest/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_inves")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	investmentRepo := repository.NewInvestmentRepository(db)

	investmentService := service.NewInvestasiService(investmentRepo)

	investmentHandler := handler.NewInvestmentHandler(investmentService)

	app.Post("/soal-satu", handlers.PerhitunganInvestasi)
	app.Post("/soal-dua", investmentHandler.SaveTransaction)
	app.Get("/soal-tiga", investmentHandler.GetInvestasiData)

	log.Fatal(app.Listen(":3000"))
}
