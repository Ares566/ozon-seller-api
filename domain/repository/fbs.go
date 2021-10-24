package repository

import "ozon-seller-api/domain/entity"

type FBSRepository interface {
	GetUnfulfilledList()
	SetStockBalances(sbl []entity.StockBalanceOffer)
}
