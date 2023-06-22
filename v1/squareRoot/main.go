package main

import (
	"context"
	"os"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/clients"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/constants"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/repositories"
	"github.com/JavierFVasquez/truenorth-calculator-backend/v1/squareRoot/controllers"
	"github.com/JavierFVasquez/truenorth-calculator-backend/v1/squareRoot/services"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	ctx := context.Background()

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// Env
	mongoUri := getEnv("MONGO_URI", "")
	dbName := getEnv("DB_NAME", "")

	// Clients
	mongoClient, mongoClose := clients.NewMongoClient(ctx, mongoUri, dbName, &logger)
	defer mongoClose()

	recordRepository := repositories.NewRecordRepository(mongoClient, string(constants.RecordCollection), &logger)
	operationRepository := repositories.NewOperationRepository(mongoClient, string(constants.OperationCollection), &logger)
	squareRootService := services.NewSquareRootService(recordRepository, operationRepository, &logger)

	squareRootController := controllers.NewSquareRootController(squareRootService)

	lambda.Start(squareRootController.SquareRootController)
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
