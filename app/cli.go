package app

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type UserConfig struct {
	binanceUserConfig BinanceUserConfig
	// for other wallets will be here!
}

type BinanceUserConfig struct {
	APIKey    string
	SecretKey string
}

func (uc *UserConfig) GetBinanceAccountConfig() BinanceUserConfig {
	return uc.binanceUserConfig
}

func GetUserConfig() *UserConfig {

	setUpDotEnvFile()

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	uc := &UserConfig{}
	APIKey, ok := viper.Get("BinanceAPIKey").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	uc.binanceUserConfig.APIKey = APIKey
	secretKey, ok := viper.Get("BinanceSecretKey").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	uc.binanceUserConfig.SecretKey = secretKey
	return uc
}

func setUpDotEnvFile() {
	f, err := os.Open(".env")
	if err != nil {
		log.Println(".env not found, creating it...")
		f, err := os.Create(".env")
		if err != nil {
			log.Fatal(err)
		}
		var binanceAPIkey string
		fmt.Println("Enter your binance API key")
		fmt.Scanln(&binanceAPIkey)

		var binanceSecretKey string
		fmt.Println("Now enter your binance secret key")
		fmt.Scanln(&binanceSecretKey)

		// store them in .env
		os.WriteFile(".env", []byte("BinanceAPIKey = " + binanceAPIkey + "\n" + "BinanceSecretKey = " + binanceSecretKey + "\n"), 0644)

		fmt.Println("Done! If you want to change them, please do it at .env file or delete .env file to make appear this configuration messages")
		defer f.Close()
		return
	}
	defer f.Close()
}
