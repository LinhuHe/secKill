package controller

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	datamodel "secKillIris/dataModel"
	"secKillIris/services"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ProductController struct {
	Ctx iris.Context
	Ps  services.IProductService
	Os  services.OrderService
}

var (
	htmlOutPath  = "../frontend/web/htmlProductShow"
	templatePath = "../frontend/web/views/template"
)

func (p *ProductController) GetGenerateHtml() {
	productId := p.Ctx.URLParam("productID")
	pid, err := strconv.ParseInt(productId, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Error("convert fail:", productId)
	}

	__file__, _ := os.Getwd()

	//1.获取模版
	contTmp, err := template.ParseFiles(filepath.Join(__file__, templatePath, "product.html"))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	//2.获取html生成路径
	fileName := filepath.Join(__file__, htmlOutPath, "htmlProduct.html")

	//3.获取模版渲染数据
	product, err := p.Ps.GetProductByID(pid)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	//4.生成静态文件
	generateStaticHtml(p.Ctx, contTmp, fileName, product)
}

func generateStaticHtml(ctx iris.Context, template *template.Template, fileName string, product *datamodel.Product) {
	if fileEixst(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			ctx.Application().Logger().Error(err)
		}
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
	defer file.Close()

	template.Execute(file, &product)
}

func fileEixst(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

func (p *ProductController) GetDetail() mvc.View {
	errFunc := func(err interface{}) {
		p.Ctx.Redirect("/user/error")
		p.Ctx.Application().Logger().Error(err)

	}

	id, err := strconv.ParseInt(p.Ctx.URLParam("productId"), 10, 64)
	if err != nil {
		errFunc(err)
		return mvc.View{}
	}

	res, err := p.Ps.GetProductByID(id)
	if err != nil {
		errFunc(err)
		return mvc.View{}
	}

	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": res,
		},
	}
}

func (p *ProductController) GetOrder() mvc.View {
	defer func() {
		if rc := recover(); rc != nil {
			fmt.Println(rc)
		}
	}()

	productId := p.Ctx.URLParam("productID")
	userId := p.Ctx.GetCookie("uid")

	pid, err := strconv.ParseInt(productId, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Error("convert fail:", productId)
	}
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Error("convert fail:", productId)
	}

	var oid int64
	var msg string
	// TODO 这里select * 可以优化
	proInfo, err := p.Ps.GetProductByID(pid)
	if err != nil {
		p.Ctx.Application().Logger().Error("can't find data by id:", productId)
	}

	if proInfo.ProductNum > 0 {
		proInfo.ProductNum -= 1
		_, err := p.Ps.UpdateProduct(proInfo)
		if err != nil {
			p.Ctx.Application().Logger().Error("update fail:")
		}

		od := &datamodel.Order{
			UserID:       uid,
			ProductId:    pid,
			OrderStataus: datamodel.OrderSuccess,
		}

		oid, err = p.Os.InsterOrder(od)
		if err != nil {
			p.Ctx.Application().Logger().Error("create order fail: ", pid, oid, err)
		}
		msg = "抢购成功!!!!!!"
	} else {
		msg = "抢购失败TAT~"
	}

	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/result.html",
		Data: iris.Map{
			"orderID":     oid,
			"showMessage": msg,
		},
	}
}

func (p *ProductController) GetOrderTest() mvc.View {
	defer func() {
		if rc := recover(); rc != nil {
			fmt.Println(rc)
		}
	}()

	start := time.Now().UnixMilli()

	productId := "1"
	userId := "2"

	pid, err := strconv.ParseInt(productId, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Error("convert fail:", productId)
	}
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Error("convert fail:", productId)
	}

	var oid int64
	var msg string
	// TODO 这里select * 可以优化
	proInfo, err := p.Ps.GetProductByID(pid)
	if err != nil {
		panic(err)
	}

	if proInfo.ProductNum > 0 {
		proInfo.ProductNum -= 1
		_, err := p.Ps.UpdateProduct(proInfo)
		if err != nil {
			p.Ctx.Application().Logger().Error("update fail:")
		}

		od := &datamodel.Order{
			UserID:       uid,
			ProductId:    pid,
			OrderStataus: datamodel.OrderSuccess,
		}

		oid, err = p.Os.InsterOrder(od)
		if err != nil {
			p.Ctx.Application().Logger().Error("create order fail: ", err)
		}
		msg = "抢购成功!!!!!!"
	} else {
		msg = "抢购失败TAT~"
	}

	fmt.Println("USE TIME: ", time.Now().UnixMilli()-start)

	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/result.html",
		Data: iris.Map{
			"orderID":     oid,
			"showMessage": msg,
		},
	}
}
