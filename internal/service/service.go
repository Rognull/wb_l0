package service

import (
	"errors"
 
	// "fmt"
	"l0/internal/db"
	"l0/internal/model"
	"github.com/sirupsen/logrus"
)

type OrderService struct{
	Storage db.OrderStorage
	cache  map[string]model.Order
	
}

func (s *OrderService) CreateOrder(newOrder model.Order) error{
	_ , ok := s.cache[newOrder.OrderUid]
if ok {
	logrus.Println("value exist")
	logrus.Println(s.cache)
    return errors.New("value exist in cashe")
}
	s.cache[newOrder.OrderUid] = newOrder
	err := s.Storage.AddOrder(newOrder)
	if err != nil {
		logrus.Println(err.Error())
	}
	 
	return err
}

func (s *OrderService) GetOrder(orderUid string) (model.Order,error){      // TODO очищение кэша и запрос в БД при ненахождении в кэше
	 result,ok:=s.cache[orderUid] // забираем значения по Uid из кэша, не из БД
	 if ok != true {
		return result, errors.New("cant find uid")
	 }
	 return result,nil
}


func NewOrderService(storage *db.OrderStorage,orders []model.Order) *OrderService{
	
	resultService := new(OrderService)
	resultService.Storage = *storage
	 resultService.cache =   make(map[string]model.Order)
	for _, order := range orders {
		resultService.cache[order.OrderUid] = order
	}
	 
	return resultService
}