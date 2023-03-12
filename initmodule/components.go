package initmodule

import (
	"github.com/spf13/viper"
	"gitlab.myshuju.top/heshiying/gin-skeleton/config"
)

func ConfigToModel(file string, m interface{}) error {
	v := viper.NewWithOptions()
	v.SetConfigType("yaml")
	//v.AddConfigPath(wd)
	v.SetConfigName("config")
	v.SetConfigFile(file)
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	return v.Unmarshal(config.GetConfig())
}
