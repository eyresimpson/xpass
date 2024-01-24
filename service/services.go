package service

import (
	"xpass/tools"
)

// 拉起UI界面
func RunUI(mode string) {
	// 调用运行当前目录下的ui.exe程序
	// 带main参数代表正常启动
	// 带log参数代表默认进入日志界面
	// 带setting参数代表默认进设置界面
}

// 同步预执行
func RunPreoperationSync() {

}

// 异步预执行
func RunPreoperationAsync() {
	// 打开UI界面，注释掉就不会默认打开
	RunUI("main")
}

// 安装服务
func InstallServe() {
	err := tools.RunExe("./bin/service.exe", "install")
	if err != nil {
		tools.ShowMessage("XPass", "服务安装失败")
		print(err.Error())
	} else {
		tools.ShowMessage("XPass", "服务安装成功")
	}
}

// 卸载服务
func UninstallServe() {
	err := tools.RunExe("./bin/service.exe", "uninstall")
	if err != nil {
		tools.ShowMessage("XPass", "服务卸载失败")
		print(err.Error())
	} else {
		tools.ShowMessage("XPass", "服务卸载成功")
	}
}

// 运行服务
func RunServe() {
	err := tools.RunExe("./bin/service.exe", "start")
	if err == nil {
		tools.ShowMessage("XPass", "服务已启动")
	} else {
		tools.ShowMessage("XPass", "服务启动失败，检查运行权限和服务状态")
	}
}

// 停止服务
func StopServe() {
	err := tools.RunExe("./bin/service.exe", "stop")
	if err == nil {
		tools.ShowMessage("XPass", "服务已停止")
	} else {
		tools.ShowMessage("XPass", "服务停止失败，检查运行权限和服务状态")
	}
}

// 检查更新
func CheckUpdate(service_addr, server_port, username, password, request_version string) {
	// 检查参数
	tools.ShowMessage("更新程序", "更新测试")
}

// 重启服务
func RestartServe() {
	StopServe()
	RunServe()
}

// 调用UI的日志界面
func ShowLog() {
	// 模拟类似于Tail的实时文件监控（前端UI）
}

// 调用UI的设置界面
func ShowSetting() {

}

// 修改配置文件，后续改为图形化界面
func EditConf() {
	err5 := tools.RunExe("notepad", "config/main.toml")
	if err5 != nil {
		tools.ShowMessage("XPass", "无法打开配置文件，检查配置是否存在")
	}
}
