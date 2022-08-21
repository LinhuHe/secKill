package controllers

import (
	"secKillIris/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type OrderController struct {
	Ctx iris.Context
	Os  services.OrderService
}

func (oc *OrderController) Get() mvc.View {
	datas, err := oc.Os.GetAllWithInfo()
	if err != nil {
		oc.Ctx.Application().Logger().Debug(err)
	}

	return mvc.View{
		Name: "order/view.html",
		Data: iris.Map{
			"ordersData": datas,
		},
	}
}
