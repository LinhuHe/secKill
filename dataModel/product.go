package datamodel

type Product struct {
	ID           uint64 `json:"id" sql:"id" hlh:"id"`
	ProductName  string `json:"prodName" sql:"product_name" hlh:"prodName"`
	ProductNum   uint64 `json:"prodNum" sql:"product_num" hlh:"prodNum"`
	ProductImage string `json:"prodImg" sql:"product_img" hlh:"prodImg"`
	ProductUrl   string `json:"prodUrl" sql:"product_url" hlh:"prodUrl"`
}
