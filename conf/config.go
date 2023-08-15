package conf

import (
	"github.com/Nolions/cuiFileBrowser/conf/db"
	"github.com/spf13/viper"
	"strings"
)

type Conf struct {
	DB db.Conf `mapstructure:"db"`
}

func init() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func New(file string) (conf *Conf, err error) {
	conf = &Conf{}
	viper.SetConfigFile(file)
	if err := viper.MergeInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
