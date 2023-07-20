package main

import (
	"fmt"
	"golang-starter/generators"
	"golang-starter/service/database"
	"golang-starter/service/delivery/handler"
	"golang-starter/service/delivery/repository"
	"golang-starter/service/delivery/usecase/sample"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("environments/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	args := os.Args
	if len(args) > 0 {
		for _, arg := range args[1:] {
			if arg == "ns" {
				generators.GenerateService()
			}
		}
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())

	timeOutStr := os.Getenv("TIMEOUT")
	timeOut, err := strconv.Atoi(timeOutStr)
	if err != nil {
		log.Fatal("Invalid TIMEOUT value")
	}
	timeOutDuration := time.Duration(timeOut) * time.Second

	// Inisialisasi database
	mongoDB := database.InitMongo(timeOutDuration)

	// Inisialisasi repository
	repository := repository.NewRepository(mongoDB)
	// Inisialisasi usecase
	sampleUc := sample.Usecase(timeOutDuration, repository)

	// Inisialisasi handler
	sampleHandler := handler.NewSampleHandler(sampleUc)

	// Register route
	sampleHandler.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("[%s] YTC API SAMPLE running on port: %s\n", time.Now().Format("2006-01-02 15:04:05"), port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
