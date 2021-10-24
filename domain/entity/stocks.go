package entity

type StockBalanceOffer struct {
	OfferID     string `json:"offer_id"`
	ProductID   int    `json:"product_id"`
	Stock       int    `json:"stock"`
	WarehouseID int    `json:"warehouse_id"`
}

type StockBalanceRequest struct {
	Stocks []StockBalanceOffer `json:"stocks"`
}

type StockBalanceResponse struct {
	Result []struct {
		Errors []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"errors"`
		OfferID     string `json:"offer_id"`
		ProductID   int    `json:"product_id"`
		Updated     bool   `json:"updated"`
		WarehouseID int    `json:"warehouse_id"`
	} `json:"result"`
}
