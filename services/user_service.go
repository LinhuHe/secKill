package services

import (
	"secKillIris/common"
	datamodel "secKillIris/dataModel"
	"secKillIris/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	IsLoginSuccess(userName string, pwd string) (*datamodel.User, bool)
	AddUser(*datamodel.User) (int64, error)
	GetUserIdByName(string) int64
}

type UserService struct {
	userRepo repositories.IUserManger
}

func NewUserService(repo repositories.IUserManger) IUserService {
	return &UserService{repo}
}

func (us *UserService) IsLoginSuccess(userName string, pwd string) (*datamodel.User, bool) {
	get, err := us.userRepo.SelectName(userName)
	if err != nil {
		return &datamodel.User{}, false
	}
	return get, checkPassword(get.Password, pwd)
}

func (us *UserService) GetUserIdByName(userName string) int64 {
	res, err := us.userRepo.SelectName(userName)
	if err != nil {
		return -1
	}
	return res.ID
}

func (us *UserService) AddUser(user *datamodel.User) (int64, error) {
	hash, err := generatePassword(user.Password)
	if err != nil {
		return -1, err
	}
	user.Password = hash
	return us.userRepo.Insert(user)
}

func generatePassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", common.NewError("密码的错误转换", pwd)
	}
	return string(hash), nil
}

func checkPassword(query string, input string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(query), []byte(input)); err != nil {
		return false
	}

	return true
}
