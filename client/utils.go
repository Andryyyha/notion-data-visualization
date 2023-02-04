package client

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetBudgetPageID(configPath string) string {
	viper.SetConfigName("credentials")
	viper.SetConfigType("ini")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return viper.GetString("pages.pageID")
}
