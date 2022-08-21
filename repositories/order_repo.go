package repositories

import (
	"database/sql"
	"secKillIris/common"
	datamodel "secKillIris/dataModel"
)

type IOrder interface {
	ConnDb() error
	Inster(*datamodel.Order) (int64, error)
	Delete(int64) error
	Update(*datamodel.Order) (int64, error)
	SelectAll() ([]*datamodel.Order, error)
	SelectByID(id int64) (*datamodel.Order, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}

type OrderMangerImp struct {
	db    *sql.DB
	table string
}

func NewOrderManger() IOrder {
	return &OrderMangerImp{}
}

func (om *OrderMangerImp) ConnDb() error {
	om.table = "orders"
	if om.db == nil {
		conn, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		om.db = conn
	}
	return nil
}

func (om *OrderMangerImp) Inster(ord *datamodel.Order) (int64, error) {
	if om.db == nil {
		return -1, common.NewError("未初始化的数据库指针")
	}

	sql := "insert " + om.table + " set user_id=?, product_id=?, order_status=?"
	stmt, err := om.db.Prepare(sql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(ord.UserID, ord.ProductId, ord.OrderStataus)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func (om *OrderMangerImp) Delete(id int64) error {
	if om.db == nil {
		return common.NewError("未初始化的数据库指针")
	}

	sql := "delete from " + om.table + " where id=?"
	stmt, err := om.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (om *OrderMangerImp) Update(ord *datamodel.Order) (int64, error) {
	if om.db == nil {
		return 0, common.NewError("未初始化的数据库指针")
	}

	sql := "update " + om.table + " set user_id=?, product_id=?, order_status=? where id=?"

	stmt, err := om.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(ord.UserID, ord.ProductId, ord.OrderStataus, ord.ID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (om *OrderMangerImp) SelectByID(id int64) (*datamodel.Order, error) {
	if om.db == nil {
		return nil, common.NewError("未初始化的数据库指针")
	}

	sql := "select * from " + om.table + " where id=?"

	stmt, err := om.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	res := common.GetResultRow(row)

	if len(res) == 0 {
		return nil, common.NewError("查无数据")
	}
	data := &datamodel.Order{}
	common.DataToStructByTagSql(res, data)

	return data, nil
}

func (om *OrderMangerImp) SelectAll() ([]*datamodel.Order, error) {
	if om.db == nil {
		return nil, common.NewError("未初始化的数据库指针")
	}

	sql := "select * from " + om.table
	rows, err := om.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := common.GetResultRows(rows)
	if len(res) == 0 {
		return nil, common.NewError("查无数据")
	}

	datas := []*datamodel.Order{}
	for _, v := range res {
		orders := &datamodel.Order{}
		common.DataToStructByTagSql(v, orders)
		datas = append(datas, orders)
	}

	return datas, nil
}

func (om *OrderMangerImp) SelectAllWithInfo() (map[int]map[string]string, error) {
	if om.db == nil {
		return nil, common.NewError("未初始化的数据库指针")
	}

	sql := "select o.id, p.product_name, o.order_status from orders as o left join product as p on o.product_id=p.id"
	rows, err := om.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return common.GetResultRows(rows), nil
}
