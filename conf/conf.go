/*
Package conf
  - @Author: it.ww@qq.com
  - @Date: 2023/7/20
*/

package conf

import (
	"os"

	"github.com/spf13/viper"
	"github.com/tangyong230/gox/xerror"
)

func LoadConfig[T any](dir, envVar, lang string) (*T, error) {
	vp := viper.New()
	value, isOk := os.LookupEnv(envVar)
	if !isOk {
		return nil, xerror.NewWithLang(xerror.ErrConf, "err.conf.env", lang)
	}
	vp.SetConfigName("app_" + value)
	vp.SetConfigType("toml")
	vp.AddConfigPath(dir)
	if err := vp.ReadInConfig(); err != nil {
		return nil, xerror.NewWithLang(xerror.ErrConf, "err.conf.readfile", lang)
	}

	var config T
	if err := vp.Unmarshal(&config); err != nil {
		return nil, xerror.NewWithLang(xerror.ErrConf, "err.conf.unmarshal", lang)
	}
	return &config, nil
}
