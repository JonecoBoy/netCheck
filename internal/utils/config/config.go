package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MongoURI   string
	MongoDB    string
	ServerPort string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName(".env") // Look for .env file
	viper.SetConfigType("env")  // Format
	viper.AddConfigPath(".")    // Look in the current directory
	viper.AutomaticEnv()        // Override with environment variables

	// Set default values
	viper.SetDefault("MONGO_URI", "mongodb://localhost:27017")
	viper.SetDefault("MONGO_DB", "cron_db")
	viper.SetDefault("SERVER_PORT", "8081")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No .env file found, using defaults and environment variables")
	}

	AppConfig = Config{
		MongoURI:   viper.GetString("MONGO_URI"),
		MongoDB:    viper.GetString("MONGO_DB"),
		ServerPort: viper.GetString("SERVER_PORT"),
	}

	log.Println("Configuration Loaded:", AppConfig)
}
