// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"goniushop/controllers"
	"goniushop/controllers/api"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/goods",
			beego.NSInclude(
				&api.NsGoodsController{},
			),
		),
		beego.NSNamespace("/adv",
			beego.NSInclude(
				&api.NsPlatformAdvController{},
			),
		),
		beego.NSNamespace("/cart",
			beego.NSInclude(
				&api.NsCartController{},
			),
		),
		beego.NSNamespace("/order",
			beego.NSInclude(
				&api.NsOrderController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
