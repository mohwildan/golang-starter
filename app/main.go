package main

import (
	"fmt"
	"golang-starter/service/database"
	"golang-starter/service/delivery/handler"
	"golang-starter/service/delivery/repository"
	"golang-starter/service/delivery/usecase/sample"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("environments/.env")
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())

	timeOut, _ := strconv.Atoi(os.Getenv("TIMEOUT"))
	timeOutDuration := time.Duration(timeOut) * time.Second

	//mongo
	mongoDB := database.InitMongo(timeOutDuration)

	//repository
	sampleRepo := repository.SampleRepo(mongoDB)
	//usecase
	sampleUc := sample.Usecase(sampleRepo, timeOutDuration)

	//handler
	handler.SampleHandler(router, sampleUc)
	port := os.Getenv("PORT")
	fmt.Printf("[%s] YTC API SAMPLE running on port: %s\n", time.Now().Format("2006-01-02 15:04:05"), port)
	router.Run(":" + port)
}
