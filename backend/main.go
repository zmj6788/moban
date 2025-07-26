package main

import (
	"backend/core"
	"backend/flag"
	"backend/global"
	"backend/routers"
)

func init() {
	// 配置信息读取
	core.InitConf()
	//日志初始化
	global.Log = core.InitLogger()
	//数据库连接
	global.DB = core.Initgorm()
	global.Log.Info("初始化配置完成")
}
func main() {
	//命令行参数绑定迁移表结构函数
	var option = flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		//控制迁移表结构后退出
		return
	}
	//路由初始化
	router := routers.InitRouter()
	//启动服务
	addr := global.Config.System.Addr()
	router.Run(addr)
}
