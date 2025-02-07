package config

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"os"
)

type Config struct {
	Password string `json:"password"`
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetOrInitConfig() (config Config) {

	configFile, err := os.OpenFile("config.json", os.O_RDONLY, 0600)
	defer configFile.Close()
	if err == nil {
		configBytes, err := io.ReadAll(configFile)
		if err != nil {
			log.Fatalf("failed reading file 'config.json': %v", err)
		}
		err = json.Unmarshal(configBytes, &config)
		if err != nil {
			log.Fatalf("failed unmarshaling file 'config.json': %v", err)
		}
	} else {
		log.Printf("failed opening file 'config.json': %v \n", err)
		var passwordRunes []rune = make([]rune, 12)
		for i := range passwordRunes {
			passwordRunes[i] = letters[rand.Intn(len(letters))]
		}

		config = Config{
			Password: string(passwordRunes),
		}

		configJsonBytes, err := json.Marshal(&config)
		if err != nil {
			log.Fatalf("failed marshaling var configJson : %v", err)
		}
		configFileNew, err := os.Create("config.json")
		defer configFileNew.Close()
		if err != nil {
			log.Fatalf("failed creating file 'config.json': %v", err)
		}
		_, err = configFileNew.Write(configJsonBytes)
		if err != nil {
			log.Fatalf("failed writing file 'config.json': %v", err)
		}
	}

	return
}
