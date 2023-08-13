package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"golang-starter/service/database"
	"golang-starter/service/delivery/handler"
	"golang-starter/service/delivery/repository"
	"golang-starter/service/delivery/usecase/sample"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("environments/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := chi.NewRouter()

	timeOutStr := os.Getenv("TIMEOUT")
	timeOut, err := strconv.Atoi(timeOutStr)
	if err != nil {
		log.Fatal("Invalid TIMEOUT value")
	}
	timeOutDuration := time.Duration(timeOut) * time.Second

	// Inisialisasi database
	mongoDB := database.InitMongo(timeOutDuration)

	// Inisialisasi newRepository
	newRepository := repository.NewRepository(mongoDB)
	// Inisialisasi usecase
	sampleUc := sample.NewUsecase(timeOutDuration, newRepository)

	// Initialising handler
	sampleHandler := handler.NewSampleHandler(sampleUc)

	// Register route
	sampleHandler.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("[%s] YTC API SAMPLE running on port: %s\n", time.Now().Format("2006-01-02 15:04:05"), port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
