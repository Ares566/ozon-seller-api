package api

import (
	"database/sql"
	"encoding/json"
	"ozon-seller-api/domain/entity"
	"ozon-seller-api/infrastructure/logger"
	"ozon-seller-api/infrastructure/net"
)

type FBSRepo struct {
	client *net.OzonClient
	db     *sql.DB
}

func NewFBSRepository(client *net.OzonClient, dbc *sql.DB) *FBSRepo {
	if client == nil || dbc == nil {
		return nil
	}

	return &FBSRepo{client, dbc}
}

func (c *FBSRepo) SetStockBalances(sbl []entity.StockBalanceOffer) {

	var chunkSize = 99
	var chunks [][]entity.StockBalanceOffer

	for i := 0; i < len(sbl); i += chunkSize {
		end := i + chunkSize

		if end > len(sbl) {
			end = len(sbl)
		}

		chunks = append(chunks, sbl[i:end])
	}

	for _, offers := range chunks {
		request := entity.StockBalanceRequest{
			Stocks: offers,
		}

		c.client.SendBalance(&request)
	}
}

func (c *FBSRepo) GetUnfulfilledList() {

	list, err := c.client.GetUnfulfilledList()
	if err != nil {
		logger.Error(err)
		return
	}

	stmt, err := c.db.Prepare("INSERT INTO 1c_queue(posting) VALUE (:posting)")
	if err != nil {
		logger.Error(err)
		return
	}
	defer stmt.Close()

	for _, posting := range list {
		bytes, jmErr := json.Marshal(posting)
		if jmErr != nil {
			logger.Error(jmErr)
		} else {

			if _, stmtErr := stmt.Exec(string(bytes)); err != nil {
				logger.Error(stmtErr)
			}
		}

	}
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
