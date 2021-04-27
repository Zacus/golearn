package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	// TaillogConf `ini:"taillog"`
	EtcdConf `ini:"etcd"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type TaillogConf struct {
	FileName string `ini:"path"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Key     string `ini:"collect_log_key"`
	Timeout int    `ini:"timeout"`
}

// func Init() {
// 	//加载文件
// 	cfg, err := ini.Load("./config.ini")
// 	if err != nil {
// 		fmt.Printf("Fail to read file: %v", err)
// 		os.Exit(1)
// 	}
// }
/*
func Init() (cfg *AppConf, err error) {
	//
	cfg = new(AppConf)
	err = ini.MapTo(cfg, "./config.ini")
	if err != nil {
		// fmt.Printf("Fail to read file: %v", err)
		return
	}
	return
}
*/
