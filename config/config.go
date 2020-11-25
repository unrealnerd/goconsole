package config

import (
	"fmt"
	"github.com/spf13/viper")

var (
	// Data ...
	Data configuration
)

//Configuration ...
type configuration struct {
	ConnectionString string
	DefaultNameSpace string
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())

		fmt.Println("ConnectionSting:", viper.GetString("DefaultNamespace"))

		err := viper.Unmarshal(&Data)
		if err != nil {
			fmt.Printf("Unable to decode into struct, %v", err)
		}
	}
}
