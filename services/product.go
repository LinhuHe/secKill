package services

import (
	datamodel "secKillIris/dataModel"
	"secKillIris/repositories"
)

type IProductService interface {
	InsertProduct(*datamodel.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*datamodel.Product) (int64, error)
	GetProductByID(int64) (*datamodel.Product, error)
	GetAllProduct() ([]*datamodel.Product, error)
}

type ProductServiceImp struct {
	pm repositories.IProductManger
}

func NewProductService(pm repositories.IProductManger) IProductService {
	return &ProductServiceImp{pm: pm}
}

func (psi *ProductServiceImp) InsertProduct(p *datamodel.Product) (int64, error) {
	return psi.pm.Insert(p)
}

func (psi *ProductServiceImp) DeleteProduct(id int64) error {
	return psi.pm.Delete(id)
}
func (psi *ProductServiceImp) UpdateProduct(p *datamodel.Product) (int64, error) {
	return psi.pm.Update(p)
}

func (psi *ProductServiceImp) GetProductByID(id int64) (*datamodel.Product, error) {
	return psi.pm.SelectByID(id)
}

func (psi *ProductServiceImp) GetAllProduct() ([]*datamodel.Product, error) {
	return psi.pm.SelectAll()
}
