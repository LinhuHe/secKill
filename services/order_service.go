package services

import (
	datamodel "secKillIris/dataModel"
	"secKillIris/repositories"
)

type OrderService interface {
	InsterOrder(*datamodel.Order) (int64, error)
	DeleteOrderById(int64) error
	UpdateOrder(*datamodel.Order) (int64, error)
	// 返回所有订单消息, 以订单形式呈现
	GetAllOrder() ([]*datamodel.Order, error)
	GetOrderByID(id int64) (*datamodel.Order, error)
	// 返回所有订单消息, 以自定义形式呈现
	GetAllWithInfo() (map[int]map[string]string, error)
}

type OrderServiceImp struct {
	orderRepo repositories.IOrder
}

func NewOrderService(repo repositories.IOrder) OrderService {
	return &OrderServiceImp{orderRepo: repo}
}

func (os *OrderServiceImp) InsterOrder(ord *datamodel.Order) (int64, error) {
	return os.orderRepo.Inster(ord)
}

func (os *OrderServiceImp) DeleteOrderById(id int64) error {
	return os.orderRepo.Delete(id)
}

func (os *OrderServiceImp) UpdateOrder(ord *datamodel.Order) (int64, error) {
	return os.orderRepo.Update(ord)
}

func (os *OrderServiceImp) GetAllOrder() ([]*datamodel.Order, error) {
	return os.orderRepo.SelectAll()
}

func (os *OrderServiceImp) GetOrderByID(id int64) (*datamodel.Order, error) {
	return os.orderRepo.SelectByID(id)
}

func (os *OrderServiceImp) GetAllWithInfo() (map[int]map[string]string, error) {
	return os.orderRepo.SelectAllWithInfo()
}
