package config

import "example.com/ningneng/pkg/global"

func Init() {
	global.Viper.SetConfigName("config")
	global.Viper.SetConfigType("yml")
	global.Viper.AddConfigPath(global.RootPath + "/config")

	if err := global.Viper.ReadInConfig(); err != nil {
		panic(err)
	}

	//global.Viper.WatchConfig()
	//global.Viper.OnConfigChange(func(in fsnotify.Event) {
	//	if in.Op.String() == "WRITE" {
	//		// 被回调了两次
	//		fmt.Println("文件被修改了", in.Name, time.Now().String())
	//	}
	//})
}
