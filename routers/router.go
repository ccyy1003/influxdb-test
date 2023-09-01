package routers

import (
	"influxdb-test/controllers"

	"github.com/astaxie/beego"
)

func init() {
	testNamespace := beego.NewNamespace("/influxdb_test",
		// admin api
		beego.NSRouter("/cq", &controllers.TaskController{}, "*:CQTask"),
		beego.NSRouter("/func", &controllers.TaskController{}, "get:FuncTask"),
		beego.NSRouter("/hint", &controllers.TaskController{}, "get:HintTask"),
		beego.NSRouter("/interface", &controllers.TaskController{}, "get:InterfaceTask"),
		beego.NSRouter("/mathopt", &controllers.TaskController{}, "get:MathOptTask"),
		beego.NSRouter("/mgdb", &controllers.TaskController{}, "get:MgDbTask"),
		beego.NSRouter("/query", &controllers.TaskController{}, "get:QueryTask"),
		beego.NSRouter("/show", &controllers.TaskController{}, "get:ShowTask"),
		beego.NSRouter("/all", &controllers.TaskController{}, "get:AllTask"),
	)
	beego.AddNamespace(testNamespace)

	// beego.Router("/influxdb_test/CQTask", &controllers.TaskController{}, "get:CQTask")
	// beego.Router("/influxdb_test/FuncTask", &controllers.TaskController{}, "get:FuncTask")
	// beego.Router("/influxdb_test/HintTask", &controllers.TaskController{}, "get:HintTask")
	// beego.Router("/influxdb_test/InterfaceTask", &controllers.TaskController{}, "get:InterfaceTask")
	// beego.Router("/influxdb_test/MathOptTask", &controllers.TaskController{}, "get:MathOptTask")
	// beego.Router("/influxdb_test/MgDbTask", &controllers.TaskController{}, "get:MgDbTask")
	// beego.Router("/influxdb_test/QueryTask", &controllers.TaskController{}, "get:QueryTask")
	// beego.Router("/influxdb_test/ShowTask", &controllers.TaskController{}, "get:ShowTask")
	// beego.Router("/influxdb_test/AllTask", &controllers.TaskController{}, "get:AllTask")
}
