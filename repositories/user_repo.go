package repositories

import (
	"database/sql"
	"errors"
	"secKillIris/common"
	datamodel "secKillIris/dataModel"
)

type IUserManger interface {
	ConnDB() error
	SelectName(string) (*datamodel.User, error)
	Insert(*datamodel.User) (int64, error)
}

type UserMangerImp struct {
	table string
	db    *sql.DB
}

func NewUserManger() IUserManger {
	return &UserMangerImp{}
}

func (um *UserMangerImp) ConnDB() error {
	um.table = "user"
	if um.db == nil {
		newDb, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		um.db = newDb
	}
	return nil
}

func (um *UserMangerImp) SelectName(userName string) (*datamodel.User, error) {
	if um.db == nil {
		return &datamodel.User{}, common.NewError("未初始化的数据库指针")
	}

	sql := "select * from " + um.table + " where user_name = ?"
	res, err := um.db.Query(sql, userName)
	if err != nil {
		return &datamodel.User{}, err
	}
	defer res.Close()

	data := common.GetResultRow(res)
	if len(data) == 0 {
		return &datamodel.User{}, errors.New("用户不存在")
	}

	ret := &datamodel.User{}
	common.DataToStructByTagSql(data, ret)

	return ret, nil
}

func (um *UserMangerImp) Insert(usr *datamodel.User) (int64, error) {
	if um.db == nil {
		return -1, common.NewError("未初始化的数据库指针")
	}

	sql := "insert " + um.table + " set nick_name = ?, user_name = ?, pass_word = ?"
	res, err := um.db.Exec(sql, usr.NickName, usr.UserName, usr.Password)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (um *UserMangerImp) SelectID(ID int) (*datamodel.User, error) {
	if um.db == nil {
		return &datamodel.User{}, common.NewError("未初始化的数据库指针")
	}

	sql := "select * from" + um.table + " where id = ?"
	res, err := um.db.Query(sql, ID)
	if err != nil {
		return &datamodel.User{}, err
	}
	defer res.Close()

	data := common.GetResultRow(res)
	if len(data) == 0 {
		return &datamodel.User{}, errors.New("用户不存在")
	}

	ret := &datamodel.User{}
	common.DataToStructByTagSql(data, ret)

	return ret, nil
}
