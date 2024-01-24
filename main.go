package main

import (
	"xpass/resource"
	"xpass/service"

	"github.com/getlantern/systray"
)

// 主入口函数
func main() {

	// 读取配置
	service.ReadConfig()

	// 执行同步预操作
	service.RunPreoperationSync()

	// 执行异步预操作
	go service.RunPreoperationAsync()

	// 加载系统托盘图标（仅MS Windows下生效）
	systray.Run(onReady, onExit)

}

func onReady() {
	// 设置任务栏Icon图标
	systray.SetIcon(resource.IconData)
	// 设置标题（仅Mac和Linux有效）
	systray.SetTitle("XPsss Client")
	// 设置提示内容（鼠标悬浮一段时间出现）
	systray.SetTooltip("XPass 内网穿透")
	install := systray.AddMenuItem("服务安装", "安装")
	uninstall := systray.AddMenuItem("服务卸载", "卸载")
	systray.AddSeparator()
	start := systray.AddMenuItem("服务启动", "启动")
	stop := systray.AddMenuItem("服务停止", "停止")
	restart := systray.AddMenuItem("服务重启", "重启")
	systray.AddSeparator()
	edit := systray.AddMenuItem("修改配置", "重启")
	showLog := systray.AddMenuItem("打开日志", "日志")
	setting := systray.AddMenuItem("系统设置", "设置")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出程序")

	go func() {
		for {
			select {
			case <-showLog.ClickedCh:
				service.ShowLog()
			case <-install.ClickedCh:
				service.InstallServe()
			case <-uninstall.ClickedCh:
				service.UninstallServe()
			case <-start.ClickedCh:
				service.RunServe()
			case <-stop.ClickedCh:
				service.StopServe()
			case <-restart.ClickedCh:
				service.RestartServe()
			case <-setting.ClickedCh:
				service.ShowSetting()
			case <-edit.ClickedCh:
				service.EditConf()
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onExit() {}
