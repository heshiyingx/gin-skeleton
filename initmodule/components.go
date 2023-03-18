package initmodule

import (
	"github.com/spf13/viper"
	"gitlab.myshuju.top/heshiying/gin-skeleton/pkg/utils"
)

func ConfigToModel(file string, m interface{}) error {
	v := viper.NewWithOptions()
	v.SetConfigType("yaml")
	//v.AddConfigPath(wd)
	//reflectT(m)

	v.SetConfigName("config")
	v.SetConfigFile(file)
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	err = v.Unmarshal(m)
	if err != nil {
		return err
	}
	utils.FillDefault(m)
	return nil
}
