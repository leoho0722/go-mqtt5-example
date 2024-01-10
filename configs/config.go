package configs

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadEnvConfig() *Configs {
	envFile, err := os.Open("config.json")
	defer func(envConfig *os.File) {
		err := envConfig.Close()
		if err != nil {
			panic(err)
		}
	}(envFile)

	if err != nil {
		panic(err)
	}
	fmt.Println("config.json open successfully")

	envConfigs, _ := io.ReadAll(envFile)
	var configs Configs
	err = json.Unmarshal(envConfigs, &configs)
	if err != nil {
		panic(err)
	}
	fmt.Println("config.json unmarshal successfully")

	return &configs
}

type Configs struct {
	MQTT []MQTTConfig `json:"mqtt"`
}
