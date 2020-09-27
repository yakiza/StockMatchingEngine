package service

import "StockMatchingEngine/models"

type ServiceTrade struct {
	trade *models.Trade
}

func (st *ServiceTrade) Process(db *DatabaseService) {

}
