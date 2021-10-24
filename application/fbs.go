package application

import (
	"ozon-seller-api/domain/entity"
	"ozon-seller-api/domain/repository"
	"ozon-seller-api/infrastructure/queue"
)

type FBSApplication struct {
	FBSRepository repository.FBSRepository
	queueConsumer *queue.RMQConsumer
}

func NewFBSApplication(r repository.FBSRepository, q *queue.RMQConsumer) *FBSApplication {
	if r == nil {
		return nil
	}

	return &FBSApplication{r, q}
}

// GetUnfulfilledList is
func (a *FBSApplication) GetUnfulfilledList() {
	// TODO на подумать: может метод только возвращать список будет, а другим методом будем отправлять
	a.FBSRepository.GetUnfulfilledList()
}

func (a *FBSApplication) CheckStockBalances() {

	// TODO взять из очереди RabbitMQ все изменения
	var sbl []entity.StockBalanceOffer
	//sbl, err := a.queueConsumer.Consume("packbalances")
	a.FBSRepository.SetStockBalances(sbl)
}
