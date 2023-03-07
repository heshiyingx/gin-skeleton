package initmodule

import (
	"github.com/spf13/viper"
	"gitlab.myshuju.top/heshiying/gin-skeleton/config"
	"os"
)

func Init() {
	configInit()

}
func configInit() {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	//executable, err := os.Executable()
	//if err != nil {
	//	return
	//}

	//g.Debug("%v;executable:%v", wd, executable)
	v := viper.NewWithOptions()
	v.SetConfigType("yaml")
	v.AddConfigPath(wd)
	v.SetConfigName("config")
	err = v.ReadInConfig()
	if err != nil {
		panic(err)
		return
	}
	err = v.Unmarshal(config.GetConfig())
	if err != nil {
		panic(err)
		return
	}
}
