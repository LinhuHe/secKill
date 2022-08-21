package main

import (
	"context"
	"secKillIris/backend/web/controllers"
	"secKillIris/frontend"
	"secKillIris/repositories"
	"secKillIris/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	// mp := map[string]string{
	// 	"id":           "111",
	// 	"product_name": "mmike",
	// 	"product_num":  "24",
	// 	"product_img":  "safaedfafafadfadfad",
	// 	"product_url":  "www.dsfnsduifhsdui.com",
	// }

	// data := &datamodel.Product{}
	// common.DataToStructByTagSql(mp, data)

	app := iris.New()
	app.Logger().SetLevel("debug")

	// 该路径是相较于main.go文件所在路径的
	tmplate := iris.HTML("./web/views", ".html").Layout("shared/layout.html").Reload(true)

	app.RegisterView(tmplate)
	// 设置模板目标
	app.HandleDir("assets", iris.Dir("./web/assets"))
	// 异常跳转
	app.OnAnyErrorCode(func(ctx iris.Context) {
		// 这里最终返回的就是一个字符串罢了, 怀疑直接value填写固定值是不是都可以
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问到外太空了"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	// 注册控制器
	//product 控制器
	prodManger := repositories.NewProductManger()
	err := prodManger.ConnDb()
	if err != nil {
		panic(err)
	}

	prodService := services.NewProductService(prodManger)
	prodParty := app.Party("/product")
	mvcPrdApp := mvc.New(prodParty)

	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	mvcPrdApp.Register(ctx, prodService)
	mvcPrdApp.Handle(new(controllers.ProductController))

	// order 控制器
	ordManger := repositories.NewOrderManger()
	err = ordManger.ConnDb()
	if err != nil {
		panic(err)
	}

	ordService := services.NewOrderService(ordManger)
	ordParty := app.Party("/order")
	mvcOrdApp := mvc.New(ordParty)

	mvcOrdApp.Register(ctx, ordService)
	mvcOrdApp.Handle(new(controllers.OrderController))

	go frontend.InitFrontendRouter()
	go frontend.InitHtmlPage()
	// 启动服务
	app.Run(iris.Addr(":8080"), iris.WithOptimizations)

}
