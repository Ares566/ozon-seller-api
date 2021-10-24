package main

import (
	"github.com/robfig/cron/v3"
	"ozon-seller-api/application"
	"ozon-seller-api/infrastructure/logger"
	"ozon-seller-api/internal/app"
)

func main() {

	var (
		appData = app.New()
		fbsApp  = appData[app.FBSApp].(*application.FBSApplication)
	)

	c := cron.New()

	// раз в 3 минуты проверяем новые заказы
	_, err := c.AddFunc("@every 1m", func() { fbsApp.GetUnfulfilledList() })
	if err != nil {
		logger.Error("Ошибка запуска забора заказов по расписанию")
		return
	}

	//err = c.AddFunc("0 */3 * * * *", func() { fbsApp.CheckStockBalances() })
	//if err != nil {
	//	logger.Error("Ошибка запуска обновления остатков по расписанию")
	//	return
	//}

	c.Start()

}
