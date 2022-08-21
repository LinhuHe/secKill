package controller

import (
	datamodel "secKillIris/dataModel"
	"secKillIris/services"
	"secKillIris/tool"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UsersController struct {
	Ctx iris.Context
	Us  services.IUserService
}

func (uc *UsersController) GetRegister() mvc.View {
	return mvc.View{
		Name: "user/register.html",
	}
}

func (uc *UsersController) PostRegister() {
	user := &datamodel.User{}
	user.NickName = uc.Ctx.FormValue("nickName")
	user.UserName = uc.Ctx.FormValue("userName")
	user.Password = uc.Ctx.FormValue("password")

	_, err := uc.Us.AddUser(user)
	if err != nil {
		uc.Ctx.Redirect("/user/error")
		defer func() {
			if rc := recover(); rc != nil {
				uc.Ctx.Application().Logger().Debug(err.Error())
			}
		}()
		panic(err)
	}

	uc.Ctx.Redirect("/user/login")
}

func (uc *UsersController) GetLogin() mvc.View {
	return mvc.View{
		Name: "user/login.html",
	}
}

func (uc *UsersController) PostLogin() mvc.Response {
	uname := uc.Ctx.FormValue("userName")
	pwd := uc.Ctx.FormValue("password")

	user, login := uc.Us.IsLoginSuccess(uname, pwd)
	if !login {
		return mvc.Response{
			Path: "login",
		}
	}

	tool.GlobalCookie(uc.Ctx, "login_name", uname)
	tool.GlobalCookie(uc.Ctx, "uid", strconv.FormatInt(user.ID, 10))

	encapUid, _ := tool.EnPwdCode([]byte(string(user.ID)))
	tool.GlobalCookie(uc.Ctx, "uidEncode", encapUid)
	return mvc.Response{
		Path: "/product/",
	}
}
