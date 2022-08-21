package repositories

import (
	"database/sql"
	"secKillIris/common"
	datamodel "secKillIris/dataModel"
)

type IProductManger interface {
	ConnDb() error
	Insert(*datamodel.Product) (int64, error)
	Delete(int64) error
	Update(*datamodel.Product) (int64, error)
	SelectByID(int64) (*datamodel.Product, error)
	SelectAll() ([]*datamodel.Product, error)
}

type ProductMangerImp struct {
	table     string
	mysqlConn *sql.DB
}

func NewProductManger() IProductManger {
	return &ProductMangerImp{}
}

func (pm *ProductMangerImp) ConnDb() error {
	pm.table = "product"
	if pm.mysqlConn == nil {
		var err error
		pm.mysqlConn, err = common.NewMysqlConn()
		if err != nil {
			return err
		}
	}

	return nil
}

func (pm *ProductMangerImp) Insert(p *datamodel.Product) (int64, error) {
	if pm.mysqlConn == nil {
		return -1, common.NewError("未初始化的数据库指针")
	}

	sql := "insert " + pm.table + " set product_name=?, product_num=?, product_img=?, product_url=?"
	stmt, err := pm.mysqlConn.Prepare(sql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(p.ProductName, p.ProductNum, p.ProductImage, p.ProductUrl)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func (pm *ProductMangerImp) Delete(id int64) error {
	if pm.mysqlConn == nil {
		return common.NewError("未初始化的数据库指针")
	}

	sql := "delete from " + pm.table + " where id=?"
	stmt, err := pm.mysqlConn.Prepare(sql)
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

func (pm *ProductMangerImp) Update(p *datamodel.Product) (int64, error) {
	if pm.mysqlConn == nil {
		return 0, common.NewError("未初始化的数据库指针")
	}

	sql := "update " + pm.table + " set product_name=?, product_num=?, product_img=?, product_url=? where id=?"

	stmt, err := pm.mysqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(p.ProductName, p.ProductNum, p.ProductImage, p.ProductUrl, p.ID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (pm *ProductMangerImp) SelectByID(id int64) (*datamodel.Product, error) {
	if pm.mysqlConn == nil {
		return nil, common.NewError("未初始化的数据库指针")
	}

	sql := "select * from " + pm.table + " where id=?"

	stmt, err := pm.mysqlConn.Prepare(sql)
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
	data := &datamodel.Product{}
	common.DataToStructByTagSql(res, data)

	return data, nil
}

func (pm *ProductMangerImp) SelectAll() ([]*datamodel.Product, error) {
	if pm.mysqlConn == nil {
		return nil, common.NewError("未初始化的数据库指针")
	}

	sql := "select * from " + pm.table
	rows, err := pm.mysqlConn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := common.GetResultRows(rows)
	if len(res) == 0 {
		return nil, common.NewError("查无数据")
	}

	datas := []*datamodel.Product{}
	for _, v := range res {
		product := &datamodel.Product{}
		common.DataToStructByTagSql(v, product)
		datas = append(datas, product)
	}

	return datas, nil
}
