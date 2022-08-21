package datamodel

type User struct {
	ID       int64  `json:"id" form:"ID" sql:"id"`
	NickName string `json:"nick_name" form:"NickName" sql:"nick_name"`
	UserName string `json:"user_name" form:"UserName" sql:"user_name" hlh:"userName"`
	Password string `json:"_" form:"Password" sql:"pass_word" hlh:"password"`
}
