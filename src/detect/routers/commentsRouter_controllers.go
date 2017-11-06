package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["detect/controllers:DetectorController"] = append(beego.GlobalControllerRouter["detect/controllers:DetectorController"],
		beego.ControllerComments{
			Method: "Detect",
			Router: `/detect`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
