package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port        int    `mapstructure:"port"`
	Version     string `mapstructure:"version"`
	MysqlConfig `mapstructure:"mysql"`
}

type MysqlConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	DbName string `mapstructure:"dbname"`
}

func main() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("read config success!!!")

	k := viper.GetString("clothing.jacket") //通过.获取嵌套字段
	fmt.Println(k)

	//提取子树
	sub := viper.Sub("mysql") //通过.获取嵌套字段
	fmt.Println(sub)

	// r := gin.Default()
	// r.GET("/name", func(c *gin.Context) {
	// 	c.String(http.StatusOK, viper.GetString("name"))
	// })
	// r.Run()

	//序列化
	var C Config

	err = viper.Unmarshal(&C)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v\n", err)
		return
	}
	fmt.Printf("c:%v\n", C)

}
