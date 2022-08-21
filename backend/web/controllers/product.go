package controllers

import (
	"fmt"
	"secKillIris/common"
	datamodel "secKillIris/dataModel"
	"secKillIris/services"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ProductController struct {
	Ctx iris.Context
	Ps  services.IProductService
}

func (p *ProductController) GetAll() mvc.View {
	data, err := p.Ps.GetAllProduct()
	defer func() {
		if f := recover(); f != nil {
			fmt.Println("Recover follows panic: \n", f)
		}
	}()
	if err != nil {
		panic(err)
	}

	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productAllInfo": data,
		},
	}
}

func (p *ProductController) PostUpdate() {
	product := &datamodel.Product{}
	p.Ctx.Request().ParseForm()

	decode := common.NewDecoder(&common.DecoderOptions{TagName: "hlh"})
	if err := decode.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	_, err := p.Ps.UpdateProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	//  Q. 这句话是干嘛的  执行完后重定向到某个页面
	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) GetAdd() mvc.View {
	return mvc.View{
		Name: "product/add.html",
	}
}

func (p *ProductController) PostAdd() {
	product := &datamodel.Product{}
	p.Ctx.Request().ParseForm()

	decode := common.NewDecoder(&common.DecoderOptions{TagName: "hlh"})
	if err := decode.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	_, err := p.Ps.InsertProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) GetManager() mvc.View {
	id, err := strconv.ParseInt(p.Ctx.URLParam("id"), 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
		id = -1
	}

	productModel, _ := p.Ps.GetProductByID(id)
	return mvc.View{
		Name: "product/manager.html",
		Data: iris.Map{
			"product": productModel,
		},
	}
}

func (p *ProductController) GetDelete() {
	id, err := strconv.ParseInt(p.Ctx.URLParam("id"), 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	err = p.Ps.DeleteProduct(id)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}
