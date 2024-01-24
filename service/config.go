package service

// TODO: 尝试读取配置文件，必须在同级目录下的config中，目前仅支持yml文件
func ReadConfig() {
	// dataBytes, err := os.ReadFile("config/sys.yml")
	// if err != nil {
	// 	tools.Err("Can't read config", err)
	// 	return
	// }
	// // fmt.Println("yaml 文件的内容: \n", string(dataBytes))
	// config := entity.SysConfig{}
	// err = yaml.Unmarshal(dataBytes, &config)
	// if err != nil {
	// 	tools.Err("Can't parse config", err)
	// 	return
	// }
	// // fmt.Printf("config → %+v\n", config) // config → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}
	// // tools.Info("config → %+v\n" + config)
	// mp := make(map[string]any, 2)
	// err = yaml.Unmarshal(dataBytes, mp)
	// if err != nil {
	// 	fmt.Println("解析 yaml 文件失败：", err)
	// 	tools.Err("Can't parse config", err)
	// 	return
	// }
	// fmt.Printf("map → %+v", config) // config → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}

}

// 尝试读取执行脚本
func ReadExecScript() {

}
