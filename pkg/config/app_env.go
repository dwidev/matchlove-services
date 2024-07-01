package config

var config *Schema

type Schema struct {
	// server property
	ServerPort string `mapstructure:"SERVER_PORT"`

	// JWT property
	AccessSecretKey  string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshSecretKey string `mapstructure:"REFRESH_TOKEN_SECRET"`

	// db property
	DB_DSN string `mapstructure:"DB_DSN"`
}

func Load() *Schema {
	configLoader := YamlLoader()

	config, err := configLoader.Run()
	if err != nil {
		panic(err)
	}

	return config
}

func Get() *Schema {
	return config
}
