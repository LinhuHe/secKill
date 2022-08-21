package middleware

import "github.com/kataras/iris/v12"

func AuthConProduct(ctx iris.Context) {
	name := ctx.GetCookie("login_name")
	if name == "" {
		ctx.Application().Logger().Debug("必须先登录!")
		ctx.Redirect("/user/login")
		return
	}
	ctx.Application().Logger().Debug("已经登陆")
	ctx.Next()
}
