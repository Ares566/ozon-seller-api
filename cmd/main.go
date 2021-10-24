package main

import (
	"ozon-seller-api/application"
	"ozon-seller-api/infrastructure/logger"
	"ozon-seller-api/internal/app"
	"github.com/robfig/cron"

)


func main() {

	var (
		appData   = app.New()
		fbsApp     = appData[app.FBSApp].(*application.FBSApplication)
	)


	c := cron.New()
	// раз в 3 минуты
	cacheError := c.AddFunc("0 */3 * * * *", func() { fbsApp.GetUnfulfilledList() })
	if cacheError != nil {
		logger.Error("Ошибка запуска обновления по расписанию")
		return
	}
	c.Start()


}
