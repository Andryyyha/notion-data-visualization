package client

import (
	"fmt"
	"github.com/jomei/notionapi"
	"github.com/spf13/viper"
)

func InitClient(configPath string) *notionapi.Client {
	viper.SetConfigName("credentials")
	viper.SetConfigType("ini")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	integrationToken := viper.GetString("secrets.token")
	return notionapi.NewClient(notionapi.Token(integrationToken))
}
