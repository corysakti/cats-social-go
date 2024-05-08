package config

import "github.com/spf13/viper"

func NewViper() {
	var config = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	//yaml
	//config.SetConfigFile("config.yaml")
	//config.AddConfigPath(".")

	//membaca config
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config.GetString("app.name")
	config.GetString("app.author")
	config.GetString("database.host")
	config.GetString("database.port")
	config.GetString("database.show_sql")
}
