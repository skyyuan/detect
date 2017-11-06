package routers

import (
	"detect/controllers"
	"github.com/astaxie/beego"
)

func init() {
	skiing := beego.NewNamespace("s1",
		beego.NSNamespace("/det",
			beego.NSInclude(
				&controllers.DetectorController{},
			),
		),
	)
	beego.AddNamespace(skiing)
}
