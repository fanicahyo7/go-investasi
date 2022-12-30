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

	// Enable some middleware
	app.Use(recover.New())
	app.Use(logger.New())

	//connect to database
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_inves")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//init repository
	investmentRepo := repository.NewInvestmentRepository(db)

	//init service
	investmentService := service.NewInvestasiService(investmentRepo)

	//init handler
	investmentHandler := handler.NewInvestmentHandler(investmentService)

	// Add route to handler
	app.Post("/soal-satu", handlers.PerhitunganInvestasi)
	app.Post("/transaction", investmentHandler.SaveTransaction)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
