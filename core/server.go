package core

import (
	"fmt"
	"time"

	"github.com/wangrui19970405/wu-shi-admin/server/global"
	"github.com/wangrui19970405/wu-shi-admin/server/initialize"
	"github.com/wangrui19970405/wu-shi-admin/server/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.WUSHI_CONFIG.System.UseMultipoint || global.WUSHI_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.WUSHI_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.WUSHI_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.WUSHI_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	波动科技，未来会更好
	药要吃私人播客:https://moonwife.top
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.WUSHI_LOG.Error(s.ListenAndServe().Error())
}
