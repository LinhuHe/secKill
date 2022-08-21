package frontend

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"secKillIris/frontend/middleware"
	"secKillIris/frontend/web/controller"
	"secKillIris/repositories"
	"secKillIris/services"
)

func InitFrontendRouter() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// 该路径是相较于main.go文件所在路径的
	tmplate := iris.HTML("../frontend/web/views", ".html").Layout("shared/layout.html").Reload(true)

	app.RegisterView(tmplate)
	// 设置模板目标
	app.HandleDir("public", iris.Dir("../frontend/web/public"))
	app.HandleDir("assets", iris.Dir("./web/assets"))
	// 异常跳转
	app.OnAnyErrorCode(func(ctx iris.Context) {
		// 这里最终返回的就是一个字符串罢了, 怀疑直接value填写固定值是不是都可以
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问到外太空了"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	// 注册控制器
	//user 控制器
	userManger := repositories.NewUserManger()
	err := userManger.ConnDB()
	if err != nil {
		panic(err)
	}

	userService := services.NewUserService(userManger)
	userParty := app.Party("/user")
	mvcUserApp := mvc.New(userParty)

	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	mvcUserApp.Register(ctx, userService)
	mvcUserApp.Handle(new(controller.UsersController))

	//product 控制器
	prodManger := repositories.NewProductManger()
	err = prodManger.ConnDb()
	if err != nil {
		panic(err)
	}

	prodService := services.NewProductService(prodManger)

	ordManger := repositories.NewOrderManger()
	err = ordManger.ConnDb()
	if err != nil {
		panic(err)
	}

	ordService := services.NewOrderService(ordManger)

	prodParty := app.Party("/product")
	prodParty.Use(middleware.AuthConProduct)
	mvcProdApp := mvc.New(prodParty)

	mvcProdApp.Register(ctx, prodService, ordService)
	mvcProdApp.Handle(new(controller.ProductController))

	app.Run(iris.Addr(":8081"), iris.WithOptimizations)
}

func InitHtmlPage() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.HandleDir("/public", iris.Dir("../frontend/web/public"))
	app.HandleDir("/html", iris.Dir("../frontend/web/htmlProductShow/"))

	app.Run(iris.Addr(":80"))
}
