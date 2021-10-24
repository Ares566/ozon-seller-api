package main

import (
	"ozon-seller-api/application"
	"ozon-seller-api/internal/app"
)


func main() {

	var (
		appData   = app.New()
		fbsApp     = appData[app.FBSApp].(*application.FBSApplication)
	)
	fbsApp.GetUnfulfilledList()

	//c := cron.New()
	//// раз в 3 минуты
	//cacheError := c.AddFunc("0 */3 * * * *", func() { fbsApp.GetUnfulfilledList() })
	//if cacheError != nil {
	//	logger.Error("Ошибка запуска обновления по расписанию")
	//	return
	//}
	//c.Start()


}
