package app

import (
	"ozon-seller-api/application"
	"ozon-seller-api/infrastructure/api"
	"ozon-seller-api/infrastructure/database"
	"ozon-seller-api/infrastructure/logger"
	"ozon-seller-api/infrastructure/net"
	"ozon-seller-api/internal/config"


	"github.com/joho/godotenv"
)


const (
	FBSApp      = iota
)

// New is
func New() map[int]interface{} {

	//load environment variable
	if err := godotenv.Load(); err != nil {
		logger.Error(err)
	}

	var (
		appConfig = config.NewAppConfig()
		dbConfig  = config.NewDatabaseConfig()
	)
	// connect to database
	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		logger.Error(err)
	}

	client, err := net.NewFBSClient(appConfig.ClientID, appConfig.Key)
	if err != nil {
		logger.Error(err)
	}
	// dependency injection
	// TODO может uber-go/dig ?
	var (
		fbsRepository = api.NewFBSRepository(client, db.Conn)
		fbsApp          = application.NewFBSApplication(fbsRepository)
	)

	return map[int]interface{}{
		FBSApp:    fbsApp,
	}
}
